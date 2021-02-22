package arkanoid


type direction int

const(
    LEFT direction = 1 + iota
    RIGHT
)

type paddle struct{
    left int
    y int
    length int
    leftLimit int
    rightLimit int
}



func makePaddle (left, y, length int) *paddle {
    p := &paddle{
        left : left,
        y : y,
        length : length,
        //leftLimit : leftLimit,
        //rightLimit : rightLimit,
    }
    return p
}

func (p *paddle)move(d direction){
    switch d{
    case LEFT:
        if p.left-1 > p.leftLimit{
            p.left--
        }
    case RIGHT:
        if p.left+p.length+1 < p.rightLimit{
            p.left++
        }
    }
}