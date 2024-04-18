package asciiArt

import (
	"fmt"
)

func PrintAscii(art []string) string {
	if len(art) == 0 {
		fmt.Println()
	}
	str := ""
	for _, v := range art {
		str += v + "\n"
	}
	return str
}
