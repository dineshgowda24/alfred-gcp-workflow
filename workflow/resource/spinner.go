package resource

import (
	"log"

	aw "github.com/deanishe/awgo"
)

type spinner struct {
	key    string
	frames []string
	wf     *aw.Workflow
	index  int
	loaded bool
}

func defaultSpinner(wf *aw.Workflow) *spinner {
	frames := []string{"🕐", "🕑", "🕒", "🕓", "🕔", "🕕", "🕖", "🕗", "🕘", "🕙", "🕚", "🕛"}
	return newSpinner(wf, "spinner-index", frames)
}

func newSpinner(wf *aw.Workflow, key string, frames []string) *spinner {
	return &spinner{
		key:    key,
		frames: frames,
		wf:     wf,
	}
}

func (s *spinner) NextFrame() string {
	s.load()
	frame := s.next()
	s.store()
	return frame
}

func (s *spinner) load() {
	if s.loaded {
		return
	}

	if err := s.wf.Cache.LoadJSON(s.key, &s.index); err != nil {
		log.Println("cache read failed:", err)
		s.index = 0
	}
	s.loaded = true
}

func (s *spinner) next() string {
	if len(s.frames) == 0 {
		return "…"
	}

	frame := s.frames[s.index]
	s.index = (s.index + 1) % len(s.frames)

	return frame
}

func (s *spinner) store() {
	if err := s.wf.Cache.StoreJSON(s.key, s.index); err != nil {
		log.Println("cache write failed:", err)
	}
}
