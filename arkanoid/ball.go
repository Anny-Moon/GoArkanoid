package arkanoid

import "math"
import "time"

const(
    dr float64 = 0.2
)

type ball struct{
    x int
    y int
    leftLimit int
    rightLimit int
    topLimit int
    bottomLimit int
}

func(b *ball) updatePosition(){
    time.Sleep(time.Second)
    var r float64 = math.Sqrt(float64(b.x*b.x + b.y*b.y))
    xNew := int(float64(b.x) + dr * float64(b.x)/r)
    yNew := int(float64(b.y) - dr * float64(b.y)/r)
    b.x = xNew
    b.y = yNew
}

