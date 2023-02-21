package lx

import (
	"fmt"
	"strings"

	"github.com/yadiiiig/tfmod/models"
)

type prs struct {
	Tokens    []Token
	LenTokens int
	Content   []byte
}

func Parser(tokens []Token, content []byte) error {
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

// expects 2 labels (type, name): `resources "foo" "bar" {}`
// len(tokens)-i-4 <= 0 is done because we need atleast 4 tokens: type, name, open-, close bracket
func (p *prs) parse_resource(i int) error {
	r_type, r_name, err := p.parse_resource_headers(i)
	if err != nil {
		return err
	}

	resource := models.Resource{
		Type: r_type,
		Name: r_name,
	}

	start, end, ok := p.find_bracket_combo(i)
	if !ok {
		return fmt.Errorf("resource does not have closing")
	}

	fmt.Printf("%+v\n", resource)
	fmt.Printf("Start -: %d | End index: %d\n", start, end)

	err = p.parse_vars(start, end)
	if err != nil {
		return err
	}

	return nil
}

func (p *prs) parse_resource_headers(start int) (string, string, error) {
	headers := []string{}

	for i := start + 1; i <= p.LenTokens; i++ {
		if p.Tokens[i].Type == TOKEN_QUOTE {
			end, end_i, err := p.find_string_end(i)
			if err != nil {
				return "", "", err
			}

			headers = append(headers, string(p.Content[p.Tokens[i].Ind.Start+1:end]))

			if len(headers) == 2 {
				return headers[0], headers[1], nil
			}

			i = end_i
		}
	}

	return "", "", fmt.Errorf("not able to find a resources type & name")
}

func (p *prs) parse_vars(start, end int) error {
	for i := start + 1; i <= end; i++ {
		if p.Tokens[i].Type == TOKEN_ASSIGN {
			field := models.Field{}
			fmt.Println("found var decl")

			// could hardcode the i-1 (since currenlty the prev string would be a nw token)
			// or loop backwards and stop at the first occurance
			nw, err := p.prev_newline(start, i)
			if err != nil {
				return err
			}

			field.Key = strings.ReplaceAll(string(p.Content[p.Tokens[nw].Ind.End:p.Tokens[i].Ind.Start]), " ", "")

			fmt.Printf("%+v\n", field)
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
		if p.Tokens[i].Type == TOKEN_LEFT_BRACKET {
			if found == 0 {
				start_index = i
			}

			found++
		} else if p.Tokens[i].Type == TOKEN_RIGHT_BRACKET {
			found--

			if found == 0 {
				return start_index + start, i + start, true
			}
		}

	}

	return 0, 0, false
}

func (p *prs) find_string_end(start int) (int, int, error) {
	for i := start + 1; i <= p.LenTokens; i++ {
		if p.Tokens[i].Type == TOKEN_QUOTE {
			return p.Tokens[i].Ind.Start, i, nil
		}
	}

	return 0, 0, fmt.Errorf("could not find string ending")
}
