package evaluator

import (
	"monkey/object"
	"reflect"
	"testing"
)

func Test_evalBangOperatorExpression(t *testing.T) {
	type args struct {
		right object.Object
	}
	tests := []struct {
		name string
		args args
		want object.Object
	}{
		{
			name: "!true",
			args: args{
				right: TRUE,
			},
			want: FALSE,
		},
		{
			name: "!false",
			args: args{
				right: FALSE,
			},
			want: TRUE,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evalBangOperatorExpression(tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("evalBangOperatorExpression() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_evalMinusPrefixOperatorExpression(t *testing.T) {
	type args struct {
		right object.Object
	}
	tests := []struct {
		name string
		args args
		want object.Object
	}{
		{
			name: "-5",
			args: args{
				right: &object.Integer{
					Value: 5,
				},
			},
			want: &object.Integer{
				Value: -5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evalMinusPrefixOperatorExpression(tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("evalMinusPrefixOperatorExpression() = %v, want %v", got, tt.want)
			}
		})
	}
}
