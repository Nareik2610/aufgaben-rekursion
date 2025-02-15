package strings

// Contains prüft, ob der String s die Sequenz seq enthält.
func Contains(s, seq string) bool {
	// TODO
	if seq == "" {
		return true
	}
	if len(s) < len(seq) {
		return false
	}
	if StartsWith(s, seq) {
		return true
	}

	return Contains(s[1:], seq)
}
