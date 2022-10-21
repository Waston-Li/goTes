package pack

type Simpler interface {
	Get() int
	Set(int)
}
type Simple struct {
	i int
}

func (g *Simple) Get() int {
	return g.i
}
func (g *Simple) Set(i int) {
	g.i = i
}
func FI(it Simpler) int {
	it.Set(5)
	return it.Get()
}
