package gcloud

import (
	"net/url"
	"strings"

	"github.com/dineshgowda24/alfred-gcp-workflow/workflow/env"
)

const (
	authUserParam = "authuser"

	// AuthUserFromGCloud is the sentinel env value that tells the workflow to
	// use the account of the selected gcloud configuration as the authuser.
	AuthUserFromGCloud = "gcloud"
)

// AuthUser resolves the value to use for the `authuser` query parameter on
// console URLs.
//
// Resolution order:
//  1. ALFRED_GCP_WORKFLOW_AUTHUSER is unset/empty  -> "" (no parameter added)
//  2. ALFRED_GCP_WORKFLOW_AUTHUSER == "gcloud"     -> core/account of the config
//  3. anything else                                -> used verbatim (email or index)
func AuthUser(config *Config) string {
	val := env.AuthUser()
	if val == "" {
		return ""
	}

	if strings.EqualFold(val, AuthUserFromGCloud) {
		if config == nil {
			return ""
		}
		return strings.TrimSpace(config.Account)
	}

	return val
}

// WithAuthUser appends `authuser=<value>` to rawURL when an authuser is
// configured. Existing query parameters and fragments are preserved, and an
// existing authuser parameter is overridden.
func WithAuthUser(rawURL string, config *Config) string {
	user := AuthUser(config)
	if user == "" || rawURL == "" {
		return rawURL
	}

	u, err := url.Parse(rawURL)
	if err != nil {
		return rawURL
	}

	q := u.Query()
	q.Set(authUserParam, user)
	u.RawQuery = q.Encode()
	return u.String()
}
