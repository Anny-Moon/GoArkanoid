package arkanoid

import "github.com/nsf/termbox-go"

var keyboardEventsChan = make(chan myKeyboardEvent)

//var s = makeScene(3, 5, 50, 20)
var s = scene{
    left: 3,
    top: 5,
    width: 50,
    height: 20,
}
//var p = makePaddle(10, 25, 10)
var p = paddle{
    left: 10,
    y: 25,
    length: 10,
    leftLimit: 3,
    rightLimit: 53,
}


func redraw(){
    s.draw()
    p.draw()
}

func Start(){
    termbox.Init()
    redraw()
    go listenToKeyboard(keyboardEventsChan)

mainloop:
    for{
        ch:=<-keyboardEventsChan
        switch ch.eventType {
            case MOVE:
                p.move(keyToDirection(ch.key))
                redraw()
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