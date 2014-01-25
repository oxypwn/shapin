package main

import (
    "fmt"
    "os"
    "os/exec"
    "time"
    "bufio"
    "log"
    "crypto/sha512"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    line, err  := reader.ReadString('\n')
    if err != nil {
        log.Println(err)
        }
    b := sha512.Sum512([]byte(line))
    fmt.Println(fmt.Sprintf("%x", b))

    timer := time.NewTimer(10 * time.Second)
    <-timer.C


    clearClipBoard := exec.Command("/usr/bin/pbcopy", "<", "/dev/null")
    clearClipBoard.Stdout = os.Stdout
    clearClipBoard.Run()
    clearScreen := exec.Command("clear")
    clearScreen.Stdout = os.Stdout
    clearScreen.Run()

}
