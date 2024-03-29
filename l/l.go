// Package l implements an L-System
// Currently this is a DOL system:
// the left-hand side of a production can only be a single letter; and
// no two productions can have the same left-hand side.
package l

import (
	"container/list"
	"fmt"
	"time"

	"github.com/timtadh/lexmachine"
)

// Rules define all the productions
type Rules map[string]string

// NewRules create new production rules
func NewRules() Rules {
	return Rules(make(map[string]string))
}

// Add adds a new rule
func (r Rules) Add(from string, to string) {
	r[from] = to
}

// Get returns the value at key, ok = false if value doesn't exist
func (r Rules) Get(key string) (value string, ok bool) {
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
	lexer *lexmachine.Lexer
}

// NewSystem returns a new system
func NewSystem(axiom string, rules Rules, lexer *lexmachine.Lexer) *System {
	s := &System{
		axiom: axiom,
		rules: rules,
		state: list.New(),
		lexer: lexer,
	}

	for _, i := range axiom {
		s.state.PushBack(string(i))
	}

	return s
}

// tokenize tokenizes the given string into a list of string
func (s *System) tokenize(in string) []string {
	var result []string
	// simple for now
	// TODO: Implement using https://blog.gopheracademy.com/advent-2017/lexmachine-advent/
	// for _, i := range in {
	// 	result = append(result, string(i))
	// }

	// lexer
	scanner, err := s.lexer.Scanner([]byte(in))
	if err != nil {
		panic(err)
	}

	for tk, err, eof := scanner.Next(); !eof; tk, err, eof = scanner.Next() {
		if err != nil {
			panic(err)
		}
		token := tk.(*lexmachine.Token)
		value := token.Value.(string)
		result = append(result, value)

		// fmt.Printf("%-7v | %-25q | %v:%v-%v:%v\n",
		// 	tokens[token.Type],
		// 	value,
		// 	token.StartLine,
		// 	token.StartColumn,
		// 	token.EndLine,
		// 	token.EndColumn)
	}

	return result
}

// Step applies the rules once
func (s *System) Step(delay time.Duration) {
	var next *list.Element

	for e := s.state.Front(); e != nil; e = next {
		i := e.Value.(string)

		if v, ok := s.rules.Get(i); ok {
			// expand
			for _, j := range s.tokenize(v) {
				if j != "" {
					s.state.InsertBefore(j, e)
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
		i := e.Value.(string)
		result = result + string(i)
	}
	return fmt.Sprintf("%v", result)
}
