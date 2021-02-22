package arkanoid

import "github.com/nsf/termbox-go"

var keyboardEventsChan = make(chan myKeyboardEvent)

func Start(){
    s := makeScene(3, 5, 50, 20)
    drawScene(s)
    
    p := makePaddle(10, 25, 10)
    drawPaddle(p)
    
    go listenToKeyboard(keyboardEventsChan)

mainloop:
    for{
        ch:=<-keyboardEventsChan
        switch ch.eventType {
            //case MOVE:
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