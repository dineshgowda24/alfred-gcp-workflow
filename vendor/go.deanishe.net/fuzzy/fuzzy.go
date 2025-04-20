// Copyright (c) 2018 Dean Jackson <deanishe@deanishe.net>
// MIT Licence - http://opensource.org/licenses/MIT

package fuzzy // import "go.deanishe.net/fuzzy"

import (
	"sort"
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// Default bonuses and penalties for fuzzy sorting. To customise
// sorting behaviour, pass corresponding Options to New() or
// Sorter.Configure().
const (
	DefaultAdjacencyBonus          = 5.0  // Bonus for adjacent matches
	DefaultSeparatorBonus          = 10.0 // Bonus if the match is after a separator
	DefaultCamelBonus              = 10.0 // Bonus if match is uppercase and previous is lower
	DefaultLeadingLetterPenalty    = -3.0 // Penalty applied for every letter in string before first match
	DefaultMaxLeadingLetterPenalty = -9.0 // Maximum penalty for leading letters
	DefaultUnmatchedLetterPenalty  = -1.0 // Penalty for every letter that doesn't match
	DefaultStripDiacritics         = true // Strip diacritics from sort keys if query is plain ASCII
)

// Sortable makes the implementer fuzzy-sortable.
// It is a superset of sort.Interface (i.e. your struct must also
// implement sort.Interface).
type Sortable interface {
	sort.Interface
	// Keywords returns the string to compare to the sort query
	Keywords(i int) string
}

// Result stores the result of a single fuzzy ranking.
type Result struct {
	// Match is whether or not the string matched the query,
	// i.e. if all characters in the query are present,
	// in order, in the string.
	Match bool
	// Query is the query that was matched against.
	Query string
	// Score is how well the string matched the query.
	// Higher is better.
	Score float64
	// SortKey is the string Query was compared to.
	SortKey string
}

// Sorter sorts Data based on the query passsed to Sorter.Sort().
type Sorter struct {
	Data                    Sortable              // Data to sort
	AdjacencyBonus          float64               // Bonus for adjacent matches
	SeparatorBonus          float64               // Bonus if the match is after a separator
	CamelBonus              float64               // Bonus if match is uppercase and previous is lower
	LeadingLetterPenalty    float64               // Penalty applied for every letter in string before first match
	MaxLeadingLetterPenalty float64               // Maximum penalty for leading letters
	UnmatchedLetterPenalty  float64               // Penalty for every letter that doesn't match
	StripDiacritics         bool                  // Strip diacritics from sort keys if query is plain ASCII
	stripDiacritics         bool                  // Internal setting based on StripDiacritics and whether query is plain ASCII
	stripper                transform.Transformer // Diacritics stripper
	query                   string                // Search query
	results                 []*Result             // Results of the fuzzy sort
}

// New creates a new Sorter for the given data.
func New(data Sortable, opts ...Option) *Sorter {
	s := &Sorter{
		Data:                    data,
		AdjacencyBonus:          DefaultAdjacencyBonus,
		SeparatorBonus:          DefaultSeparatorBonus,
		CamelBonus:              DefaultCamelBonus,
		LeadingLetterPenalty:    DefaultLeadingLetterPenalty,
		MaxLeadingLetterPenalty: DefaultMaxLeadingLetterPenalty,
		UnmatchedLetterPenalty:  DefaultUnmatchedLetterPenalty,
		StripDiacritics:         DefaultStripDiacritics,
		stripDiacritics:         false,
		stripper:                transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn))),
		results:                 make([]*Result, data.Len()),
	}
	s.Configure(opts...)
	return s
}

// Configure applies one or more Options to Sorter.
func (s *Sorter) Configure(opts ...Option) Option {
	var undo Option
	for _, opt := range opts {
		undo = opt(s)
	}
	return undo
}

// Len implements sort.Interface.
func (s *Sorter) Len() int { return s.Data.Len() }

// Less implements sort.Interface.
func (s *Sorter) Less(i, j int) bool {
	a, b := s.results[i], s.results[j]

	// Matches beat non-matches.
	if a.Match && !b.Match {
		return true
	}
	if !a.Match && b.Match {
		return false
	}

	if a.Score == b.Score {
		// Normal comparison because A comes before B.
		return s.Data.Less(i, j)
	}
	// Reverse comparison because higher score is better.
	return b.Score < a.Score
}

// Swap implements sort.Interface.
func (s *Sorter) Swap(i, j int) {
	s.results[i], s.results[j] = s.results[j], s.results[i]
	s.Data.Swap(i, j)
}

// Sort sorts data against query.
func (s *Sorter) Sort(query string) []*Result {
	s.query = query

	if s.isASCII(query) && s.StripDiacritics {
		s.stripDiacritics = true
	}

	for i := 0; i < s.Data.Len(); i++ {
		s.results[i] = s.Match(s.Data.Keywords(i))
	}

	sort.Sort(s)

	return s.results
}

