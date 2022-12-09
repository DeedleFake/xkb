package xkb

import "io"

const (
	tokEndOfFile = iota
	tokEndOfLine
	tokInclude
	tokIncludeString
	tokLHSKeysym
	tokColon
	tokBang
	tokTilde
	tokString
	tokIdent
	tokError
)

type (
	lvalue  any
	lstring []byte
)

func peek(r io.RuneScanner) (rune, error) {
	c, _, err := r.ReadRune()
	if err != nil {
		return c, err
	}
	return c, r.UnreadRune()
}
