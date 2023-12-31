// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

// Injectors from wire.go:

// NewNaraku returns a new Naraku struct.
func NewNaraku() *Naraku {
	naraku := newNaraku()
	return naraku
}

// wire.go:

// Naraku is a struct that contains the settings for the naraku.
type Naraku struct{}

// newNaraku returns a new naraku struct.
func newNaraku() *Naraku {
	return &Naraku{}
}
