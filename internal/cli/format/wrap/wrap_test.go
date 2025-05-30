package wrap_test

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/verily-src/fsh-lint/internal/cli/format/wrap"
)

func TestWrapper_String(t *testing.T) {
	testCases := []struct {
		name      string
		width     int
		splitter  wrap.Splitter
		separator string
		input     string
		want      string
	}{
		{
			name:  "text fits within split",
			width: 40,
			input: "Lorem ipsum dolor sit amet.",
			want:  "Lorem ipsum dolor sit amet.",
		},
		{
			name:  "splits a single line",
			width: 40,
			input: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
			want:  "Lorem ipsum dolor sit amet, consectetur\nadipiscing elit.",
		},
		{
			name:  "preserves empty lines",
			width: 40,
			input: "Lorem ipsum dolor sit amet.\n\nConsectetur adipiscing elit.",
			want:  "Lorem ipsum dolor sit amet.\n\nConsectetur adipiscing elit.",
		},
		{
			name:  "splits multiple lines",
			width: 10,
			input: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
			want:  "Lorem\nipsum\ndolor sit\namet,\nconsectetur\nadipiscing\nelit.",
		},
		{
			name:  "splits overlong token at start of string",
			width: 10,
			input: "Loremipsumdolorsitamet ipsum dolor sit amet.",
			want:  "Loremipsumdolorsitamet\nipsum\ndolor sit\namet.",
		}, {
			name:  "splits overlong token in middle of string",
			width: 10,
			input: "Lorem loremipsumdalorsitamet ipsum dolor sit amet.",
			want:  "Lorem\nloremipsumdalorsitamet\nipsum\ndolor sit\namet.",
		}, {
			name:  "splits overlong token at end of string",
			width: 10,
			input: "Lorem ipsum dalor sit loremipsumdalorsitamet.",
			want:  "Lorem\nipsum\ndalor sit\nloremipsumdalorsitamet.",
		}, {
			name:  "negative width does not wrap",
			width: -1,
			input: "Lorem ipsum dolor sit amet,\nconsectetur adipiscing elit.",
			want:  "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
		}, {
			name:  "custom splitter",
			width: 10,
			splitter: wrap.SplitFunc(func(s string) []string {
				return strings.Split(s, "-")
			}),
			input: "Lorem-ipsum-dolor-sit-amet.",
			want:  "Lorem\nipsum\ndolor sit\namet.",
		}, {
			name:      "custom separator",
			width:     10,
			separator: "-",
			input:     "Lorem ipsum dolor sit amet.",
			want:      "Lorem\nipsum\ndolor-sit\namet.",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			wrapper := wrap.NewWrapper(tc.width).WithSeparator(tc.separator).WithSplitter(tc.splitter)

			got := wrapper.String(tc.input)

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("Wrapper.Strings(%q) mismatch (-want +got):\n%s", tc.input, diff)
			}
		})
	}
}
