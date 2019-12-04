package lexer

type Lexer struct {
	input        string
	position     int  // Current position in input (points to current char)
	readPosition int  // Current reading position in input (after current char)
	ch           byte // Current char under examination
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}
