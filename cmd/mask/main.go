package main

import (
    "bufio"
    "strings"
    "os"
    "fmt"
)


// Read in standard input and write mask to standard output
// Mask output format -> stdin[0]-stdin[3] rest of stdin is '*' for each character. suffix output with length of stdin.
// Ex: randomsecretissecret -> rand****************[57]
func main(){
    stdInReader := bufio.NewReader(os.Stdin)
    text, err := stdInReader.ReadString('\n')
    if err != nil {
        panic(err)
    }

    if len(text) >= 3 {
        fmt.Printf("%s%s[%d]\n", text[:3], strings.Repeat("*", len(text) - 3), len(text))
    } else {
        fmt.Println(text)
    }
}
