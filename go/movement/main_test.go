package main

import (
	"bytes"
	"reflect"
	"testing"
)

func Test_genNumbers(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"n = 0",
			args{0},
			nil,
		},
		{
			"n = 1",
			args{1},
			[][]int{
				{1},
			},
		},
		{
			"n = 2",
			args{2},
			[][]int{
				{1, 2},
				{4, 3},
			},
		},
		{
			"n = 3",
			args{3},
			[][]int{
				{1, 2, 3},
				{8, 9, 4},
				{7, 6, 5},
			},
		},
		{
			"n = 4",
			args{4},
			[][]int{
				{1, 2, 3, 4},
				{12, 13, 14, 5},
				{11, 16, 15, 6},
				{10, 9, 8, 7},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := genNumbers(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("genNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_printNumbers(t *testing.T) {
	type args struct {
		numbers [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantOut string
	}{
		{
			"n = 1",
			args{[][]int{{1}}},
			"1 \n",
		},
		{
			"n = 2",
			args{[][]int{
				{1, 2},
				{4, 3},
			}},
			"1 2 \n" +
				"4 3 \n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := make([]byte, 0)
			out := bytes.NewBuffer(buf)
			printNumbers(tt.args.numbers, out)

			if gotOut := out.String(); gotOut != tt.wantOut {
				t.Errorf("printNumbers() = `%v`, want `%v`", gotOut, tt.wantOut)
			}
		})
	}
}
