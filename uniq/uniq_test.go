package uniq

import (
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

var formattingTest = []struct {
	opt Options
	in      []string
	out     []string
}{
	{
		opt: Options{I: true, FFields: 0, SChars: 0},
		in: []string{"HelloEverybody", "Hell"},
		out: []string{"helloeverybody", "hell"},
	},
	{
		opt: Options{I: true, FFields: 0, SChars: 0},
		in: []string{"HelloEverybody", "Hell", "CMON", "HoWArEY o U"},
		out: []string{"helloeverybody", "hell", "cmon", "howarey o u"},
	},
	{
		opt: Options{I: true, FFields: 0, SChars: 3},
		in: []string{"HelloEverybody", "Hell", "CMON", "HoWArEY o U"},
		out: []string{"loeverybody", "l", "n", "arey o u"},
	},
	{
		opt: Options{I: true, FFields: 0, SChars: 5},
		in: []string{"HelloEverybody", "Hell", "CMON", "HoWArEY o U"},
		out: []string{"everybody", "", "", "ey o u"},
	},
	{
		opt: Options{I: false, FFields: 1, SChars: 4},
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
			var result []string
			for _, item := range tt.in {
				result = append(result, formatLineWithOptions(item, tt.opt))
			}

			if len(result) != len(tt.out) {
				t.Fatal("Arrays not equal length")
			}

			for i := 0; i < len(result); i++ {
				require.Equal(t, tt.out[i], result[i], "Cell in array %d", i)
			}
		})
	}
}

var similarTestCase = []string{
	"I love music.",
	"I love music.",
	"I love music.",
	"",
	"I love music of Kartik.",
	"I love music of Kartik.",
	"Thanks.",
	"I love music of Kartik.",
	"I love music of Kartik.",
}

var uniqTestCases = []struct {
	opt Options
	in      []string
	out     []string
}{
	{
		opt: Options{WorkMode: None, I: false, FFields: 0, SChars: 0},
		in: similarTestCase,
		out: []string {
			"I love music.",
			"",
			"I love music of Kartik.",
			"Thanks.",
			"I love music of Kartik.",
		},
	},
	{
		opt: Options{WorkMode: None, I: true, FFields: 0, SChars: 0},
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
		opt: Options{WorkMode: None, I: false, FFields: 1, SChars: 0},
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
		opt: Options{WorkMode: None, I: false, FFields: 0, SChars: 1},
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
	{
		opt: Options{WorkMode: C, I: false, SChars: 0, FFields: 0},
		in: similarTestCase,
		out: []string{
			"3 I love music.",
			"1 ",
			"2 I love music of Kartik.",
			"1 Thanks.",
			"2 I love music of Kartik.",
		},
	},
	{
		opt: Options{WorkMode: C, I: true, SChars: 0, FFields: 1},
		in: []string{
			"i love music",
			"a lovE music",
			"fedor Love music",
			"",
			"Hello",
			"hello",
			"",
		},
		out: []string{
			"3 i love music",
			"4 ",
		},
	},
	{
		opt: Options{WorkMode: U, I: false, FFields: 0, SChars: 0},
		in: similarTestCase,
		out: []string {
			"",
			"Thanks.",
		},
	},
	{
		opt: Options{WorkMode: D, I: false, FFields: 0, SChars: 0},
		in: similarTestCase,
		out: []string {
			"I love music.",
			"I love music of Kartik.",
			"I love music of Kartik.",
		},
	},
	{
		opt: Options{WorkMode: U, I: false, FFields: 0, SChars: 0},
		in: similarTestCase,
		out: []string {
			"",
			"Thanks.",
		},
	},
	{
		opt: Options{WorkMode: None, I: false, FFields: 0, SChars: 0},
		in: []string {
			"Hello",
			"Hello",
			"Hello",
			"Hell",
			"Hell",
		},
		out: []string {
			"Hello",
			"Hell",
		},
	},
}

func testingUniqFunc(t *testing.T, tt struct{
											opt Options
											in[]string
											out []string}) {
	result := Uniq(tt.in, tt.opt)
	if len(result) != len(tt.out) {
		t.Fatalf("Arrays not equal length result: %d out: %d", len(result), len(tt.out))
	}

	for i := 0; i < len(result); i++ {
		require.Equal(t, tt.out[i], result[i], "Cell in array %d", i)
	}
}


func TestUniq(t *testing.T) {
	for index, tt := range uniqTestCases {
		t.Run("case num" + strconv.Itoa(index) , func(t *testing.T) {
			testingUniqFunc(t, tt)
		})
	}
}