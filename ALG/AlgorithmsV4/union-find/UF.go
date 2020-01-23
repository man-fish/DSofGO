package UF

type IUF interface {
	Union(p int, q int)
	Connected(p int, q int) bool
	Count() int
}

type UF struct {
	id		[]int
	count	int
}

func New(n int) *UF {
	id := make([]int,n)
	for i := 0; i < n; i++ {
		id[i] = i
	}
	return &UF{id,n}
}

func (u *UF) Union(p int, q int) {

}
