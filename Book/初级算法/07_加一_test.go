package book01

import (
	"reflect"
	"testing"
)

func Test_plusOne(t *testing.T) {
	type args struct {
		digits []int
	}
	tests := []struct {
		name       string
		args       args
		wantOutput []int
	}{
		{
			name:       "Success",
			args:       args{digits: []int{4, 3, 2, 1}},
			wantOutput: []int{4, 3, 2, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := plusOne(tt.args.digits); !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("plusOne() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
