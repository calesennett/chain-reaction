package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	chains := make(map[string]string)
	var wordChain []string
	chainLength, _ := strconv.Atoi(os.Args[1])
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		firstPhrase := strings.Split(scanner.Text(), ":")[0]
		secondPhrase := strings.Split(scanner.Text(), ":")[1]
		chains[firstPhrase] = secondPhrase
	}
	for fPhrase, lPhrase := range chains {
		// add initial phrases to chain
		wordChain = append(wordChain, fPhrase, lPhrase)

		// append nextPhrase as long as chain length is not satisifed
		for i := 0; i < chainLength-2; i++ {
			nextPhrase, err := nextPhrase(wordChain, chains)
			if err != nil {
				break
			}
			wordChain = append(wordChain, nextPhrase)
		}
		if len(wordChain) == chainLength {
			break
		}
		// no chain of correct length found, start again
		wordChain = []string{}
	}
	fmt.Println(wordChain)
}

func nextPhrase(chain []string, chains map[string]string) (string, error) {
	// get last word of current chain
	lastWord := strings.Fields(chain[len(chain)-1])[1]
	for nPhrase, _ := range chains {
		if lastWord == strings.Fields(nPhrase)[0] {
			return nPhrase, nil
		}
	}
	return "", errors.New("No phrase found")
}
