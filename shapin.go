package main

import (
    "fmt"
    "os"
    "bufio"
    "crypto/sha512"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    for {
        line, err := reader.ReadString('\n')

        if err != nil {
            break 
            }
        

        b := sha512.Sum512([]byte(line))
        fmt.Println(fmt.Sprintf("%x", b))
    }
}
