package arkanoid

//import "math"
//import "time"

const(
    dt float64 = 1
)

type ball struct{
    x int
    y int
    leftLimit int
    rightLimit int
    topLimit int
    bottomLimit int
    vx float64
    vy float64
}

func (b *ball) inverseVX(){
    b.vx = -b.vx
}
func (b *ball) inverseVY(){
    b.vy = -b.vy
}


func(b *ball) updatePosition(p *paddle){
    //var r float64 = math.Sqrt(float64(b.x*b.x + b.y*b.y))
    xNew := int(float64(b.x) + b.vx * dt)
    yNew := int(float64(b.y) + b.vy * dt)
    
    if yNew < b.topLimit-1{
        b.inverseVY()
        xNew = int(float64(b.x) + b.vx * dt)
        yNew = int(float64(b.y) + b.vy * dt)
    }
    
    if yNew == p.y{
        for i := 0; i < p.length+1; i++ {
            if xNew == p.left+i{
                b.inverseVY()
                xNew = int(float64(b.x) + b.vx * dt)
                yNew = int(float64(b.y) + b.vy * dt)
            }
        }
    }
    
    if xNew < b.leftLimit+1{
        b.inverseVX()
        xNew = int(float64(b.x) + b.vx * dt)
        yNew = int(float64(b.y) + b.vy * dt)
    }
    
    if xNew > b.rightLimit-2{
        b.inverseVX()
        xNew = int(float64(b.x) + b.vx * dt)
        yNew = int(float64(b.y) + b.vy * dt)
    }
    
    b.x = xNew
    b.y = yNew
}

