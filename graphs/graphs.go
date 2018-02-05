package graphs

type GraphType string

const (
	DIRECTED   GraphType = "DIRECTED"
	UNDIRECTED GraphType = "UNDIRECTED"
)

type Graphs interface {
	Init()

	AddEdge(vertexOne interface{}, vertexTwo interface{}) error

	AddEdgeWithWeight(vertexOne interface{}, vertexTwo interface{}, weight int) error

	RemoveEdge(vertexOne interface{}, vertexTwo interface{}) error

	HasEdge(vertexOne interface{}, vertexTwo interface{}) bool

	GetGraphType() GraphType

	GetAdjacentNodesForVertex(vertex interface{}) map[interface{}]bool

	GetWeightOfEdge(vertexOne interface{}, vertexTwo interface{}) (int, error)

	GetNumberOfVertices() int

	GetIndegreeForVertex(vertex interface{}) int
}
