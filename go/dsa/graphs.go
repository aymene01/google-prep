package dsa


var (
	DirectedGraph = make(map[int][]int, 0)
)

type Graph struct {}

func BuildGraph[T comparable](edges [][]T) map[T][]T{
	graph := make(map[T][]T)

	for _, edge := range edges {
		src, dst := edge[0], edge[1]
		graph[src] = append(graph[src], dst)
	}

	return graph
}

func Dfs[T comparable](graph map[T][]T, src T, dst T, visited map[T]bool) bool {
	if visited[src] {
		return false
	}

	if src == dst {
		return true
	}

	visited[src] = true

	for _, neighbor := range graph[src] {
		if Dfs(graph, neighbor, dst, visited) {
			return true
		}
	}

	return false
}