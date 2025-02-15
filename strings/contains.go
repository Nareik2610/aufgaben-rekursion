package strings

// Contains prüft, ob der String s die Sequenz seq enthält.
func Contains(s, seq string) bool {
	// TODO
	if seq == "" {
		return true
	}
	if s == "" {
		return false
	}

	return StartsWith(s[1:], seq[1:]) || Contains(s[1:], seq)
}
