package iri

import (
	"strings"
	"unicode/utf8"
)


// IsUCSChar returns whether a rune is in the ucschar ranges as per IETF RFC-3987 Section 2.2:
//
//	ucschar = %xA0-D7FF / %xF900-FDCF / %xFDF0-FFEF
//	        / %x10000-1FFFD / %x20000-2FFFD / %x30000-3FFFD
//	        / %x40000-4FFFD / %x50000-5FFFD / %x60000-6FFFD
//	        / %x70000-7FFFD / %x80000-8FFFD / %x90000-9FFFD
//	        / %xA0000-AFFFD / %xB0000-BFFFD / %xC0000-CFFFD
//	        / %xD0000-DFFFD / %xE1000-EFFFD
func IsUCSChar(r rune) bool {
	switch {
	case 0xA0 <= r && r <= 0xD7FF:
		return true
	case 0xF900 <= r && r <= 0xFDCF:
		return true
	case 0xFDF0 <= r && r <= 0xFFEF:
		return true
	case 0x10000 <= r && r <= 0x1FFFD:
		return true
	case 0x20000 <= r && r <= 0x2FFFD:
		return true
	case 0x30000 <= r && r <= 0x3FFFD:
		return true
	case 0x40000 <= r && r <= 0x4FFFD:
		return true
	case 0x50000 <= r && r <= 0x5FFFD:
		return true
	case 0x60000 <= r && r <= 0x6FFFD:
		return true
	case 0x70000 <= r && r <= 0x7FFFD:
		return true
	case 0x80000 <= r && r <= 0x8FFFD:
		return true
	case 0x90000 <= r && r <= 0x9FFFD:
		return true
	case 0xA0000 <= r && r <= 0xAFFFD:
		return true
	case 0xB0000 <= r && r <= 0xBFFFD:
		return true
	case 0xC0000 <= r && r <= 0xCFFFD:
		return true
	case 0xD0000 <= r && r <= 0xDFFFD:
		return true
	case 0xE1000 <= r && r <= 0xEFFFD:
		return true
	default:
		return false
	}
}
