package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
	"reflect"
	"testing"
)

func TestParser_parseLetStatement(t *testing.T) {

	tests := []struct {
		name   string
		fields *Parser
		want   *ast.LetStatement
	}{
		{
			name:   "",
			fields: New(lexer.New("let five = 5;")),
			want: &ast.LetStatement{
				Token: token.Token{
					Type:    token.LET,
					Literal: "let",
				},
				Name: &ast.Identifier{
					Token: token.Token{
						Type:    token.IDENT,
						Literal: "five",
					},
					Value: "five",
				},
				Value: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.fields
			if got := p.parseLetStatement(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseLetStatement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParser_parseReturnStatement(t *testing.T) {

	tests := []struct {
		name   string
		fields *Parser
		want   *ast.ReturnStatement
	}{
		{
			name:   "Parse return statement",
			fields: New(lexer.New("return 5;")),
			want: &ast.ReturnStatement{
				Token: token.Token{
					Type:    token.RETURN,
					Literal: "return",
				},
				ReturnValue: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.fields
			if got := p.parseReturnStatement(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseReturnStatement() = %v, want %v", got, tt.want)
			}
		})
	}
}
