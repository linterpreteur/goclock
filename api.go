package goclock

import (
	"time"
)

type Goclock struct {
	Source string        `json:"source"`
	Offset time.Duration `json:"offset"`
}

type Request struct {
	Url        string
	ClientTime time.Time
}

func New(request Request) (*Goclock, error) {
	g := &Goclock{}
	err := g.initialize(request)
	return g, err
}

func (g *Goclock) Initialize(request Request) error {
	return g.initialize(request)
}

func (g Goclock) Time() time.Time {
	return time.Now().Add(g.Offset)
}

func (g *Goclock) initialize(request Request) error {
	g.Source = request.Url
	o, err := offset(g.Source)
	if err != nil {
		return err
	}
	g.Offset = o
	return nil
}
