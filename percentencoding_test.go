package iri_test

import (
	"testing"

	"github.com/reiver/go-iri"
)

func TestPercentEncodeString(t *testing.T) {
	tests := []struct {
		Name   string
		Value  string
		Expected string
	}{
		// Empty string.
		{
			Name:     "empty",
			Value:    "",
			Expected: "",
		},

		// ASCII unreserved characters — should NOT be encoded.
		{
			Name:     "lowercase-letters",
			Value:    "abcdefghijklmnopqrstuvwxyz",
			Expected: "abcdefghijklmnopqrstuvwxyz",
		},
		{
			Name:     "uppercase-letters",
			Value:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			Expected: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		},
		{
			Name:     "digits",
			Value:    "0123456789",
			Expected: "0123456789",
		},
		{
			Name:     "hyphen",
			Value:    "-",
			Expected: "-",
		},
		{
			Name:     "period",
			Value:    ".",
			Expected: ".",
		},
		{
			Name:     "underscore",
			Value:    "_",
			Expected: "_",
		},
		{
			Name:     "tilde",
			Value:    "~",
			Expected: "~",
		},

		// ASCII reserved characters — should be encoded.
		{
			Name:     "at-sign",
			Value:    "@",
			Expected: "%40",
		},
		{
			Name:     "colon",
			Value:    ":",
			Expected: "%3A",
		},
		{
			Name:     "slash",
			Value:    "/",
			Expected: "%2F",
		},
		{
			Name:     "question-mark",
			Value:    "?",
			Expected: "%3F",
		},
		{
			Name:     "hash",
			Value:    "#",
			Expected: "%23",
		},
		{
			Name:     "open-bracket",
			Value:    "[",
			Expected: "%5B",
		},
		{
			Name:     "close-bracket",
			Value:    "]",
			Expected: "%5D",
		},
		{
			Name:     "exclamation",
			Value:    "!",
			Expected: "%21",
		},
		{
			Name:     "dollar",
			Value:    "$",
			Expected: "%24",
		},
		{
			Name:     "ampersand",
			Value:    "&",
			Expected: "%26",
		},
		{
			Name:     "single-quote",
			Value:    "'",
			Expected: "%27",
		},
		{
			Name:     "open-paren",
			Value:    "(",
			Expected: "%28",
		},
		{
			Name:     "close-paren",
			Value:    ")",
			Expected: "%29",
		},
		{
			Name:     "asterisk",
			Value:    "*",
			Expected: "%2A",
		},
		{
			Name:     "plus",
			Value:    "+",
			Expected: "%2B",
		},
		{
			Name:     "comma",
			Value:    ",",
			Expected: "%2C",
		},
		{
			Name:     "semicolon",
			Value:    ";",
			Expected: "%3B",
		},
		{
			Name:     "equals",
			Value:    "=",
			Expected: "%3D",
		},

		// Other ASCII characters — should be encoded.
		{
			Name:     "space",
			Value:    " ",
			Expected: "%20",
		},
		{
			Name:     "percent",
			Value:    "%",
			Expected: "%25",
		},
		{
			Name:     "backslash",
			Value:    "\\",
			Expected: "%5C",
		},
		{
			Name:     "caret",
			Value:    "^",
			Expected: "%5E",
		},
		{
			Name:     "backtick",
			Value:    "`",
			Expected: "%60",
		},
		{
			Name:     "open-brace",
			Value:    "{",
			Expected: "%7B",
		},
		{
			Name:     "pipe",
			Value:    "|",
			Expected: "%7C",
		},
		{
			Name:     "close-brace",
			Value:    "}",
			Expected: "%7D",
		},
		{
			Name:     "tab",
			Value:    "\t",
			Expected: "%09",
		},
		{
			Name:     "newline",
			Value:    "\n",
			Expected: "%0A",
		},
		{
			Name:     "null",
			Value:    "\x00",
			Expected: "%00",
		},

		// ucschar — should NOT be encoded (RFC 3987 Section 2.2).
		{
			Name:     "hangul-korean",
			Value:    "세상",
			Expected: "세상",
		},
		{
			Name:     "arabic",
			Value:    "مرحبا",
			Expected: "مرحبا",
		},
		{
			Name:     "cyrillic",
			Value:    "Привет",
			Expected: "Привет",
		},
		{
			Name:     "kanji",
			Value:    "漢字",
			Expected: "漢字",
		},
		{
			Name:     "persian",
			Value:    "سلام",
			Expected: "سلام",
		},
		{
			Name:     "persian-alphabet-alef",
			Value:    "ا",
			Expected: "ا",
		},
		{
			Name:     "persian-alphabet-be",
			Value:    "ب",
			Expected: "ب",
		},
		{
			Name:     "persian-alphabet-pe",
			Value:    "پ",
			Expected: "پ",
		},
		{
			Name:     "persian-alphabet-te",
			Value:    "ت",
			Expected: "ت",
		},
		{
			Name:     "persian-alphabet-se",
			Value:    "ث",
			Expected: "ث",
		},
		{
			Name:     "persian-alphabet-jim",
			Value:    "ج",
			Expected: "ج",
		},
		{
			Name:     "persian-alphabet-che",
			Value:    "چ",
			Expected: "چ",
		},
		{
			Name:     "persian-alphabet-he",
			Value:    "ح",
			Expected: "ح",
		},
		{
			Name:     "persian-alphabet-khe",
			Value:    "خ",
			Expected: "خ",
		},
		{
			Name:     "persian-alphabet-dal",
			Value:    "د",
			Expected: "د",
		},
		{
			Name:     "persian-alphabet-zal",
			Value:    "ذ",
			Expected: "ذ",
		},
		{
			Name:     "persian-alphabet-re",
			Value:    "ر",
			Expected: "ر",
		},
		{
			Name:     "persian-alphabet-ze",
			Value:    "ز",
			Expected: "ز",
		},
		{
			Name:     "persian-alphabet-zhe",
			Value:    "ژ",
			Expected: "ژ",
		},
		{
			Name:     "persian-alphabet-sin",
			Value:    "س",
			Expected: "س",
		},
		{
			Name:     "persian-alphabet-shin",
			Value:    "ش",
			Expected: "ش",
		},
		{
			Name:     "persian-alphabet-sad",
			Value:    "ص",
			Expected: "ص",
		},
		{
			Name:     "persian-alphabet-zad",
			Value:    "ض",
			Expected: "ض",
		},
		{
			Name:     "persian-alphabet-ta",
			Value:    "ط",
			Expected: "ط",
		},
		{
			Name:     "persian-alphabet-za",
			Value:    "ظ",
			Expected: "ظ",
		},
		{
			Name:     "persian-alphabet-eyn",
			Value:    "ع",
			Expected: "ع",
		},
		{
			Name:     "persian-alphabet-gheyn",
			Value:    "غ",
			Expected: "غ",
		},
		{
			Name:     "persian-alphabet-fe",
			Value:    "ف",
			Expected: "ف",
		},
		{
			Name:     "persian-alphabet-ghaf",
			Value:    "ق",
			Expected: "ق",
		},
		{
			Name:     "persian-alphabet-kaf",
			Value:    "ک",
			Expected: "ک",
		},
		{
			Name:     "persian-alphabet-gaf",
			Value:    "گ",
			Expected: "گ",
		},
		{
			Name:     "persian-alphabet-lam",
			Value:    "ل",
			Expected: "ل",
		},
		{
			Name:     "persian-alphabet-mim",
			Value:    "م",
			Expected: "م",
		},
		{
			Name:     "persian-alphabet-nun",
			Value:    "ن",
			Expected: "ن",
		},
		{
			Name:     "persian-alphabet-vav",
			Value:    "و",
			Expected: "و",
		},
		{
			Name:     "persian-alphabet-he-do-cheshm",
			Value:    "ه",
			Expected: "ه",
		},
		{
			Name:     "persian-alphabet-ye",
			Value:    "ی",
			Expected: "ی",
		},
		{
			Name:     "persian-alphabet-full-word",
			Value:    "ایران",
			Expected: "ایران",
		},
		{
			Name:     "persian-number-0",
			Value:    "۰",
			Expected: "۰",
		},
		{
			Name:     "persian-number-1",
			Value:    "۱",
			Expected: "۱",
		},
		{
			Name:     "persian-number-2",
			Value:    "۲",
			Expected: "۲",
		},
		{
			Name:     "persian-number-3",
			Value:    "۳",
			Expected: "۳",
		},
		{
			Name:     "persian-number-4",
			Value:    "۴",
			Expected: "۴",
		},
		{
			Name:     "persian-number-5",
			Value:    "۵",
			Expected: "۵",
		},
		{
			Name:     "persian-number-6",
			Value:    "۶",
			Expected: "۶",
		},
		{
			Name:     "persian-number-7",
			Value:    "۷",
			Expected: "۷",
		},
		{
			Name:     "persian-number-8",
			Value:    "۸",
			Expected: "۸",
		},
		{
			Name:     "persian-number-9",
			Value:    "۹",
			Expected: "۹",
		},
		{
			Name:     "persian-numbers-all",
			Value:    "۰۱۲۳۴۵۶۷۸۹",
			Expected: "۰۱۲۳۴۵۶۷۸۹",
		},
		{
			Name:     "persian-mixed-text-and-numbers",
			Value:    "سال۱۴۰۴",
			Expected: "سال۱۴۰۴",
		},
		{
			Name:     "emoji-smiling-horns-U+1F608",
			Value:    "😈",
			Expected: "😈",
		},
		{
			Name:     "emoji-grinning-face-U+1F600",
			Value:    "😀",
			Expected: "😀",
		},
		{
			Name:     "ucschar-boundary-low-U+00A0",
			Value:    "\u00A0",
			Expected: "\u00A0",
		},
		{
			Name:     "ucschar-boundary-high-U+D7FF",
			Value:    "\uD7FF",
			Expected: "\uD7FF",
		},
		{
			Name:     "ucschar-range-U+F900",
			Value:    "\uF900",
			Expected: "\uF900",
		},
		{
			Name:     "ucschar-range-U+FDF0",
			Value:    "\uFDF0",
			Expected: "\uFDF0",
		},

		// Characters outside ucschar ranges — should be encoded.
		{
			Name:     "outside-ucschar-U+FFFE",
			Value:    "\uFFFE",
			Expected: "%EF%BF%BE",
		},
		{
			Name:     "outside-ucschar-U+FFFF",
			Value:    "\uFFFF",
			Expected: "%EF%BF%BF",
		},

		// Mixed content.
		{
			Name:     "mixed-unreserved-and-reserved",
			Value:    "hello@world",
			Expected: "hello%40world",
		},
		{
			Name:     "mixed-ascii-and-unicode",
			Value:    "user-세상",
			Expected: "user-세상",
		},
		{
			Name:     "mixed-unicode-and-reserved",
			Value:    "세상@漢字",
			Expected: "세상%40漢字",
		},
		{
			Name:     "iri-user-component",
			Value:    "joe@example",
			Expected: "joe%40example",
		},
		{
			Name:     "multiple-reserved-characters",
			Value:    "a:b@c/d?e#f",
			Expected: "a%3Ab%40c%2Fd%3Fe%23f",
		},
	}

	for testNumber, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			actual := iri.PercentEncodeString(test.Value)

			expected := test.Expected

			if test.Expected != actual {
				t.Errorf("For test #%d, the actual IRI-style percent-encoded it not what was expected.", testNumber)
				t.Logf("EXPECTED: %q", expected)
				t.Logf("ACTUAL:   %q", actual)
				t.Logf("VALUE:    %q", test.Value)
			}
		})
	}
}
