package main

import (
    "fmt"
    "os"
    "os/exec"
    "time"
    "bufio"
    "log"
    "crypto/sha512"
//    "github.com/howeyc/gopass"
)
func clearScreen () {
    clearScreen := exec.Command("clear")
    clearScreen.Stdout = os.Stdout
    clearScreen.Run()
}

func clearClipboard () {
    clearClipboard := exec.Command("/usr/bin/pbcopy", "<", "/dev/null")
    clearClipboard.Stdout = os.Stdout
    clearClipboard.Run()
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    line, err  := reader.ReadString('\n')
    if err != nil {
        log.Println(err)
        }
    clearScreen()

    //line := gopass.GetPasswd()
    b := sha512.Sum512([]byte(line))
    fmt.Println(fmt.Sprintf("%x", b))

    timer := time.NewTimer(10 * time.Second)
    <-timer.C

    clearScreen()
    clearClipboard()
}
