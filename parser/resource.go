package parser

import (
	"fmt"

	"github.com/yadiiiig/tfmod/models"
)

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

// returns header type, name or error
func (p *prs) parse_resource_headers(start int) (string, string, error) {
	headers := []string{}

	for i := start + 1; i <= p.LenTokens; i++ {
		if p.Tokens[i].Type == TOKEN_QUOTE {
			end, end__tkn_i, err := p.find_string_end(i)
			if err != nil {
				return "", "", err
			}

			headers = append(headers, p.get_string(p.Tokens[i].Ind.Start+1, end, true))

			if len(headers) == 2 {
				return headers[0], headers[1], nil
			}

			i = end__tkn_i
		}
	}

	return "", "", fmt.Errorf("not able to find a resources type & name")
}

func (p *prs) parse_vars(start, end int) error {
	for i := start + 1; i <= end; i++ {
		if p.Tokens[i].Type == TOKEN_ASSIGN {
			field := models.Argument{}

			nw, err := p.prev_newline(start, i)
			if err != nil {
				return err
			}

			field.Key = p.get_string(p.Tokens[nw].Ind.End, p.Tokens[i].Ind.Start, true)

			struct_type, err := p.parse_var_type(i, end, &field)
			if err != nil {
				return err
			}

			if struct_type {
				fmt.Println("found structure decleration")
				// parse_vars
			}

			fmt.Printf("%+v\n", field)
		}
	}

	return nil
}

// returns true if value is a structure def or an error
func (p *prs) parse_var_type(start, end int, field *models.Argument) (bool, error) {
	for i := start + 1; i <= end; i++ {
		switch p.Tokens[i].Type {
		case TOKEN_VAR:
			field.Type = TOKEN_VAR
			field.Edit = true
			return false, nil

		case TOKEN_DATA:
			field.Type = TOKEN_DATA
			return false, nil

		case TOKEN_POINT:
			field.Type = TOKEN_CUSTOM_TYPE
			field.CustomType = p.get_string(p.Tokens[start].Ind.End, p.Tokens[i].Ind.Start, true)
			return false, nil

		case TOKEN_LEFT_BRACKET:
			return true, nil

		case TOKEN_QUOTE:
			end_ind, _, err := p.find_string_end(i)
			if err != nil {
				return false, err
			}

			field.Type = TOKEN_HARDCODED_STRING
			field.Hardcoded = p.get_string(p.Tokens[i].Ind.End, end_ind, false)

			return false, nil
		}
	}

	return false, fmt.Errorf("could not find type decleration")
}
