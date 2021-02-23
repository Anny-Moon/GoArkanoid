package arkanoid

//import "math"
//import "time"

const(
    dt float64 = 1
    //cos = 0.87 // 30 grad
    //sin = 0.5 // 30 grad
    cos = 0 // 30 grad
    sin = 1 // 30 grad
)

type ball struct{
    x float64
    y float64
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

func (b *ball) reflectFromPaddleMiddle(){
    b.vy = -b.vy
}

func (b *ball) reflectFromPaddleRightEnd(){
    vxNew := b.vx * cos + b.vy * sin
    vyNew := -b.vx * sin + b.vy * cos
    b.vx = vxNew
    b.vy = -vyNew

}
func (b *ball) reflectFromPaddleLeftEnd(){
    vxNew := b.vx * cos - b.vy * sin
    vyNew := b.vx * sin + b.vy * cos
    b.vx = vxNew
    b.vy = - vyNew

}


func(b *ball) updatePosition(p *paddle){
    xNew := b.x + b.vx * dt
    yNew := b.y + b.vy * dt
    
    if int(yNew) == p.y{
        for i := 0; i < p.length+1; i++ {
            if int(xNew) == p.left+i{
                
                if i < p.zone[0]{
                    tbprint(0,1, "left")
                    b.reflectFromPaddleLeftEnd()
                } else if i >= p.zone[1]{
                    b.reflectFromPaddleRightEnd()
                    tbprint(0,1, "righ")
                } else{
                    b.reflectFromPaddleMiddle()
                    tbprint(0,1, "midd")
                }
            }
        }
    } else if int(yNew) < b.topLimit-1{
        b.inverseVY()
    }
    
    if int(xNew) < b.leftLimit+1{
        b.inverseVX()
    } else if int(xNew) > b.rightLimit-2{
        b.inverseVX()
    }
    
    xNew = b.x + b.vx * dt
    yNew = b.y + b.vy * dt
    b.x = xNew
    b.y = yNew
}

