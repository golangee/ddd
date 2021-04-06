package parser

import "testing"

func TestParseText(t *testing.T) {
	tests := []struct {
		name    string
		adl     string
		wantErr bool
	}{
		{
			name:    "hello",
			adl:     `package mypkg



type XY interface{}

type Hello struct{
	World int
}

`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ParseText(tt.adl); (err != nil) != tt.wantErr {
				t.Errorf("ParseText() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
