package orchestrator

import (
	"strings"

	aw "github.com/deanishe/awgo"
	"github.com/dineshgowda24/alfred-gcp-workflow/services"
)

const (
	contributingURL = "https://github.com/dineshgowda24/alfred-gcp-workflow/CONTRIBUTING.md"
)

func buildAutocomplete(ctx *Context, service *services.Service) string {
	query := ctx.Args.Query
	remainingQuery := ctx.ParsedQuery.RemainingQuery

	if !ctx.ParsedQuery.IsConfigQuery && !ctx.ParsedQuery.IsRegionQuery {
		return service.Autocomplete()
	}

	if ctx.IsSubServiceQuery() {
		return query
	}

	// freze the autocomplete if the query is messy
	if isMessyQuery(query, remainingQuery) {
		return query
	}

	prefix := strings.TrimSuffix(query, remainingQuery)
	prefix = strings.TrimSuffix(prefix, " ")

	if prefix != "" {
		return prefix + " " + service.ID
	}
	return query + " " + service.ID
}

func isMessyQuery(query, remaining string) bool {
	queryWords := strings.Fields(query)
	remainingWords := strings.Fields(remaining)

	// should not happen
	if len(remainingWords) > len(queryWords) {
		return true
	}

	start := len(queryWords) - len(remainingWords)
	for i := 0; i < len(remainingWords); i++ {
		if queryWords[start+i] != remainingWords[i] {
			return true
		}
	}

	return false
}

func emptyResultItem(wf *aw.Workflow, title string) {
	wf.NewItem(title).
		Subtitle("Try a different query").
		Icon(aw.IconNote).
		Valid(false)
}

func addContributingItem(wf *aw.Workflow, title string) {
	wf.NewFileItem(title).
		Subtitle("Open contributing guide to add them").
		Arg(contributingURL).
		Icon(aw.IconNote).
		Valid(true)
}
