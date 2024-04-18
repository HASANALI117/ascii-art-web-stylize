package asciiArt

func AppendAscii(art []string, char []string) []string {
	if len(art) == 0 {
		return char
	}
	for i, v := range char {
		art[i] = art[i] + v
	}
	return art
}
