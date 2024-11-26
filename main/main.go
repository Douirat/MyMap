package main

import (
	"github.com/Douirat/MyMap/logic"
)

func main() {
	cityMap := logic.NewGraph(6)
	a := logic.NewCity(1, "a")
	b := logic.NewCity(2, "b")
	c := logic.NewCity(3, "c")
	d := logic.NewCity(4, "d")
	e := logic.NewCity(5, "e")
	f := logic.NewCity(6, "f")
	cityMap.AddRoad(a, b)
	cityMap.AddRoad(a, d)
	cityMap.AddRoad(a, e)
	cityMap.AddRoad(b, d)
	cityMap.AddRoad(b, f)
	cityMap.AddRoad(c, f)
	cityMap.AddRoad(d, e)
	// cityMap.Display()
	cityMap.BFS(a)
}
