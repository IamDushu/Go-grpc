package main

import (
	"fmt"
)

func printUTF8Encoding(s string) {
	fmt.Printf("Input: %q\n\n", s)
	for i, r := range s {
		utf8Bytes := []byte(string(r))
		unicodeBinary := fmt.Sprintf("%016b", r) // 16-bit binary representation

		fmt.Printf("Character #%d: %q\n", i+1, r)
		fmt.Printf("→ Unicode:  U+%04X (decimal: %d)\n", r, r)
		fmt.Printf("→ Binary (Unicode): %s\n", insertEvery4Bits(unicodeBinary))
		fmt.Printf("→ UTF-8 Bytes: ")

		for _, b := range utf8Bytes {
			fmt.Printf("%d ", b)
		}

		fmt.Print("\n→ Hex:       ")
		for _, b := range utf8Bytes {
			fmt.Printf("0x%X ", b)
		}

		fmt.Print("\n→ Binary:    ")
		for _, b := range utf8Bytes {
			fmt.Printf("%08b ", b)
		}

		fmt.Println("\n---")
	}
}

// Optional: pretty-print binary by grouping 4 bits
func insertEvery4Bits(s string) string {
	out := ""
	for i, c := range s {
		if i > 0 && (len(s)-i)%4 == 0 {
			out += " "
		}
		out += string(c)
	}
	return out
}

func main() {
	// printUTF8Encoding("త")
	// Try with more examples:
	// printUTF8Encoding("😊")
	printUTF8Encoding("hi")
	// printUTF8Encoding("తెలుగు")
}
