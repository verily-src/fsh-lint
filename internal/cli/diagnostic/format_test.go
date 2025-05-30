package diagnostic_test

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/verily-src/fsh-lint/internal/cli/diagnostic"
)

func TestParseFormat(t *testing.T) {
	testCases := []struct {
		name    string
		input   string
		want    diagnostic.Format
		wantErr error
	}{
		{
			name:  "text input",
			input: "text",
			want:  diagnostic.FormatText,
		}, {
			name:  "github input",
			input: "github",
			want:  diagnostic.FormatGitHub,
		}, {
			name:  "json input",
			input: "json",
			want:  diagnostic.FormatJSON,
		}, {
			name:    "invalid input",
			input:   "invalid",
			wantErr: cmpopts.AnyError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := diagnostic.ParseFormat(tc.input)

			if got, want := err, tc.wantErr; !cmp.Equal(got, want, cmpopts.EquateErrors()) {
				t.Errorf("ParseFormat(%q) got error %v, want %v", tc.input, got, want)
			}

			if got, want := got, tc.want; !cmp.Equal(got, want) {
				t.Errorf("ParseFormat(%q) got %v, want %v", tc.input, got, want)
			}
		})
	}
}

func TestFormat_NewPrinter(t *testing.T) {
	testCases := []struct {
		name     string
		format   diagnostic.Format
		wantType reflect.Type
	}{
		{
			name:     "text format",
			format:   diagnostic.FormatText,
			wantType: reflect.TypeFor[*diagnostic.ANSIPrinter](),
		}, {
			name:     "github format",
			format:   diagnostic.FormatGitHub,
			wantType: reflect.TypeFor[*diagnostic.GitHubPrinter](),
		}, {
			name:     "json format",
			format:   diagnostic.FormatJSON,
			wantType: reflect.TypeFor[*diagnostic.JSONPrinter](),
		}, {
			name:     "invalid format returns default printer",
			format:   "invalid",
			wantType: reflect.TypeOf(diagnostic.DefaultPrinter),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.format.Printer(nil)

			if got, want := reflect.TypeOf(got), tc.wantType; got != want {
				t.Errorf("Format.Printer() got %v, want %v", got, want)
			}
		})
	}
}
