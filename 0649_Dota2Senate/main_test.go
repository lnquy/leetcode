package main

import "testing"

func Test_predictPartyVictory(t *testing.T) {
	type args struct {
		senate string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"tc1", args{"RD"}, "Radiant"},
		{"tc2", args{"RDD"}, "Dire"},
		{"tc3", args{"DDRRR"}, "Dire"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := predictPartyVictory(tt.args.senate); got != tt.want {
				t.Errorf("predictPartyVictory() = %v, want %v", got, tt.want)
			}
		})
	}
}
