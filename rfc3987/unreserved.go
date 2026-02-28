package iri

// IsUnreserved returns whether a rune is an ASCII unreserved character as per IETF RFC-3986 Section 2.3:
//
//	unreserved = ALPHA / DIGIT / "-" / "." / "_" / "~"
func IsUnreserved(r rune) bool {
	switch {
	case 'A' <= r && r <= 'Z':
		return true
	case 'a' <= r && r <= 'z':
		return true
	case '0' <= r && r <= '9':
		return true
	case r == '-', r == '.', r == '_', r == '~':
		return true
	default:
		return false
	}
}
