package lib

import "testing"

func TestMin(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{
			name: "invalid argument",
			args: args{
				arr: nil,
			},
			want:  0,
			want1: 0,
		},
		{
			name: "1st",
			args: args{
				arr: []int{4, 3, 1, 4},
			},
			want:  1,
			want1: 2,
		},
		{
			name: "2nd",
			args: args{
				arr: []int{3, 2, 2, 3},
			},
			want:  2,
			want1: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Min(tt.args.arr)
			if got != tt.want {
				t.Errorf("Min() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Min() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