// Match scores str against Sorter's query using fuzzy matching.
func (s *Sorter) Match(str string) *Result {
	if s.stripDiacritics {
		str = s.strip(str)
	}

	var (
		match    = false           // Whether or not str matches query
		score    = 0.0             // How well str matches query
		uStr     = []rune(str)     // str as slice of Unicode chars
		uQuery   = []rune(s.query) // query as slice of Unicode chars
		strLen   = len(uStr)       // Length of Unicode str
		queryLen = len(uQuery)     // Length of Unicode query
	)
	var (
		queryIdx, strIdx                   int
		newScore, penalty, bestLetterScore float64
		queryChar, queryLower              string
		strChar, strLower, strUpper        string
		bestLetter, bestLower              string
		advanced, queryRepeat              bool
		nextMatch, rematch                 bool
		prevMatched, prevLower             bool
		prevSeparator                      = true
	)

	// Loop through each character in str
	for strIdx != strLen {
		strChar = string(uStr[strIdx])

		if queryIdx != queryLen {
			queryChar = string(uQuery[queryIdx])
		} else {
			queryChar = ""
		}

		queryLower = strings.ToLower(queryChar)
		strLower = strings.ToLower(strChar)
		strUpper = strings.ToUpper(strChar)

		if queryChar != "" && queryLower == strLower {
			nextMatch = true
		} else {
			nextMatch = false
		}
		if bestLetter != "" && bestLower == strLower {
			rematch = true
		} else {
			rematch = false
		}

		if nextMatch && bestLetter != "" {
			advanced = true
		} else {
			advanced = false
		}

		if bestLetter != "" && strChar != "" && bestLower == queryLower {
			queryRepeat = true
		} else {
			queryRepeat = false
		}

		if advanced || queryRepeat {
			score += bestLetterScore
			bestLetter = ""
			bestLower = ""
			bestLetterScore = 0.0
		}

		if nextMatch || rematch {
			newScore = 0.0

			// Apply penalty for letters before first match
			if queryIdx == 0 {
				penalty = float64(strIdx) * s.LeadingLetterPenalty
				if penalty <= s.MaxLeadingLetterPenalty {
					penalty = s.MaxLeadingLetterPenalty
				}
				score += penalty
			}

			// Apply bonus for consecutive matches
			if prevMatched {
				newScore += s.AdjacencyBonus
			}

			// Apply bonus for match after separator
			if prevSeparator {
				newScore += s.SeparatorBonus
			}

			// Apply bonus across camel case boundaries
			if prevLower && strChar == strUpper && strLower != strUpper {
				newScore += s.CamelBonus
			}

			// Update query index if next query letter was matched
			if nextMatch {
				queryIdx++
			}

			// Update best letter in key, which may be for a "next" letter or a reMatch
			if newScore >= bestLetterScore {
				if bestLetter != "" {
					score += s.UnmatchedLetterPenalty
				}

				bestLetter = strChar
				bestLower = strings.ToLower(bestLetter)
				bestLetterScore = newScore
			}

			prevMatched = true
		} else {
			score += s.UnmatchedLetterPenalty
			prevMatched = false
		}

		// IsLetter check
		if strChar == strLower && strLower != strUpper {
			prevLower = true
		} else {
			prevLower = false
		}
		if strChar == "_" || strChar == " " || strChar == "." || strChar == "-" || strChar == "/" {
			prevSeparator = true
		} else {
			prevSeparator = false
		}

		strIdx++
	}

	if bestLetter != "" {
		score += bestLetterScore
	}

	if queryIdx == queryLen {
		match = true
	}

	return &Result{match, s.query, score, str}
}

func (s *Sorter) strip(str string) string {
	if stripped, _, err := transform.String(s.stripper, str); err == nil {
		return stripped
	}
	return str
}

func (s *Sorter) isASCII(str string) bool { return s.strip(str) == str }

// Sort sorts data against query using a new default Sorter.
func Sort(data Sortable, query string) []*Result {
	return New(data).Sort(query)
}

// SortStrings fuzzy-sorts a slice of strings.
func SortStrings(data []string, query string) []*Result {
	return strSlice(data).Sort(query)
}

// Match scores str against query using the specified sort options.
//
// WARNING: Match creates a new Sorter for every call.
// Don't use this on large datasets.
func Match(str, query string, opts ...Option) *Result {
	return New(strSlice([]string{str}), opts...).Sort(query)[0]
}

// strSlice implements Sortable for []string.
// It is a helper for SortStrings.
type strSlice []string

// Len etc. implement sort.Interface.
func (s strSlice) Len() int           { return len(s) }
func (s strSlice) Less(i, j int) bool { return s[i] < s[j] }
func (s strSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// Keywords implements Sortable.
func (s strSlice) Keywords(i int) string { return s[i] }

// Sort is a convenience method.
func (s strSlice) Sort(query string) []*Result {
	return Sort(s, query)
}
