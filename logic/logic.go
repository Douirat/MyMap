package logic

import "fmt"

// Create a vertex to represent a city:
type City struct {
	Id   int
	Name string
}

// Create a structure to reprsent a Graph:
type CityRoads struct {
	Size int
	Map  [][]int
}

// Instantiate a new City:
func NewCity(id int, name string) *City {
	city := new(City)
	city.Id = id
	city.Name = name
	return city
}

// Instantiate the Graph:
func MapInstance(size int) *CityRoads {
	cityMap := new(CityRoads)
	cityMap.Size = size
	cityMap.Map = make([][]int, cityMap.Size)
	for i := 0; i < cityMap.Size; i++ {
		cityMap.Map[i] = make([]int, cityMap.Size)
	}
	return cityMap
}

// Display the Map graph:
func (roads *CityRoads) DisplayMap() {
	fmt.Printf("  ")
	for i := 0; i < roads.Size; i++ {
		fmt.Printf("%d ", i+1)
	}
	fmt.Printf("\n")
	for x:=0; x<roads.Size; x++{
		fmt.Printf("%d ", x+1)
		for z:=0; z<roads.Size; z++{
			fmt.Printf("%d ", roads.Map[x][z])
		}
		fmt.Println()
	}
}
