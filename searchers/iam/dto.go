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

type Account struct {
	DisplayName string
	Email       string
	UniqueID    string
}

func (a Account) Title() string {
	return a.DisplayName
}

func (a Account) Subtitle() string {
	return a.Email
}

func (a Account) UID() string {
	return a.UniqueID
}

func (a Account) URL(config *gcloud.Config) string {
	return "https://console.cloud.google.com/iam-admin/serviceaccounts/details/" + a.UID() + "?project=" + config.Project
}

func FromGCloudIAMServiceAccount(accounts *gcloud.IAMServiceAccount) Account {
	return Account{
		DisplayName: accounts.DisplayName,
		Email:       accounts.Email,
		UniqueID:    accounts.UniqueID,
	}
}
