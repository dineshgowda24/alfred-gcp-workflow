package iam

import (
	"strings"

	"github.com/dineshgowda24/alfred-gcp-workflow/gcloud"
)

type Roleslist struct {
	Description  string
	Name         string
	DisplayTitle string
}

func (r Roleslist) Details() string {
	return r.Name
}

func (r Roleslist) Subtitle() string {
	return r.Description
}

func (r Roleslist) Title() string {
	return r.DisplayTitle
}

func (r Roleslist) URL(config *gcloud.Config) string {
	id := r.ID()
	return "https://console.cloud.google.com/iam-admin/roles/details/" + id + "?project=" + config.Project

}

func (r Roleslist) ID() string {
	parts := strings.ReplaceAll(r.Name, "/", "%2F")
	return parts
}
func FromGCloudIAMRoles(roles *gcloud.IAMRole) Roleslist {
	return Roleslist{
		Description:  roles.Description,
		Name:         roles.Name,
		DisplayTitle: roles.Title,
	}
}
