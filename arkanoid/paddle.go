package arkanoid


type coord struct{
    x int
    y int
}

type paddle struct{
    left int
    right int
    y int
    length int
}

func MakePaddle (left, y, length int) *paddle {
    p := &paddle{
        left : left,
        right : left + length,
        y : y,
        length : length,
    }
    return p
}