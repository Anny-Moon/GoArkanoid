package arkanoid

import "github.com/nsf/termbox-go"


type myKeyboardEventType int

const(
    MOVE myKeyboardEventType = 1 + iota
    END
)

type myKeyboardEvent struct{
    eventType   myKeyboardEventType
    key         termbox.Key
}

func keyToDirection(k termbox.Key) direction {
    switch k{
    case termbox.KeyArrowLeft:
        return LEFT
    case termbox.KeyArrowRight:
        return RIGHT
    default:
        return 0
    }
}

func listenToKeyboard(evChan chan myKeyboardEvent){
    termbox.SetInputMode(termbox.InputEsc)
    for {
        switch ev := termbox.PollEvent(); ev.Type {
        case termbox.EventKey:
            switch ev.Key {
            case termbox.KeyArrowLeft:
                evChan <- myKeyboardEvent{eventType: MOVE, key: ev.Key}
            
            case termbox.KeyArrowRight:
                evChan <- myKeyboardEvent{eventType: MOVE, key: ev.Key}
            
            case termbox.KeyEsc:
                evChan <- myKeyboardEvent{eventType: END, key: ev.Key}
            
            }
        case termbox.EventError:
            panic(ev.Err)
        }
    }
        
}