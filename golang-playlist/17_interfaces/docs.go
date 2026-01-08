package main

import (
    "fmt"
    "math"
)
//To implement an interface in Go, we just need to implement all the methods in the interface. 

type geometry interface {
    area() float64
    perim() float64
}

type rect struct {
    width, height float64
}
type circle struct {
    radius float64
}

//can see rect and circle have implemented all methods of geometry interface
func (r rect) area() float64 {
    return r.width * r.height
}
func (r rect) perim() float64 {
    return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}

//If a variable has an interface type, then we can call methods that are in the named interface. 
// Here’s a generic measure function taking advantage of this to work on any geometry.
func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}

func detectCircle(g geometry) {
    if c, ok := g.(circle); ok {
        fmt.Println("circle with radius", c.radius)
    }
	//g.(circle) //checking type of g is circle ( like done in switch , i.(type))
}

func main() {
    r := rect{width: 3, height: 4}
    c := circle{radius: 5}

    measure(r)
    measure(c)

    detectCircle(r)
    detectCircle(c)
}