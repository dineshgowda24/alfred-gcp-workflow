fuzzy
=====

![Build Status][github-status-icon]
[![Go Report Card][goreport-icon]][goreport-link]
[![Codacy coverage][coverage-icon]][codacy-link]
[![GitHub licence][licence-icon]][licence-link]
[![GoDoc][godoc-icon]][godoc-link]

Package fuzzy implements fuzzy matching/sorting of string slices and custom types (via `fuzzy.Sortable` interface).

Import path is `go.deanishe.net/fuzzy`

The fuzzy matching algorithm is based on [Forrest Smith's reverse engineering of Sublime Text's search][forrest].


Features
--------

- Sublime Text-like fuzzy matching
- Simple sorting of string slices via `fuzzy.SortStrings()`
- Sorting of custom types via `fuzzy.Sortable` interface
- Intelligent handling of diacritics


Usage
-----

See [Godoc][godoc] for full documentation.


### Slice of strings ###

```go
actors := []string{"Tommy Lee Jones", "James Earl Jones", "Keanu Reeves"}
fuzzy.SortStrings(actors, "jej")
fmt.Println(actors[0])
// -> James Earl Jones
```


### Custom types ###

To sort a custom type, it must implement `fuzzy.Sortable`:

```go
type Sortable interface {
    sort.Interface
    // Keywords returns the string to compare to the sort query
    Keywords(i int) string
}
```

This is a superset of `sort.Interface` (i.e. your type must also implement `sort.Interface`). The string returned by `Keywords()` is compared to the search query using the fuzzy algorithm. The `Less()` function of `sort.Interface` is used as a tie-breaker if multiple items have the same fuzzy matching score.

```go
// Player is a very simple data model.
type Player struct {
    Firstname string
    Lastname  string
}

// Name returns the full name of the Player.
func (p *Player) Name() string {
    return strings.TrimSpace(p.Firstname + " " + p.Lastname)
}

// Team is a collection of Player items. This is where fuzzy.Sortable
// must be implemented to enable fuzzy sorting.
type Team []*Player

// Default sort.Interface methods
func (t Team) Len() int      { return len(t) }
func (t Team) Swap(i, j int) { t[i], t[j] = t[j], t[i] }

// Less is used as a tie-breaker when fuzzy match score is the same.
func (t Team) Less(i, j int) bool { return t[i].Name() < t[j].Name() }

// Keywords implements Sortable.
// Comparisons are based on the the full name of the player.
func (t Team) Keywords(i int) string { return t[i].Name() }

// Fuzzy sort players by name.
var t = Team{
    &Player{"Alisson", "Becker"},
    &Player{Firstname: "Adrián"},
    &Player{"Andy", "Lonergan"},
    &Player{"Caoimhín", "Kelleher"},
    &Player{"Virgil", "van Dijk"},
    &Player{"Joe", "Gomez"},
    &Player{"Andy", "Robertson"},
    &Player{"Joel", "Matip"},
    &Player{"Ki-Jana", "Hoever"},
    &Player{"Trent", "Alexander-Arnold"},
    &Player{"Sepp", "van den Berg"},
    &Player{"Neco", "Williams"},
    &Player{Firstname: "Fabinho"},
    &Player{"Georginio", "Wijnaldum"},
    &Player{"James", "Milner"},
    &Player{"Naby", "Keita"},
    &Player{"Jordan", "Henderson"},
    &Player{"Alex", "Oxlade-Chamberlain"},
    &Player{"Xherdan", "Shaqiri"},
    &Player{"Curtis", "Jones"},
    &Player{"Harvel", "Elliott"},
    &Player{"Roberto", "Firmino"},
    &Player{"Sadio", "Mané"},
    &Player{"Mohamed", "Salah"},
    &Player{"Takumi", "Minamino"},
    &Player{"Divock", "Origi"},
}
// Unsorted
fmt.Println(t[0].Name()) // --> Alisson Becker

// Initials
Sort(t, "taa")
fmt.Println(t[0].Name()) // --> Trent Alexander-Arnold

// Initials beat start of string
Sort(t, "al")
fmt.Println(t[0].Name()) // --> Andy Lonergan

// Start of word
Sort(t, "ox")
fmt.Println(t[0].Name()) // --> Alex Oxlade-Chamberlain

// Earlier in string = better match
Sort(t, "x")
fmt.Println(t[0].Name()) // --> Xherdan Shaqiri

// Diacritics are considered
Sort(t, "ne")
fmt.Println(t[0].Name()) // --> Neco Williams
Sort(t, "né")
fmt.Println(t[0].Name()) // --> Sadio Mané

// But ignored if query is ASCII
Sort(t, "mane")
fmt.Println(t[0].Name()) // --> Sadio Mané
```


Licence
-------

`fuzzy` is released under the [MIT licence][mit].

[godoc]: https://godoc.org/go.deanishe.net/fuzzy
[forrest]: https://blog.forrestthewoods.com/reverse-engineering-sublime-text-s-fuzzy-match-4cffeed33fdb
[mit]: ./LICENCE.txt
[godoc-icon]: https://godoc.org/go.deanishe.net/fuzzy?status.svg
[godoc-link]: https://godoc.org/go.deanishe.net/fuzzy
[goreport-link]: https://goreportcard.com/report/github.com/deanishe/go-fuzzy
[goreport-icon]: https://goreportcard.com/badge/github.com/deanishe/go-fuzzy
[coverage-icon]: https://img.shields.io/codacy/coverage/9cdd179cb6ce4236979ef01915b9e6eb?color=brightgreen
[codacy-link]: https://www.codacy.com/app/deanishe/go-fuzzy
[licence-icon]: https://img.shields.io/github/license/deanishe/go-fuzzy
[licence-link]: https://github.com/deanishe/go-fuzzy/blob/master/LICENCE.txt
[github-status-icon]: https://img.shields.io/github/workflow/status/deanishe/go-fuzzy/Test
