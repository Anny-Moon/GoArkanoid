package arkanoid

import (
    "time"
    "fmt"
    "github.com/nsf/termbox-go"
)


var keyboardEventsChan = make(chan myKeyboardEvent)

//var s = makeScene(3, 5, 50, 20)
var s = scene{
    left: 3,
    top: 5,
    width: 60,
    height: 20,
}
//var p = makePaddle(10, 25, 10)
var p = paddle{
    left: 10,
    y: s.top+s.height-1,
    length: 10,
    leftLimit: s.left-1,
    rightLimit: s.left+s.width+1,
}

var b = ball{
    x: p.left + 1,
    y: p.y-1,
    leftLimit: s.left-1,
    rightLimit: s.left+s.width+1,
    topLimit: s.top + 1,
    bottomLimit: s.top+s.height+1,
    vx: 1.5,
    vy: -1,
}

func redraw(){
    s.draw()
    p.draw()
    b.draw()
}

func loopForBall(){
    for {
        time.Sleep(time.Millisecond * 200)
        b.updatePosition(&p)
        redraw()
    }
}

func Start(){
    termbox.Init()
    redraw()
    go listenToKeyboard(keyboardEventsChan)
    go loopForBall()

mainloop:
   for{
//       select{
        ch:=<-keyboardEventsChan
            switch ch.eventType {
                case MOVE:
                    p.move(keyToDirection(ch.key))
                    redraw()
                case END:
                    break mainloop
            
            
            }
//        default:
//                b.updatePosition()
//                redraw()
//                time.Sleep(time.Second)
//        }
    }
    
    
/*    for{
        //b.updatePosition()
        redraw()
        
        ch:=<-keyboardEventsChan
            switch ch.eventType {
                case MOVE:
                    p.move(keyToDirection(ch.key))
                    redraw()
                case END:
                    break mainloop
                }
                
        
    }*/
    
    termbox.Close()
    fmt.Println(b.vx, b.vy)
}