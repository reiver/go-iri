package iri

import (
	"net/url"
)

// IRI represents a parsed IRI (internationalized resource identifier).
// Which means that IRI also supports URI and URL, since all valid URI and URL are also value IRI.
//
// You can think of IRI as URI and URL that supports Unicode.
// So, an IRI can include Arabic, Cyrillic, Hangul (Korean), Kanji, Persian, emojis, etc.
// Where, URI and URL would habe to percent-encode to punycode-encode these.
//
// Here is an example IRI:
//
//	http://안녕.kr/세상
//
// The equivalent URI / URL is:
//
//	http://xn--o70b819a.kr/%EC%84%B8%EC%83%81
//
// IRI is intentionally similar to [url.URL].
type IRI struct {
	Scheme   string
	User     *url.Userinfo
	Host     string
	Path     string
	RawQuery string
	Fragment string
}
