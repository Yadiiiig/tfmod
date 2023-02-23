package parser

import (
	"fmt"
	"strings"
)

type prs struct {
	Tokens    []Token
	LenTokens int
	Content   []byte
}

func Parse(tokens []Token, content []byte) error {
	p := prs{tokens, len(tokens), content}

	for k, v := range tokens {
		switch v.Type {
		case TOKEN_RESOURCE:
			fmt.Println("Found resource decleration")
			err := p.parse_resource(k)
			if err != nil {
				return err
			}
		}

	}

	return nil
}

func (p *prs) prev_newline(start, index int) (int, error) {
	for i := index; i >= start; i-- {
		if p.Tokens[i].Type == TOKEN_NEWLINE {
			return i, nil
		}
	}

	return 0, fmt.Errorf("could not find newline")
}

func (p *prs) find_bracket_combo(start int) (int, int, bool) {
	found := 0
	start_index := 0

	for i := start; i <= p.LenTokens; i++ {
		if p.Tokens[i].Type == TOKEN_LEFT_C_BRACKET {
			if found == 0 {
				start_index = i
			}

			found++
		} else if p.Tokens[i].Type == TOKEN_RIGHT_C_BRACKET {
			found--

			if found == 0 {
				return start_index + start, i + start, true
			}
		}

	}

	return 0, 0, false
}

// returns start index, token index, err
func (p *prs) find_string_end(start int) (int, int, error) {
	for i := start + 1; i <= p.LenTokens; i++ {
		if p.Tokens[i].Type == TOKEN_QUOTE {
			return p.Tokens[i].Ind.Start, i, nil
		}
	}

	return 0, 0, fmt.Errorf("could not find string ending")
}

func (p *prs) get_string(start, end int, rm bool) string {
	if rm {
		return strings.ReplaceAll(string(p.Content[start:end]), " ", "")
	}

	return string(p.Content[start:end])
}
