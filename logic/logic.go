package logic

import "fmt"

// Create a vertex to represent a city:
type City struct {
	Id   int
	Name string
	Next *City
}

// Create a structure to reprsent a Graph:
type Cities struct {
	Cities []*City
}

// Instantiate a new a graph:
func NewGraph(size int) *Cities {
	graph := new(Cities)
	graph.Cities = make([]*City, size)
	return graph
}

// Instantiate a new structure to represent the city:
func NewCity(id int, name string) *City {
	city := new(City)
	city.Id = id
	city.Name = name
	return city
}

// Add a road between tow cities:
func (cities *Cities) AddRoad(dep, des *City) {
	if cities.Cities[dep.Id-1] == nil {
		dep.Next = des
		cities.Cities[dep.Id-1] = dep
	} else {
		temp := cities.Cities[dep.Id-1]
		for temp.Next != nil {
			temp = temp.Next
		}
		temp.Next = des
	}
}

// Display the graph:
func (cities *Cities) Display() {
	for i := 0; i < len(cities.Cities); i++ {
		fmt.Printf("%c.%d -----> ", i+97, i+1)
		temp := cities.Cities[i]
		for temp != nil {
			fmt.Printf("%d, %s -> ", temp.Id, temp.Name)
			temp = temp.Next
		}
		fmt.Printf("\n")
	}
}

// Application of the breadth first search:
func (cities *Cities)BFS(dep *City) {
// Keep track of visited nodes: 
visited := make(map[*City]bool)
// Initialize a node with the depart node:
Queue := []*City{dep}
fmt.Println("BFS Traversal:")

for len(Queue) > 0 {
	// Dequeue the front of the queue:
	node := Queue[0]
	Queue = Queue[1:]

	// skip if node have already been visited:
	if visited[node] {
		continue
	}

	// Mark as visited an process the node:
	visited[node] = true
	fmt.Printf("We are at point %s and node %d\n", node.Name, node.Id)
	// Enqueue all non visited neigbors
	for node != nil {
		if !visited[node] {
			Queue = append(Queue, node)
		}
		node = node.Next
	}
}
}