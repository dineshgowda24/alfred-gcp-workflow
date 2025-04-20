package main

import (
	"embed"
	"flag"
	"log"
	"strings"

	aw "github.com/deanishe/awgo"
	"github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
	"github.com/dineshgowda24/alfred-gcp-workflow/parser"
	"github.com/dineshgowda24/alfred-gcp-workflow/searchers"
	"github.com/dineshgowda24/alfred-gcp-workflow/services"
)

var (
	wf           *aw.Workflow
	query        string
	forceFetch   bool
	activeConfig *gcloud.Config
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
		wf.NewItem("🔍 Search GCP Services").
			Subtitle("Try keywords like: sql, redis, storage").
			Icon(aw.IconHelp).
			Valid(false)

		wf.NewItem("🧭 Open Console").
			Subtitle("Go to https://console.cloud.google.com/").
			Icon(aw.IconHome).
			Arg("https://console.cloud.google.com/").
			Valid(true)

		if activeConfig := getActiveConfig(wf); activeConfig != nil {
			wf.NewItem("☁️ Active gcloud config: " + activeConfig.Name).
				Subtitle("Currently selected gcloud configuration").
				Icon(aw.IconInfo).
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

	active := getActiveConfig(wf)
	parsedQuery := parser.Parse(query, svcList)

	if parsedQuery.SubService != nil {
		key := parsedQuery.Service.ID + "_" + parsedQuery.SubService.ID
		if searcher := searchers.SubserviceSearchers[key]; searcher != nil {
			err := searcher.Search(wf, parsedQuery.SubService, active, parsedQuery.Filter)
			if err != nil {
				wf.NewItem("Error executing handler").
					Subtitle(err.Error()).
					Valid(false)
			}

			wf.Filter(parsedQuery.Filter)
			wf.SendFeedback()
			return
		}
	}

	if parsedQuery.Service != nil {
		for _, sub := range parsedQuery.Service.SubServices {
			url, err := sub.Url(active)
			if err != nil {
				log.Println("LOG: error rendering subservice URL:", err)
			}
			wf.NewItem(sub.Name).
				Subtitle(sub.Description).
				Autocomplete(parsedQuery.Service.ID + " " + sub.ID).
				Icon(sub.Icon()).
				Arg(url).
				Valid(true)
		}
		wf.SendFeedback()
		return
	}

	for _, svc := range svcList {
		url, err := svc.Url(active)
		if err != nil {
			log.Println("LOG: error rendering URL:", err)
		}
		wf.NewItem(svc.Name).
			Autocomplete(svc.ID).
			Subtitle(svc.Description).
			Arg(url).
			Icon(svc.Icon()).
			Valid(true)
	}
	wf.Filter(query)
	wf.SendFeedback()
}

func getActiveConfig(wf *aw.Workflow) *gcloud.Config {
	if activeConfig != nil {
		return activeConfig
	}

	activeConfig = gcloud.GetActiveConfig(wf)
	if activeConfig == nil {
		wf.NewItem("No active gcloud config found").
			Subtitle("Please set up a gcloud config").
			Icon(aw.IconWarning).
			Valid(false)
		wf.SendFeedback()
		return nil
	}

	return activeConfig
}
