package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"unicode"

	"github.com/pkg/profile"
	// "runtime/pprof"
)

func readbyte(r io.Reader) (rune, error) {
	var buf [1]byte
	_, err := r.Read(buf[:])
	return rune(buf[0]), err
}

func calcWord() {
	f, err := os.Open("moby.txt")
	if err != nil {
		log.Fatalf("could not open file %q: %v", err)
	}

	words := 0
	inword := false
	for {
		r, err := readbyte(f)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("could not read file %q: %v", os.Args[1], err)
		}
		if unicode.IsSpace(r) && inword {
			words++
			inword = false
		}
		inword = unicode.IsLetter(r)
	}
	fmt.Printf("%d words\n", words)

}

func main() {
	//defer profile.Start().Stop()
	defer profile.Start(profile.MemProfile).Stop()
	calcWord()
}

// b := bufio.NewReader(f)
// defer profile.Start(profile.MemProfile).Stop()
