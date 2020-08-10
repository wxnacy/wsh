package main

import (
    "fmt"


    "github.com/wxnacy/wsh/wsh"
    // "github.com/nsf/termbox-go"
    "os"
    "os/exec"
    "strconv"
    "strings"
)

func LogFile(str ...string) {
    file, _ := os.OpenFile("wsh.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
    file.WriteString(strings.Join(str, " ") + "\n")
}


func main() {

    cmd := exec.Command("/bin/bash", "-c", "git status")
    _ = cmd
    // out, err := cmd.Output()
    // if err != nil {
        // panic(err)
    // }
    outs := make([]string, 0)
    for i := 0; i < 50; i++ {
        outs = append(outs, strconv.Itoa(i))
    }
    out := []byte(strings.Join(outs, "\n"))


    t, err := wsh.NewTerminal()
    if err != nil {
        panic(err)
    }
    defer t.Close()
    for {
        t.Run(out)
        e := t.PollEvent()
        if e.Ch > 0 {
            switch e.Ch {
                case 'q': {
                    os.Exit(0)
                }
                case 'l': {
                    t.MoveCursor(1, 0)
                }
                case 'h': {
                    t.MoveCursor(-1, 0)
                }
                case 'j': {
                    t.MoveCursor(0, 1)
                }
                case 'k': {
                    t.MoveCursor(0, -1)
                }
                case 'g': {
                    if t.E.PreCh == 'g' {
                        t.SetCursor(0, 0)
                    }
                }
                case 'G': {
                    t.SetCursor(0, t.Height - 1)
                }
                default: {
                    fmt.Println(string(rune(e.Ch)))
                }
            }
            // LogFile(termbox.ColorRed)
            LogFile(strconv.Itoa(t.Width) , strconv.Itoa(t.Height))
            LogFile(
                string(e.Ch), 
                strconv.Itoa(t.CursorX), 
                strconv.Itoa(t.CursorY),
                strconv.Itoa(t.OffsetX),
                strconv.Itoa(t.OffsetY),
            )
        } else {

        }
    }
    fmt.Println("Hello World")

}
