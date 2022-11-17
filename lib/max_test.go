package lib

import "testing"

func TestMax(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "invalid argument",
			args: args{
				arr: nil,
			},
			want: 0,
		},
		{
			name: "1st",
			args: args{
				arr: []int{3, 1, 2},
			},
			want: 3,
		},
		{
			name: "2nd",
			args: args{
				arr: []int{3, 1, 2, 4},
			},
			want: 4,
		},
		{
			name: "3rd",
			args: args{
				arr: []int{1, 3, 5, 4, 2, 3, 1, 1},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.args.arr); got != tt.want {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}
