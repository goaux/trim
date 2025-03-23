package trim_test

import (
	"fmt"
	"testing"

	"github.com/goaux/trim"
)

func ExampleIndent() {
	fmt.Println(trim.Indent(`
	  ABC
	    123
	  456
	`))
	// Output:
	// ABC
	//   123
	// 456
}

func ExampleMargin() {
	fmt.Println(trim.Margin(`
		    > ABC
		    >   123
		    > 456
		`, "> ",
	))
	// Output:
	// ABC
	//   123
	// 456
}

func ExampleIsBlank() {
	fmt.Println(trim.IsBlank("  \t\n"))
	fmt.Println(trim.IsBlank("hello"))
	s := " \u0020\u3000　\t\v\n\r\n"
	fmt.Printf("% 02x : %v\n", s, trim.IsBlank(s))
	// Output:
	// true
	// false
	// 20 20 e3 80 80 e3 80 80 09 0b 0a 0d 0a : true
}

func TestIndent(t *testing.T) {
	testcases := []struct {
		Text string
		Want string
	}{
		// Unicode whitespace
		{"\u2002\u2003ABC", "ABC"},

		// Deep nested indentation
		{`
			    First Level
				    Second Level
					    Third Level
		`, "First Level\n Second Level\n  Third Level"},

		// Mixed indentation styles
		{`
		    	TAB and spaces
		    		More mixed
		`, "TAB and spaces\n\tMore mixed"},
	}

	for i, tc := range testcases {
		got := trim.Indent(tc.Text)
		if got != tc.Want {
			t.Errorf("[%d] text=%q got=%q, want=%q", i, tc.Text, got, tc.Want)
		}
	}
}

func TestMargin(t *testing.T) {
	testcases := []struct {
		Text   string
		Prefix string
		Want   string
	}{
		// Unicode prefix
		{"\n❯ Hello\n❯ World\n", "❯ ", "Hello\nWorld"},

		// Complex Unicode scenario
		{`
		🌈 Rainbow
		🌈   Colorful
		🌈     Text
		`, "🌈 ", "Rainbow\n  Colorful\n    Text"},

		{"\n       ABC\nprefix 123\n       456\n              ",
			"prefix ",
			"       ABC\n123\n       456"},
	}

	for i, tc := range testcases {
		got := trim.Margin(tc.Text, tc.Prefix)
		if got != tc.Want {
			t.Errorf("[%d] text=%q got=%q, want=%q", i, tc.Text, got, tc.Want)
		}
	}
}

func TestIsBlank(t *testing.T) {
	t.Run("Unicodes", func(t *testing.T) {
		blankTests := []struct {
			Text string
			Want bool
		}{
			{"\u0020", true}, // SPACE
			{"\u00A0", true}, // NO-BREAK
			{"\u2000", true}, // EN QUAD
			{"\u2001", true}, // EM QUAD
			{"\u2002", true}, // EN SPACE
			{"\u2003", true}, // EM SPACE
			{"\u2004", true}, // THREE-PER-EM SPACE
			{"\u2005", true}, // FOUR-PER-EM SPACE
			{"\u2006", true}, // SIX-PER-EM SPACE
			{"\u2007", true}, // FIGURE SPACE
			{"\u2008", true}, // PUNCTUATION SPACE
			{"\u2009", true}, // THIN SPACE
			{"\u200A", true}, // HAIR SPACE
			{"\u3000", true}, // IDEOGRAPHIC SPACE

			{"\u2002\u2003\u2004\u2005\u200A", true}, // Various Unicode spaces

			{"\u2002a\u2003", false},
		}

		for i, tt := range blankTests {
			got := trim.IsBlank(tt.Text)
			if got != tt.Want {
				t.Errorf("[%d] text=%q got=%v want=%v", i, tt.Text, got, tt.Want)
			}
		}
	})
}
