package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
    pic := make([][]uint8, dy)
    for i := range pic {
        pic[i] = make([]uint8, dx)
        
        for j := range pic[i] {
        	pic[i][j] = uint8(i) * uint8(j)
        }
    }
    return pic
}

func main() {
    pic.Show(Pic)
}
