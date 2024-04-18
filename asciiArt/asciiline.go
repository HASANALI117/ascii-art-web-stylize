package asciiArt

func AsciiLine(str string, font string) ([]string, error) {
	fileName := "banner/" + font + ".txt"
	art := []string{}
	for _, v := range str {
		startLine, err := StartLineCalc(v)
		if err != nil {
			// log.Println(err)
			return nil, err
		}
		art = AppendAscii(art, ReadLines(fileName, startLine, 8))
	}
	return art, nil
}
