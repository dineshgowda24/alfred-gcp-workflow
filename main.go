package main

import (
	"bytes"
	"embed"
	"flag"
	"log"
	"strings"
	"text/template"

	aw "github.com/deanishe/awgo"
	"github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
	"github.com/dineshgowda24/alfred-gcp-workflow/parser"
	"github.com/dineshgowda24/alfred-gcp-workflow/searchers"
	"github.com/dineshgowda24/alfred-gcp-workflow/services"
)

var (
	wf         *aw.Workflow
	query      string
	forceFetch bool
	//go:embed services.yml
	servicesFile embed.FS
)

func init() {
	flag.StringVar(&query, "query", "", "Query to use")
	flag.BoolVar(&forceFetch, "fetch", false, "Force fetch bypassing cache")
	flag.Parse()
	wf = aw.New()
}

func main() {
	wf.Run(run)
}

func run() {
	query = strings.TrimSpace(query)
	log.Printf("Running with query: %q", query)

	if query == "" {
		wf.NewItem("Search for a GCP Service ...").
			Subtitle("e.g., sql, redis, gke ...").
			Valid(false)

		if active := gcloud.GetActiveConfig(wf); active != nil {
			wf.NewItem("✅ Using config: " + active.Name).
				Subtitle("This is your currently active gcloud config").
				Valid(false)
		}

		wf.SendFeedback()
		return
	}

	svcList, err := services.LoadServices(servicesFile)
	if err != nil {
		wf.NewItem("Error loading services").
			Subtitle(err.Error()).
			Valid(false)
		wf.SendFeedback()
		return
	}

	active := gcloud.GetActiveConfig(wf)
	parsedQuery := parser.Parse(query, svcList)

	if parsedQuery.SubService != nil {
		key := parsedQuery.Service.ID + "_" + parsedQuery.SubService.ID
		if searcher := searchers.SubserviceSearchers[key]; searcher != nil {
			err := searcher.Search(wf, active, parsedQuery.Filter)
			if err != nil {
				wf.NewItem("Error executing handler").
					Subtitle(err.Error()).
					Valid(false)
			}
			wf.SendFeedback()
			return
		}
	}

	if parsedQuery.Service != nil {
		for _, sub := range parsedQuery.Service.SubServices {
			url, err := RenderURL(sub.URL, active)
			if err != nil {
				log.Println("LOG: error rendering subservice URL:", err)
			}
			wf.NewItem(sub.Name).
				Subtitle(sub.Description).
				Autocomplete(parsedQuery.Service.ID + " " + sub.ID).
				Arg(url).
				Valid(true)
		}
		wf.SendFeedback()
		return
	}

	for _, svc := range svcList {
		url, err := RenderURL(svc.URL, active)
		if err != nil {
			log.Println("LOG: error rendering URL:", err)
		}
		wf.NewItem(svc.Name).
			Autocomplete(svc.ID).
			Subtitle(svc.Description).
			Arg(url).
			Valid(true)
	}
	wf.Filter(query)
	wf.SendFeedback()
}

func RenderURL(tmpl string, cfg *gcloud.Config) (string, error) {
	t, err := template.New("url").Parse(tmpl)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	if err := t.Execute(&buf, cfg); err != nil {
		return "", err
	}
	return buf.String(), nil
}
