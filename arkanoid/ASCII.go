package arkanoid

import(
    //"fmt"
//    "time"
    "github.com/nsf/termbox-go"
    "github.com/mattn/go-runewidth"
)

const (
//    defaultColor = termbox.ColorDefault
//    defaultColor = termbox.ColorBlack

    bgColor     = termbox.ColorBlack
    fgColor     = termbox.ColorLightGray
    paddleColor = termbox.ColorCyan
    redColor = termbox.ColorRed
    yellowColor = termbox.ColorYellow
    snakeColor   = termbox.ColorGreen
)


func (s *scene)draw(){
    
    for i := s.top; i < s.top+s.height; i++ {
        termbox.SetCell(s.left-1, i, '|', fgColor, fgColor)
        termbox.SetCell(s.left+s.width, i, '|', fgColor, fgColor)
    }
    
    for i := s.left; i < s.left+s.width; i++ {
        termbox.SetCell(i, s.top-1, '_', fgColor, bgColor)
        termbox.SetCell(i, s.top+s.height, '^', fgColor, bgColor)
    }
    
}

func (s *scene)clearScene(){
    for j := s.left-1; j < s.left+s.width +1; j++ {
        for i := s.top-1; i < s.top+s.height+1; i++ {
            termbox.SetCell(j, i, ' ', bgColor, redColor)
        }
    }
}

func (p *paddle)draw(){
    for i := p.left; i < p.left+p.length; i++ {
        termbox.SetCell(i, p.y, 'H', fgColor, paddleColor)
    }
    
    for i := p.zone[0]; i < p.zone[1]; i++ {
        termbox.SetCell(i+p.left, p.y, 'H', fgColor, yellowColor)
    }
    
    termbox.Flush()
    //time.Sleep(time.Second * 5);termbox.Close()
}

func (b *ball)draw(){
    termbox.SetCell(int(b.x), int(b.y), 'o', fgColor, bgColor)
    termbox.Flush()
}


func tbprint(x, y int,  msg string) {
    for _, c := range msg {
        termbox.SetCell(x, y, c, fgColor, redColor)
        x += runewidth.RuneWidth(c)
    }
}