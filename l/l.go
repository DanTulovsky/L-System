// Package l implements an L-System
// Currently this is a DOL system:
// the left-hand side of a production can only be a single letter; and
// no two productions can have the same left-hand side.
package l

import (
	"container/list"
	"fmt"
	"time"
)

// Rules define all the productions
type Rules map[rune]string

// NewRules create new production rules
func NewRules() Rules {
	return Rules(make(map[rune]string))
}

// Add adds a new rule
func (r Rules) Add(from rune, to string) {
	r[from] = to
}

// Get returns the value at key, ok = false if value doesn't exist
func (r Rules) Get(key rune) (value string, ok bool) {
	if v, ok := r[key]; ok {
		return v, true
	}
	return "", false
}

// System defines an L-System
type System struct {
	axiom string
	rules Rules
	state *list.List
}

// NewSystem returns a new system
func NewSystem(axiom string, rules Rules) *System {
	s := &System{
		axiom: axiom,
		rules: rules,
		state: list.New(),
	}

	for _, i := range axiom {
		s.state.PushBack(i)
	}

	return s
}

// Step applies the rules once
func (s *System) Step(delay time.Duration) {
	var next *list.Element

	for e := s.state.Front(); e != nil; e = next {
		i := e.Value.(rune)

		if v, ok := s.rules.Get(i); ok {
			// expand
			for j := 0; j < len(v); j++ {
				if string(v[j]) != "" {
					s.state.InsertBefore(rune(v[j]), e)
				}
			}
			next = e.Next()
			s.state.Remove(e)
		} else {
			next = e.Next()
		}
		time.Sleep(delay)
		e = next
	}
}

// State returns the current state of the system
func (s *System) State() *list.List {
	return s.state
}

// String returns state as string
func (s *System) String() string {
	var result string
	for e := s.state.Front(); e != nil; e = e.Next() {
		i := e.Value.(rune)
		result = result + string(i)
	}
	return fmt.Sprintf("%v", result)
}
