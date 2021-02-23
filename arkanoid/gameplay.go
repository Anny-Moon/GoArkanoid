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
    width: 50,
    height: 50,
}
//var p = makePaddle(10, 25, 10)
var p = paddle{
    left: 10,
    y: s.top+s.height-1,
    length: 20,
    leftLimit: s.left-1,
    rightLimit: s.left+s.width+1,
}

var b = ball{
    x: float64(p.left + 1),
    y: float64(p.y-1),
    leftLimit: s.left-1,
    rightLimit: s.left+s.width+1,
    topLimit: s.top + 1,
    bottomLimit: s.top+s.height+1,
    vx: 3,
    vy: -2,
}

func redraw(){
    s.clearScene()
    s.draw()
    p.draw()
    b.draw()
}

func loopForBall(){
    for {
        time.Sleep(time.Millisecond * 100)
        b.updatePosition(&p)
        redraw()
    }
}

func Start(){
    p.makeZones()
    termbox.Init()
        
    termbox.Clear(bgColor, bgColor)
    tbprint(0,0, "hello")
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
    //tbprint(0,0, "bye")
    //time.Sleep(time.Second)
    
    termbox.Clear(bgColor, bgColor)
    termbox.Close()
    fmt.Println(b.vx, b.vy)
}