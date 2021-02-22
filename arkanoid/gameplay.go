package arkanoid

import "github.com/nsf/termbox-go"

var keyboardEventsChan = make(chan myKeyboardEvent)

var s = makeScene(3, 5, 50, 20)
var p = makePaddle(10, 25, 10)

func Start(){
    termbox.Init()
    drawScene(s)
    drawPaddle(p)
    
    go listenToKeyboard(keyboardEventsChan)

mainloop:
    for{
        ch:=<-keyboardEventsChan
        switch ch.eventType {
            case MOVE:
                p.move(keyToDirection(ch.key))
                drawScene(s)
                drawPaddle(p)
            //    d := keyToDirection(e.key)
            //    g.arena.snake.changeDirection(d)
            //case RETRY:
            //    g.retry()
            case END:
                break mainloop
            }    
    }
    
    termbox.Close()
}