package summary_ranges_228

import (
	"reflect"
	"testing"
)

func Test_summaryRanges(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"example 1",
			args{nums: []int{0, 1, 2, 4, 5, 7}},
			[]string{"0->2", "4->5", "7"},
		},
		{
			"example 2",
			args{nums: []int{0, 2, 3, 4, 6, 8, 9}},
			[]string{"0", "2->4", "6", "8->9"},
		},
		{
			"empty input",
			args{nums: []int{}},
			[]string{},
		},
		{
			"single numbers",
			args{nums: []int{0, 2, 4, 8, 10}},
			[]string{"0", "2", "4", "8", "10"},
		},
		{
			"single number",
			args{nums: []int{0}},
			[]string{"0"},
		},
		{
			"single range",
			args{nums: []int{0, 1, 2}},
			[]string{"0->2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := summaryRanges(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("summaryRanges() = %v, want %v", got, tt.want)
			}
		})
	}
}
