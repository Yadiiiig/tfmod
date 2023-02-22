package parser

import (
	"bytes"
	"fmt"
	"io"
	"unicode"

	"github.com/yadiiiig/tfmod/models"
)

type lxr struct {
	Token      string
	RuneString string

	Total    int
	InString bool

	Tokens []Token

	Pos Position
}

func Lexer(file *models.File) ([]Token, error) {
	l := lxr{
		Token:    "",
		Total:    0,
		InString: false,
		Tokens:   []Token{},
		Pos: Position{
			Line:   0,
			Column: -1,
		},
	}

	reader := bytes.NewReader(file.Content)

	for {
		l.Total++
		l.Pos.Column++

		r, _, err := reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			}

			return l.Tokens, err
		}

		if unicode.IsSpace(r) {
			if r == '\n' {
				l.add_token(TOKEN_NEWLINE, true, "")

				l.Pos.Line++
				l.Pos.Column = -1
			}

			continue
		}

		l.Token += string(r)

		tkn_type, ok := Keywords[l.Token]
		if ok {
			l.add_token(tkn_type, false, "")
			continue
		}

		l.RuneString = string(r)
		tkn_type, ok = Chars[l.RuneString]
		if ok {
			switch tkn_type {
			case TOKEN_QUOTE:
				if !l.InString {
					l.InString = true
					l.add_token(tkn_type, true, `"`)

					continue
				}

				l.InString = false
				l.add_token(tkn_type, true, `"`)

				continue
			case TOKEN_ASSIGN:
				l.add_token(tkn_type, true, "=")
				continue

			case TOKEN_POINT:
				l.add_token(tkn_type, true, ".")
				continue
			}

			if !l.InString {
				l.add_token(tkn_type, false, "")
			}
		}
	}

	return l.Tokens, nil
}

func (l *lxr) add_token(t int, chk bool, token string) {
	if !chk {
		token = l.Token
	}

	l.Tokens = append(l.Tokens, Token{
		Type:  t,
		Value: token,
		Pos:   l.Pos,
		Ind: Indices{
			Start: (l.Total - len(token)),
			End:   l.Total,
		},
	})

	l.Token = ""

	fmt.Printf("%+v\n", l.Tokens[len(l.Tokens)-1])
}
