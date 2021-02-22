package main

import(
     "example.com/arkanoid"
)

func main(){

    s := arkanoid.MakeScene(10, 10)
    arkanoid.DrawScene(s, 3,5)
}