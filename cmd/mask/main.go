package main

import (
    "bufio"
    "strings"
    "os"
    "fmt"
)

// Read in standard input and write mask to standard output
// Mask output format -> stdin[0]-stdin[3] rest of stdin is '*' for each character. suffix output with length of stdin.
// Ex: testthis -> tes*****[8]
func main(){
    stdInReader := bufio.NewReader(os.Stdin)
    text, err := stdInReader.ReadString('\n')
    if err != nil {
        panic(err)
    }

    fmt.Println(mask.Mask(text))
}
