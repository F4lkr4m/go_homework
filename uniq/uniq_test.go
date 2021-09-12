package uniq

import (
	"github.com/stretchr/testify/require"
	"go_homework/uniq/utils"
	"strconv"
	"testing"
)

var formattingTest = []struct {
	opt utils.Options
	in  []string
	out []string
}{
	{
		opt: utils.Options{CaseInsensitive: true, FFields: 0, SChars: 0},
		in:  []string{"HelloEverybody", "Hell"},
		out: []string{"helloeverybody", "hell"},
	},
	{
		opt: utils.Options{CaseInsensitive: true, FFields: 0, SChars: 0},
		in:  []string{"HelloEverybody", "Hell", "CMON", "HoWArEY o U"},
		out: []string{"helloeverybody", "hell", "cmon", "howarey o u"},
	},
	{
		opt: utils.Options{CaseInsensitive: true, FFields: 0, SChars: 3},
		in:  []string{"HelloEverybody", "Hell", "CMON", "HoWArEY o U"},
		out: []string{"loeverybody", "l", "n", "arey o u"},
	},
	{
		opt: utils.Options{CaseInsensitive: true, FFields: 0, SChars: 5},
		in:  []string{"HelloEverybody", "Hell", "CMON", "HoWArEY o U"},
		out: []string{"everybody", "", "", "ey o u"},
	},
	{
		opt: utils.Options{CaseInsensitive: false, FFields: 1, SChars: 4},
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
	testName string
	opt      utils.Options
	in       []string
	out      []string
}{
	{
		testName: "Uniq standard test",
		opt:      utils.Options{WorkMode: utils.None, CaseInsensitive: false, FFields: 0, SChars: 0},
		in:       similarTestCase,
		out: []string{
			"I love music.",
			"",
			"I love music of Kartik.",
			"Thanks.",
			"I love music of Kartik.",
		},
	},
	{
		testName: "Uniq standard different case test",
		opt:      utils.Options{WorkMode: utils.None, CaseInsensitive: true, FFields: 0, SChars: 0},
		in: []string{
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
		out: []string{
			"I LOVE MUSIC.",
			"",
			"I love MuSIC of Kartik.",
			"Thanks.",
			"I love music of Kartik.",
		},
	},
	{
		testName: "Uniq standard with -f 1 test",
		opt:      utils.Options{WorkMode: utils.None, CaseInsensitive: false, FFields: 1, SChars: 0},
		in: []string{
			"We love music.",
			"I love music.",
			"They love music.",
			"",
			"I love music of Kartik.",
			"We love music of Kartik.",
			"Thanks.",
		},
		out: []string{
			"We love music.",
			"",
			"I love music of Kartik.",
			"Thanks.",
		},
	},
	{
		testName: "Uniq standard with -s 1 test",
		opt:      utils.Options{WorkMode: utils.None, CaseInsensitive: false, FFields: 0, SChars: 1},
		in: []string{
			"I love music.",
			"A love music.",
			"C love music.",
			"",
			"I love music of Kartik.",
			"We love music of Kartik.",
			"Thanks.",
		},
		out: []string{
			"I love music.",
			"",
			"I love music of Kartik.",
			"We love music of Kartik.",
			"Thanks.",
		},
	},
	{
		testName: "Uniq counting lines mode test",
		opt:      utils.Options{WorkMode: utils.C, CaseInsensitive: false, SChars: 0, FFields: 0},
		in:       similarTestCase,
		out: []string{
			"3 I love music.",
			"1 ",
			"2 I love music of Kartik.",
			"1 Thanks.",
			"2 I love music of Kartik.",
		},
	},
	{
		testName: "Uniq counting lines mode with case-insensitive and -f 1 test",
		opt:      utils.Options{WorkMode: utils.C, CaseInsensitive: true, SChars: 0, FFields: 1},
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
		testName: "Uniq not repeating lines mode test",
		opt:      utils.Options{WorkMode: utils.U, CaseInsensitive: false, FFields: 0, SChars: 0},
		in:       similarTestCase,
		out: []string{
			"",
			"Thanks.",
		},
	},
	{
		testName: "Uniq repeating lines mode test",
		opt:      utils.Options{WorkMode: utils.D, CaseInsensitive: false, FFields: 0, SChars: 0},
		in:       similarTestCase,
		out: []string{
			"I love music.",
			"I love music of Kartik.",
			"I love music of Kartik.",
		},
	},
	{
		testName: "Uniq standard mode with repeating lines input test",
		opt:      utils.Options{WorkMode: utils.None, CaseInsensitive: false, FFields: 0, SChars: 0},
		in: []string{
			"Hello",
			"Hello",
			"Hello",
			"Hell",
			"Hell",
		},
		out: []string{
			"Hello",
			"Hell",
		},
	},
}

func testingUniqFunc(t *testing.T, tt struct {
	testName string
	opt      utils.Options
	in       []string
	out      []string
}) {
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
		t.Run("case num"+strconv.Itoa(index)+" "+tt.testName, func(t *testing.T) {
			testingUniqFunc(t, tt)
		})
	}
}
