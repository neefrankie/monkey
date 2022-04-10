package evaluator

import (
	"monkey/object"
	"testing"
)

func Test_nativeBoolToBooleanObject(t *testing.T) {
	type args struct {
		input bool
	}
	tests := []struct {
		name string
		args args
		want *object.Boolean
	}{
		{
			name: "true to TRUE",
			args: args{input: true},
			want: TRUE,
		},
		{
			name: "false to FALSE",
			args: args{input: false},
			want: FALSE,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := nativeBoolToBooleanObject(tt.args.input)

			if got != tt.want {
				t.Errorf("nativeBoolToBooleanObject() = %v, want %v", got, tt.want)
			}
		})
	}
}
