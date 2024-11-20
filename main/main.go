package main

import(
	"github.com/Douirat/MyMap/logic"
)

func main(){
	cityMap := logic.MapInstance(5)
	cityMap.DisplayMap()
}