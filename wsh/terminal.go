package wsh

import (
    "github.com/nsf/termbox-go"
)

type TerminalText struct {
    CursorX, CursorY int
    Value rune
    Fg termbox.Attribute
}

type Line struct {
    Line int
}

type Output struct {
    Line int
}

type Event struct {
    PreCh rune
    Ch rune
    // E termbox.Event
}

func StringToTT(str string) []TerminalText {
    values := []rune(str)
    res := make([]TerminalText, 0)
    x := 0
    y := 0
    for _, d := range values {
        tt := TerminalText{CursorX: x, CursorY: y, Value: d}
        x++
        if d == 10 {
            y++
            x = 0
        } else {
            res  = append(res, tt)
        }
    }
    return res
}

type Terminal struct {
    Width, Height    int
	CursorX, CursorY int
    WindowX, WindowY int
	OffsetX, OffsetY int
    E *Event
}

func NewTerminal() (*Terminal, error){
    err := termbox.Init()
    if err != nil {
        return nil, err
    }

    w, h := termbox.Size()

    return &Terminal{Width: w, Height: h, E: &Event{}}, nil
}

func (t *Terminal) Run(input []byte) {

    termbox.Clear(termbox.ColorWhite, termbox.ColorDefault)
    str := string(input)
    tts := StringToTT(str)
    for _, d := range tts {
        termbox.SetCell(d.CursorX, d.CursorY, d.Value, termbox.ColorDefault, termbox.ColorDefault)
    }
    termbox.SetCursor(t.CursorX, t.CursorY)
    termbox.Flush()
}

func (t *Terminal) MoveCursor(x, y int) {
    cx := t.CursorX + x
    cy := t.CursorY + y
    var flagX, flagY bool
    flagX = true
    flagY = true
    if cx < 0 || cx == t.CursorX || cx >= t.Width {
        flagX = false
    }

    if cy < 0 || cy == t.CursorY || cy >= t.Height {
        flagY = false
    }

    if flagX {
        t.CursorX = cx
    }

    if flagY {
        t.CursorY = cy
    }

    if flagX || flagY {
        termbox.SetCursor(t.CursorX, t.CursorY)
        termbox.Flush()
    }

    t.OffsetX = cx
    t.OffsetY = cy
}

func (t *Terminal) SetCursor(x, y int) {
    t.CursorX = x
    t.CursorY = y

    termbox.SetCursor(t.CursorX, t.CursorY)
    termbox.Flush()
}

func (t *Terminal) Close() {
    termbox.Close()
}

func (t *Terminal) PollEvent() termbox.Event{
    e := termbox.PollEvent()
    t.E.PreCh = t.E.Ch
    t.E.Ch = e.Ch
    return e
}
