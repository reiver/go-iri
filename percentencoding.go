package iri

import (
	"unicode/utf8"

	"github.com/reiver/go-iri/rfc3987"
)

const upperhex = "0123456789ABCDEF"

// PercentEncodeString returns the percent-encoded version of value as per IETF RFC-3987.
//
// ASCII unreserved characters (A-Z, a-z, 0-9, '-', '.', '_', '~') and Unicode characters in the ucschar ranges defined in IETF RFC 3987 Section 2.2  are left as-is.
// All other characters are percent-encoded as their UTF-8 byte sequences.
func PercentEncodeString(value string) string {
	var buffer [256]byte
	var p []byte = buffer[0:0]

	for index := 0; index < len(value); {
		r, size := utf8.DecodeRuneInString(value[index:])

		switch {
		case rfc3987.IsUnreserved(r) || rfc3987.IsUCSChar(r):
			p = append(p, value[index : index+size]...)
		default:
			for _, b := range []byte(value[index : index+size]) {
				p = append(p, '%')
				p = append(p, upperhex[b >> 4])
				p = append(p, upperhex[b & 0x0F])
			}
		}

		index += size
	}

	return string(p)
}
