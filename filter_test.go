package itertools

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	type args struct {
		slice []int
		f     func(item int) bool
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "can filter out simple numbers",
			args: args{
				slice: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
				f: func(item int) bool {
					return item > 5
				},
			},
			want: []int{6, 7, 8, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.args.slice, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}
