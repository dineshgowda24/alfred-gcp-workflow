package parser

import (
	"strings"

	"github.com/dineshgowda24/alfred-gcp-workflow/services"
)

type Query struct {
	RawQuery   string
	Service    *services.Service
	SubService *services.Service
	Filter     string
}

func Parse(input string, svcList []services.Service) *Query {
	words := strings.Fields(strings.ToLower(strings.TrimSpace(input)))
	pq := &Query{RawQuery: input}

	if len(words) == 0 {
		return pq
	}

	serviceMap := make(map[string]*services.Service)
	for i := range svcList {
		serviceMap[strings.ToLower(svcList[i].ID)] = &svcList[i]
	}

	service, ok := serviceMap[words[0]]
	if !ok {
		return pq
	}
	pq.Service = service

	if len(words) >= 2 {
		subMap := make(map[string]*services.Service)
		for i := range service.SubServices {
			subMap[strings.ToLower(service.SubServices[i].ID)] = &service.SubServices[i]
		}
		sub, ok := subMap[words[1]]
		if ok {
			pq.SubService = sub
			if len(words) > 2 {
				pq.Filter = strings.Join(words[2:], " ")
			}
		}
	}

	return pq
}
