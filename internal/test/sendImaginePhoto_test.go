package test

import (
	bot "openjourney/internal/telegram"
	"testing"
)

func isOK(text string) string {
	if bot.IsPrompt(text) {
		return "Done"
	} else {
		return "Err"
	}
}

func TestSendImaginePhoto(t *testing.T) {
	type testCase struct {
		name string
		arg  string
		want string
	}

	tests := []testCase{
		{
			name: "Empty command",
			arg:  "/imagine",
			want: "Err",
		},
		{
			name: "Correct request",
			arg:  "/imagine beaver",
			want: "Done",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			response := isOK(test.arg)

			if response != test.want {
				t.Fatal()
			}
		})
	}
}
