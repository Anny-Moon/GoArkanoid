package arkanoid

type scene struct{
    left int
    top int
    width int
    height int
    
}

func makeScene (left, top, width, height int) *scene{
    s := &scene{
        left: left,
        top: top,
        width: width,
        height: height,
    }
    return s
}