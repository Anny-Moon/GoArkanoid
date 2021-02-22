package arkanoid

type scene struct{
    width int
    height int
}

func MakeScene (width, height int) *scene{
    s := &scene{
        width: width,
        height: height,
    }
    return s
}