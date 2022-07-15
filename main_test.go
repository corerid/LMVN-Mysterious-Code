package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func Test_reverseMsg(t *testing.T) {
	tests := []struct {
		name     string
		inputMsg []byte
		expected string
	}{
		{
			name:     "test_reverse_message_success",
			inputMsg: []byte("this is test"),
			expected: "tset si siht",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseMsg(tt.inputMsg); got != tt.expected {
				t.Errorf("reverseMsg() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func Test_Main(t *testing.T) {
	tests := []struct {
		name     string
		expected string
	}{
		{
			name:     "test_main_success",
			expected: "Join us at LINE MAN Wongnai. Okay!, I accept the invitation.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			stdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			main()

			w.Close()
			out, _ := ioutil.ReadAll(r)
			os.Stdout = stdout

			if string(out) != tt.expected {
				t.Errorf("main() = %v, want %v", string(out), tt.expected)
			}
		})
	}
}
