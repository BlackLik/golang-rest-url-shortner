package utils_test

import (
	"testing"

	"urlshort.ru/m/utils"
)

func TestGenerateShortHash(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "Empty string",
			input: "",
			want:  "d41d8cd98f00b204e9800998ecf8427e",
		},
		{
			name:  "Short string",
			input: "hello",
			want:  "5d41402abc4b2a76b9719d911017c592",
		},
		{
			name:  "Long string",
			input: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
			want:  "818c6e601a24f72750da0f6c9b8ebe28",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := utils.GenerateShortHash(tt.input)
			if got != tt.want {
				t.Errorf("GenerateShortHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
