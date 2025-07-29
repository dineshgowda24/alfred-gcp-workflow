package iam

import (
	"strings"

	"github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
)

type Role struct {
	Description  string
	Name         string
	DisplayTitle string
}

func (r Role) Details() string {
	return r.Name
}

func (r Role) Subtitle() string {
	return r.Description
}

func (r Role) Title() string {
	return r.DisplayTitle
}

func (r Role) URL(config *gcloud.Config) string {
	id := r.ID()
	return "https://console.cloud.google.com/iam-admin/roles/details/" + id + "?project=" + config.Project

}

func (r Role) ID() string {
	parts := strings.ReplaceAll(r.Name, "/", "%2F")
	return parts
}
func FromGCloudIAMRoles(roles *gcloud.IAMRole) Role {
	return Role{
		Description:  roles.Description,
		Name:         roles.Name,
		DisplayTitle: roles.Title,
	}
}
