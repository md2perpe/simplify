package simplify

type Mesh struct {
	Triangles []*Triangle
}

func NewMesh(triangles []*Triangle) *Mesh {
	return &Mesh{triangles}
}

func (m *Mesh) SaveBinarySTL(path string) error {
	return SaveBinarySTL(path, m)
}

func (m *Mesh) SaveAsciiSTL(path string) error {
	return SaveAsciiSTL(path, m)
}

func (m *Mesh) Simplify(factor float64) *Mesh {
	return Simplify(m, factor)
}
