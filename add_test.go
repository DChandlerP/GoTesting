package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	type args struct {
		x int
		y int
	}
	var tests = []struct {
		name string
		args args
		want int
	}{
		{"5 + 5", args{x: 5, y: 5}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
