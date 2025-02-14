package strings

//Bestimme die Länge von s

func Length(s string) int {
	// TODO
	if s == "" {
		return 0
	}
	return 1 + Length(s[1:])

}
