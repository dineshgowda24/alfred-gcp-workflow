package parser

import (
	"log"
	"strings"

	"github.com/dineshgowda24/alfred-gcp-workflow/services"
	"github.com/dineshgowda24/alfred-gcp-workflow/workflow/arg"
)

type Result struct {
	SearchArgs     *arg.SearchArgs
	Service        *services.Service
	SubService     *services.Service
	RemainingQuery string
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

func Parse(searchArgs *arg.SearchArgs, svcList []services.Service) *Result {
	words := strings.Fields(strings.ToLower(strings.TrimSpace(searchArgs.Query)))
	pq := &Result{SearchArgs: searchArgs}

	if len(words) == 0 {
		return pq
	}

	log.Println("LOG: parsing words:", words)

	serviceMap := buildServiceMap(svcList)
	service, ok := serviceMap[words[0]]
	if !ok {
		return pq
	}
	pq.Service = service

	if len(words) >= 2 {
		subMap := buildServiceMap(service.SubServices)
		sub, ok := subMap[words[1]]
		if ok {
			pq.SubService = sub
			if len(words) > 2 {
				pq.RemainingQuery = strings.Join(words[2:], " ")
			}
		} else {
			pq.RemainingQuery = strings.Join(words[1:], " ")
		}
	}

	return pq
}

func buildServiceMap(svcList []services.Service) map[string]*services.Service {
	serviceMap := make(map[string]*services.Service)
	for i := range svcList {
		serviceMap[strings.ToLower(svcList[i].ID)] = &svcList[i]
	}

	return serviceMap
}
