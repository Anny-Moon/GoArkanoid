package arkanoid

import(
    //"fmt"
//    "time"
    "github.com/nsf/termbox-go"
)

const (
//    defaultColor = termbox.ColorDefault
//    defaultColor = termbox.ColorBlack

    bgColor     = termbox.ColorBlack
    fgColor     = termbox.ColorLightGray
    paddleColor = termbox.ColorCyan
    
    snakeColor   = termbox.ColorGreen
)

func (s *scene)draw(){
    
    
    termbox.Clear(bgColor, bgColor)
    //tbprint(left, top-1, defaultColor, defaultColor, "Snake Game")
    for i := s.top; i < s.top+s.height; i++ {
        termbox.SetCell(s.left-1, i, '|', fgColor, fgColor)
        termbox.SetCell(s.left+s.width, i, '|', fgColor, fgColor)
    }
    
    for i := s.left; i < s.left+s.width; i++ {
        termbox.SetCell(i, s.top-1, '_', fgColor, bgColor)
        termbox.SetCell(i, s.top+s.height, '^', fgColor, bgColor)
    }
    
}

func (p *paddle)draw(){
    for i := p.left; i < p.left+p.length; i++ {
        termbox.SetCell(i, p.y, 'H', fgColor, paddleColor)
    }
    termbox.Flush()
    //time.Sleep(time.Second * 5);termbox.Close()
}

func (b *ball)draw(){
    termbox.SetCell(b.x, b.y, 'o', fgColor, bgColor)
    termbox.Flush()
}
