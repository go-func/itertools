package itertools

import (
	"reflect"
	"testing"
)

func TestReduce(t *testing.T) {
	type args struct {
		slice  []int
		iValue int
		f      func(pValue int, cValue int, index int, slice []int) int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "can reduce a bunch of numbers with no initial value",
			args: args{
				slice:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				iValue: 0,
				f: func(pValue int, cValue int, index int, slice []int) int {
					return pValue + slice[index]
				},
			},
			want: 45,
		},
		{
			name: "can reduce a bunch of numbers with initial value",
			args: args{
				slice:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				iValue: 10,
				f: func(pValue int, cValue int, index int, slice []int) int {
					return pValue + slice[index]
				},
			},
			want: 55,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reduce(tt.args.slice, tt.args.iValue, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reduce() = %v, want %v", got, tt.want)
			}
		})
	}
}
