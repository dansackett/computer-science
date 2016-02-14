package graph

import q "../../data-structures/queue"

func BreadthFirstSearch(graph [][]int, source int) []map[string]int {
	data := make([]map[string]int, len(graph))

	for i, _ := range graph {
		data[i] = map[string]int{
			"distance": -1,
			"parent":   -1,
		}
	}

	data[source]["distance"] = 0

	queue := q.NewQueue()
	queue.Enqueue(source)

	for !queue.IsEmpty() {
		currentVertex := queue.Dequeue()

		for i, _ := range graph[currentVertex] {
			neighbor := graph[currentVertex][i]

			if data[neighbor]["distance"] < 0 {
				data[neighbor]["distance"] = data[currentVertex]["distance"] + 1
				data[neighbor]["parent"] = currentVertex
				queue.Enqueue(neighbor)
			}
		}
	}

	return data
}
