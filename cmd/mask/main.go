package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/kpderoshxyz/system-utils/internal/mask"
)

// Read in standard input and write mask to standard output
// Mask output format -> stdin[0]-stdin[2] rest of stdin is '*' for each character. suffix output with length of stdin.
// Ex: testthis -> tes*****[8]
func main(){
    stdInReader := bufio.NewReader(os.Stdin)
    text, err := stdInReader.ReadString('\n')
    if err != nil {
        panic(err)
    }

    fmt.Println(mask.Mask(text))
}
