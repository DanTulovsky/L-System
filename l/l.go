// Package l implements an L-System
// Currently this is a DOL system:
// the left-hand side of a production can only be a single letter; and
// no two productions can have the same left-hand side.
package l

import (
	"fmt"
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
	state string
}

// NewSystem returns a new system
func NewSystem(axiom string, rules Rules) *System {
	s := &System{
		axiom: axiom,
		rules: rules,
		state: axiom,
	}

	return s
}

// Step makes the system take one step (apply the rules once)
func (s *System) Step() {
	var state string

	for _, i := range s.state {
		if v, ok := s.rules.Get(i); ok {
			// expand
			state = state + v
		} else {
			state = state + string(i)
		}
	}

	s.state = state
}

// State returns the current state of the system
func (s *System) State() string {
	return s.state
}

// String returns state as string
func (s *System) String() string {
	return fmt.Sprintf("%v", s.state)
}
