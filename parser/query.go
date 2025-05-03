package parser

import (
	"log"
	"strings"

	"github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
	"github.com/dineshgowda24/alfred-gcp-workflow/services"
	"github.com/dineshgowda24/alfred-gcp-workflow/workflow/arg"
)

type Result struct {
	SearchArgs     *arg.SearchArgs
	Service        *services.Service
	SubService     *services.Service
	RemainingQuery string

	IsConfigQuery      bool
	Config             *gcloud.Config
	PartialConfigQuery string

	IsRegionQuery      bool
	Region             *gcloud.Region
	PartialRegionQuery string

	Intent      string
	IntentValue string
}

func (r *Result) IsEmptyQuery() bool {
	return strings.TrimSpace(r.SearchArgs.Query) == ""
}

func (r *Result) HasServiceOnly() bool {
	return r.Service != nil && r.SubService == nil
}

func (r *Result) HasSubService() bool {
	return r.Service != nil && r.SubService != nil
}

func (r *Result) HasIntent() bool {
	return r.Intent != ""
}

func Parse(searchArgs *arg.SearchArgs, svcList []services.Service) *Result {
	query := strings.TrimSpace(searchArgs.Query)
	words := strings.Fields(query)
	result := &Result{SearchArgs: searchArgs}

	if len(words) == 0 {
		return result
	}

	if result.extractIntent(words) {
		return result
	}

	filtered := result.extractConfig(words)
	filtered = result.extractRegion(filtered)

	if len(filtered) == 0 {
		return result
	}

	result.extractServiceAndSubService(filtered, svcList)
	return result
}

func (r *Result) extractIntent(words []string) bool {
	if len(words) == 0 {
		return false
	}

	intent := strings.ToLower(words[0])
	switch intent {
	case "gcloud-path":
		r.Intent = intent
		r.IntentValue = strings.Join(words[1:], " ")
		return true
	default:
		return false
	}
}

func (r *Result) extractConfig(words []string) []string {
	var filtered []string
	for _, w := range words {
		if strings.HasPrefix(w, "@") {
			r.IsConfigQuery = true
			r.matchConfig(strings.TrimPrefix(w, "@"))
		} else {
			filtered = append(filtered, strings.ToLower(w))
		}
	}
	return filtered
}

func (r *Result) matchConfig(name string) {
	configs, err := gcloud.GetAllConfigs()
	if err != nil {
		log.Println("error fetching configs:", err)
		r.PartialConfigQuery = name
		return
	}

	for _, c := range configs {
		if strings.EqualFold(c.Name, name) {
			r.Config = c
			return
		}
	}

	r.PartialConfigQuery = name
}

func (r *Result) extractRegion(words []string) []string {
	var filtered []string
	for _, w := range words {
		if strings.HasPrefix(w, "$") {
			r.IsRegionQuery = true
			r.matchRegion(strings.TrimPrefix(w, "$"))
		} else {
			filtered = append(filtered, strings.ToLower(w))
		}
	}

	return filtered
}

func (r *Result) matchRegion(name string) {
	regions := gcloud.GetAllRegions()
	for _, reg := range regions {
		if strings.EqualFold(reg.Name, name) {
			r.Region = &reg
			return
		}
	}

	r.PartialRegionQuery = name
}

func (r *Result) extractServiceAndSubService(words []string, svcList []services.Service) {
	serviceMap := buildServiceMap(svcList)

	service, ok := serviceMap[words[0]]
	if !ok {
		r.RemainingQuery = strings.Join(words, " ")
		return
	}
	r.Service = service

	if len(words) >= 2 {
		subMap := buildServiceMap(service.SubServices)
		sub, ok := subMap[words[1]]
		if ok {
			r.SubService = sub
			if len(words) > 2 {
				r.RemainingQuery = strings.Join(words[2:], " ")
			}
		} else {
			r.RemainingQuery = strings.Join(words[1:], " ")
		}
	}
}

func buildServiceMap(svcList []services.Service) map[string]*services.Service {
	m := make(map[string]*services.Service)
	for i := range svcList {
		m[strings.ToLower(svcList[i].ID)] = &svcList[i]
	}
	return m
}
