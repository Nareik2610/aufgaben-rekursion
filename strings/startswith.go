package strings

// StartsWith liefert true, falls der string s mit der Sequenz seq beginnt.
func StartsWith(s, seq string) bool {
	// TODO
	if seq == "" {
		return true
	}
	if len(s) < len(seq) {
		return false
	}

	return s[:len(seq)] == seq

}
