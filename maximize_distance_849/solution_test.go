package maximize_distance_849

import "testing"

func Test_maxDistToClosest(t *testing.T) {
	type args struct {
		seats []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"example 1",
			args{seats: []int{1, 0, 0, 0, 1, 0, 1}},
			2,
		},
		{
			"example 2",
			args{seats: []int{1, 0, 0, 0}},
			3,
		},
		{
			"example 3",
			args{seats: []int{0, 1}},
			1,
		},
		{
			"one free seat",
			args{seats: []int{1, 0, 1}},
			1,
		},
		{
			"best seat from left",
			args{seats: []int{0, 0, 0, 1}},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDistToClosest(tt.args.seats); got != tt.want {
				t.Errorf("maxDistToClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}
