package uniq

import (
	"github.com/stretchr/testify/require"
	"testing"
)

var formattingTest = []struct {
	i       bool
	fFields int
	sChars  int
	in      []string
	out     []string
}{
	{ i: true,
		fFields: 0,
		sChars: 0,
		in: []string{"HelloEverybody", "Hell"},
		out: []string{"helloeverybody", "hell"},
	},
	{ i: true,
		fFields: 0,
		sChars: 0,
		in: []string{"HelloEverybody", "Hell", "CMON", "HoWArEY o U"},
		out: []string{"helloeverybody", "hell", "cmon", "howarey o u"},
	},
	{ i: true,
		fFields: 0,
		sChars: 3,
		in: []string{"HelloEverybody", "Hell", "CMON", "HoWArEY o U"},
		out: []string{"loeverybody", "l", "n", "arey o u"},
	},
	{ i: true,
		fFields: 0,
		sChars: 5,
		in: []string{"HelloEverybody", "Hell", "CMON", "HoWArEY o U"},
		out: []string{"everybody", "", "", "ey o u"},
	},
	{ i: false,
		fFields: 1,
		sChars: 4,
		in: []string{
			"We 1ove music.",
			"I lo1e music.",
			"They l1ve music.",
		},
		out: []string{"e music.", "e music.", "e music."},
	},
}

func TestFormattingLinesWithOptions(t *testing.T) {
	for _, tt := range formattingTest {
		t.Run(tt.in[0], func(t *testing.T) {
			result := FormatLinesWithOptions(tt.in, tt.i, tt.sChars, tt.fFields)
			if len(result) != len(tt.out) {
				t.Fatal("Arrays not equal length")
			}

			for i := 0; i < len(result); i++ {
				require.Equal(t, tt.out[i], result[i], "Cell in array %d", i)
			}
		})
	}
}

var stdTest = []struct {
	i       bool
	fFields int
	sChars  int
	in      []string
	out     []string
}{
	{
		i: false,
		fFields: 0,
		sChars: 0,
		in: []string {
			"I love music.",
			"I love music.",
			"I love music.",
			"",
			"I love music of Kartik.",
			"I love music of Kartik.",
			"Thanks.",
			"I love music of Kartik.",
			"I love music of Kartik.",
		},
		out: []string {
			"I love music.",
			"",
			"I love music of Kartik.",
			"Thanks.",
			"I love music of Kartik.",
		},
	},
	{
		i: true,
		fFields: 0,
		sChars: 0,
		in: []string {
			"I LOVE MUSIC.",
			"I love music.",
			"I LoVe MuSiC.",
			"",
			"I love MuSIC of Kartik.",
			"I love music of Kartik.",
			"Thanks.",
			"I love music of Kartik.",
			"I love MuSIC of Kartik.",
		},
		out: []string {
			"I LOVE MUSIC.",
			"",
			"I love MuSIC of Kartik.",
			"Thanks.",
			"I love music of Kartik.",
		},
	},
	{
		i: false,
		fFields: 1,
		sChars: 0,
		in: []string {
			"We love music.",
			"I love music.",
			"They love music.",
			"",
			"I love music of Kartik.",
			"We love music of Kartik.",
			"Thanks.",
		},
		out: []string {
			"We love music.",
			"",
			"I love music of Kartik.",
			"Thanks.",
		},
	},
	{
		i: false,
		fFields: 0,
		sChars: 1,
		in: []string {
			"I love music.",
			"A love music.",
			"C love music.",
			"",
			"I love music of Kartik.",
			"We love music of Kartik.",
			"Thanks.",
		},
		out: []string {
			"I love music.",
			"",
			"I love music of Kartik.",
			"We love music of Kartik.",
			"Thanks.",
		},
	},

}

func TestUniqStandard(t *testing.T) {
	for _, tt := range stdTest {
		t.Run(tt.in[0], func(t *testing.T) {
			opt := Options{WorkMode: 0,
				I: tt.i,
				FFields: tt.fFields,
				SChars: tt.sChars,
			}
			result := Uniq(tt.in, opt)
			if len(result) != len(tt.out) {
				t.Fatalf("Arrays not equal length result: %d out: %d", len(result), len(tt.out))
			}

			for i := 0; i < len(result); i++ {
				require.Equal(t, tt.out[i], result[i], "Cell in array %d", i)
			}
		})
	}
}