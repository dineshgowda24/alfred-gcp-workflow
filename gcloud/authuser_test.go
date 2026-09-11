package gcloud

import (
	"testing"

	"github.com/dineshgowda24/alfred-gcp-workflow/workflow/env"
)

func TestWithAuthUser(t *testing.T) {
	cfg := &Config{Name: "work", Project: "my-proj", Account: "me@example.com"}

	tests := []struct {
		name   string
		env    string
		config *Config
		in     string
		want   string
	}{
		{
			name: "unset env leaves url untouched",
			env:  "", config: cfg,
			in:   "https://console.cloud.google.com/sql/instances?project=my-proj",
			want: "https://console.cloud.google.com/sql/instances?project=my-proj",
		},
		{
			name: "literal email is appended",
			env:  "other@example.com", config: cfg,
			in:   "https://console.cloud.google.com/sql/instances?project=my-proj",
			want: "https://console.cloud.google.com/sql/instances?authuser=other%40example.com&project=my-proj",
		},
		{
			name: "numeric index is appended",
			env:  "2", config: cfg,
			in:   "https://console.cloud.google.com",
			want: "https://console.cloud.google.com?authuser=2",
		},
		{
			name: "gcloud keyword uses config account",
			env:  "gcloud", config: cfg,
			in:   "https://console.cloud.google.com/storage/browser",
			want: "https://console.cloud.google.com/storage/browser?authuser=me%40example.com",
		},
		{
			name: "gcloud keyword is case-insensitive",
			env:  "GCloud", config: cfg,
			in:   "https://console.cloud.google.com/storage/browser",
			want: "https://console.cloud.google.com/storage/browser?authuser=me%40example.com",
		},
		{
			name: "gcloud keyword with nil config is a no-op",
			env:  "gcloud", config: nil,
			in:   "https://console.cloud.google.com",
			want: "https://console.cloud.google.com",
		},
		{
			name: "gcloud keyword with config lacking account is a no-op",
			env:  "gcloud", config: &Config{Project: "p"},
			in:   "https://console.cloud.google.com",
			want: "https://console.cloud.google.com",
		},
		{
			name: "existing authuser is overridden",
			env:  "1", config: cfg,
			in:   "https://console.cloud.google.com/run?authuser=0&project=p",
			want: "https://console.cloud.google.com/run?authuser=1&project=p",
		},
		{
			name: "fragment is preserved",
			env:  "1", config: cfg,
			in:   "https://console.cloud.google.com/run?project=p#tab",
			want: "https://console.cloud.google.com/run?authuser=1&project=p#tab",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv(env.AuthUserEnv, tt.env)
			if got := WithAuthUser(tt.in, tt.config); got != tt.want {
				t.Errorf("WithAuthUser() = %q, want %q", got, tt.want)
			}
		})
	}
}
