package itertools

import (
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	type args struct {
		slice []int
		f     func(item int) int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "can map a simple list of numbers",
			args: args{
				slice: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				f: func(item int) int {
					return 2 * item
				},
			},
			want: []int{2, 4, 6, 8, 10, 12, 14, 16, 18},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Map(tt.args.slice, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}
