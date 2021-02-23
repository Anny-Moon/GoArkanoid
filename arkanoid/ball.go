package arkanoid

import "fmt"
import "math"
//import "time"

const(
    dt float64 = 1
    cos float64 = 0.17 // 30 grad
    sin float64 = -0.98 // 30 grad
    //cos float64 = 0 // 90 grad
    //sin float64 = -1 // 90 grad
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

/*func (b *ball) reflectFromPaddleRightEnd(){
    b.vy = -b.vy
    vxNew := b.vx * cos + b.vy * sin
    vyNew := -b.vx * sin + b.vy * cos
    b.vx = vxNew
    b.vy = vyNew

}
*/

func (b *ball) reflectFromPaddleRightEnd(){
    vxNew := b.vx - 2.0*(cos*b.vx + sin*b.vy)*cos
    vyNew := b.vy - 2.0*(cos*b.vx + sin*b.vy)*sin
    b.vx = vxNew
    b.vy = vyNew
}

/*func (b *ball) reflectFromPaddleLeftEnd(){
    b.vy = -b.vy
    vxNew := b.vx * cos - b.vy * sin
    vyNew := b.vx * sin + b.vy * cos
    b.vx = vxNew
    b.vy = vyNew
}
*/

func (b *ball) reflectFromPaddleLeftEnd(){
    nx := -cos
    ny := sin
    tbprint(0,2,fmt.Sprintf("v: %f; %f", b.vx, b.vy))
    tbprint(0,3,fmt.Sprintf("n: %f; %f\t",nx, ny))
    vxNew := b.vx - 2.0*(nx*b.vx + ny*b.vy)*nx
    vyNew := b.vy - 2.0*(nx*b.vx + ny*b.vy)*ny
    b.vx = vxNew
    b.vy = vyNew
    tbprint(60,2,fmt.Sprintf("v: %f; %f", b.vx, b.vy))
    tbprint(60,3,fmt.Sprintf("n: %f; %f\t",nx, ny))
}

func (b *ball) reflectFromPaddle(nx, ny float64){
    tbprint(0,2,fmt.Sprintf("v: %f; %f", b.vx, b.vy))
    tbprint(0,3,fmt.Sprintf("n: %f; %f\t",nx, ny))
    vxNew := b.vx - 2.0*(nx*b.vx + ny*b.vy)*nx
    vyNew := b.vy - 2.0*(nx*b.vx + ny*b.vy)*ny
    b.vx = vxNew
    b.vy = vyNew
    tbprint(60,2,fmt.Sprintf("v: %f; %f", b.vx, b.vy))
    tbprint(60,3,fmt.Sprintf("n: %f; %f\t",nx, ny))
}

func (b *ball) correctVelocity(){
    v := math.Sqrt(b.vx*b.vx + b.vy*b.vy)
    if b.vx/math.Abs(b.vy) > 1.74{
        b.vx = v * 0.87
        b.vy = v * (-0.5)
    } else if b.vx/math.Abs(b.vy) < -1.74{
        b.vx = -v * 0.87
        b.vy = v * (-0.5)
    }
}

func(b *ball) updatePosition(p *paddle){
    xNew := b.x + b.vx * dt
    yNew := b.y + b.vy * dt
    
    if int(yNew) == p.y{
        for i := 0; i < p.length+1; i++ {
            if int(xNew) == p.left+i{
                
                if i < p.zone[0]{
                    tbprint(0,1, "left")
                    b.reflectFromPaddle(-cos, sin)
                } else if i >= p.zone[1]{
                    b.reflectFromPaddle(cos, sin)
                    tbprint(0,1, "righ")
                } else{
                    b.reflectFromPaddle(0, -1)
                    tbprint(0,1, "midd")
                }
                b.correctVelocity()
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

