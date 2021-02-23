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
    zone [2]int
}


/*
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
*/

func (p *paddle) makeZones(){
    //zoneLength := int(p.length/3)
    zoneLength:=5
    tbprint(0,2,"hi")
    tbprint(0,2,string(zoneLength))
    p.zone[0] = zoneLength
    p.zone[1] = p.length - zoneLength
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