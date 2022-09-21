package main

import "testing"

func Test_unpackingString(t *testing.T) {
	type args struct {
		src string
	}
	tests := []struct {
		name          string
		args          args
		wantUnpacking string
		wantErr       bool
	}{
		// TODO: Add test cases.
		{"packed string", args{"iu2cd3c5r"}, "iuucdddcccccr", false},
		{"error string", args{"45"}, "", true},
		{"string with escape", args{"jnraw\\5mksd\\6\\2"}, "jnraw5mksd62 (*)", false},
		{"string with packed escape", args{"gh\\66uy2\\\\4"}, "gh666666uyy\\\\\\\\ (*)", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUnpacking, err := unpackingString(tt.args.src)
			if (err != nil) != tt.wantErr {
				t.Errorf("unpackingString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotUnpacking != tt.wantUnpacking {
				t.Errorf("unpackingString() gotUnpacking = %v, want %v", gotUnpacking, tt.wantUnpacking)
			}
		})
	}
}
