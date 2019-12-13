package l

import (
	"log"

	"github.com/timtadh/lexmachine"
	"github.com/timtadh/lexmachine/machines"
)

var tokens = []string{
	"MOVE_DRAW", "MOVE", "TURN_RIGHT", "TURN_LEFT", "SCALE", "PUSH_STATE", "POP_STATE", "COLOR", "COLOR_AFTER", "COLOR_BEFORE",
}

var tokmap map[string]int

func getToken(tokenType int) lexmachine.Action {
	return func(s *lexmachine.Scanner, m *machines.Match) (interface{}, error) {
		return s.Token(tokenType, string(m.Bytes), m), nil
	}
}

// NewDefaultLexer returns the default lexer
func NewDefaultLexer(rules Rules) *lexmachine.Lexer {
	log.Println("Initializing default lexer...")
	tokmap = make(map[string]int)

	lexer := lexmachine.NewLexer()
	// tokens that cause give commands to the turtle
	for key := range rules {
		tokens = append(tokens, key)
	}
	for id, name := range tokens {
		tokmap[name] = id
	}
	lexer.Add([]byte(`F`), getToken(tokmap["MOVE_DRAW"]))
	lexer.Add([]byte(`G`), getToken(tokmap["MOVE"]))
	lexer.Add([]byte(`-`), getToken(tokmap["TURN_RIGHT"]))
	lexer.Add([]byte(`\+`), getToken(tokmap["TURN_LEFT"]))
	// fix regex later
	lexer.Add([]byte(`\@(I|Q)*0*(\.)*\d*`), getToken(tokmap["SCALE"]))
	lexer.Add([]byte(`\[`), getToken(tokmap["PUSH_STATE"]))
	lexer.Add([]byte(`\]`), getToken(tokmap["POP_STATE"]))
	lexer.Add([]byte(`\%\d*`), getToken(tokmap["COLOR"]))
	lexer.Add([]byte(`\>\d*`), getToken(tokmap["COLOR_AFTER"]))
	lexer.Add([]byte(`\<\d*`), getToken(tokmap["COLOR_BEFORE"]))

	// tokens present in the rules

	for key := range rules {
		tokens = append(tokens, key)
		lexer.Add([]byte(key), getToken(tokmap[key]))

	}

	err := lexer.Compile()
	if err != nil {
		panic(err)
	}
	return lexer
}
