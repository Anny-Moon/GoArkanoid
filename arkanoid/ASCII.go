package arkanoid

import(
    "fmt"
    "time"
    "github.com/nsf/termbox-go"
)

const (
//    defaultColor = termbox.ColorDefault
//    defaultColor = termbox.ColorBlack

    bgColor     = termbox.ColorBlack
    fgColor     = termbox.ColorLightGray
    
    snakeColor   = termbox.ColorGreen
)

func DrawScene (s *scene, left, top int){
    fmt.Printf("Hello1\n")
    
    var (
    IsInit bool = false
    )
    termbox.Init()
    fmt.Println(IsInit)
    
    termbox.Clear(bgColor, bgColor)
    //tbprint(left, top-1, defaultColor, defaultColor, "Snake Game")
    for i := top; i < top+s.height; i++ {
        termbox.SetCell(left-1, i, '|', fgColor, bgColor)
        termbox.SetCell(left+s.width, i, '|', fgColor, bgColor)
    }
    
    for i := left; i < left+s.width; i++ {
        termbox.SetCell(i, top-1, '_', fgColor, bgColor)
        termbox.SetCell(i, top+s.height, '^', fgColor, bgColor)
    }
    termbox.Flush()
    time.Sleep(time.Second * 5);termbox.Close()
}