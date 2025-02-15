package strings

// StartsWith liefert true, falls der string s mit der Sequenz seq beginnt.
func StartsWith(s, seq string) bool {
	// TODO
	if seq == "" {
		return true
	}
	if s == "" {
		return false
	}
	return s[0] == seq[0]

}
