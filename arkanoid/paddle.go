package arkanoid


type direction int

const(
    LEFT direction = 1 + iota
    RIGHT
)

type paddle struct{
    left int
    right int
    y int
    length int
}



func makePaddle (left, y, length int) *paddle {
    p := &paddle{
        left : left,
        right : left + length,
        y : y,
        length : length,
    }
    return p
}

func (p *paddle)move(d direction){
    switch d{
    case LEFT:
        p.left--
        p.right--
    case RIGHT:
        p.left++
        p.right++
    }
}