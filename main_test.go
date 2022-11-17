package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_bubbleSort(t *testing.T) {
	type args struct {
		arg []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "arg is null",
			args: args{
				arg: nil,
			},
			want: nil,
		},
		{
			name: "length of arg is 1",
			args: args{
				arg: []int{1},
			},
			want: []int{1},
		},
		{
			name: "1st",
			args: args{
				arg: []int{3, 1, 2},
			},
			want: []int{1, 2, 3},
		},
		{
			name: "2nd",
			args: args{
				arg: []int{4, 1, 2, 3, 2, 2, 5, 1},
			},
			want: []int{1, 1, 2, 2, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := bubbleSort(tt.args.arg)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("unexpected diff; %s", diff)
			}
		})
	}
}

func Test_selectSort(t *testing.T) {
	type args struct {
		arg []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "arg is null",
			args: args{
				arg: nil,
			},
			want: nil,
		},
		{
			name: "1st",
			args: args{
				arg: []int{3, 1, 4, 2},
			},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "2nd",
			args: args{
				arg: []int{2, 2, 1, 4, 5, 4},
			},
			want: []int{1, 2, 2, 4, 4, 5},
		},
		{
			name: "3rd",
			args: args{
				arg: []int{4, 1, 2, 3, 2, 2, 5, 1},
			},
			want: []int{1, 1, 2, 2, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := selectSort(tt.args.arg)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("unexpected diff; %s", diff)
			}
		})
	}
}
