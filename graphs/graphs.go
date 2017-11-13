package graphs

type Graphs interface {
	GetAdjacentVerticesNodesForVertex(vertex int) []int

	GetWeightedOfEdge(vertexOne int, vertexTwo int)

	GetNumberOfVertices() int

	GetIndegreeForVertex(vertex int) int

	GetOutdegreeForVertex(vertex int) int

	AddEdge(vertexOne int, vertexTwo int)

	AddEdgeWithWeight(vertexOne int, vertexTwo int, weight int)
}
