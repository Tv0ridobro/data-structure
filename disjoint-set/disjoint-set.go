package disjoint_set

//https://en.wikipedia.org/wiki/Disjoint-set_data_structure
type DisjointSet struct {
	ranks   []int
	parents []int
	sizes   []int
}

func New(size int) *DisjointSet {
	d := &DisjointSet{
		ranks:   make([]int, size),
		parents: make([]int, size),
		sizes:   make([]int, size),
	}
	for i := 0; i < size; i++ {
		d.parents[i] = i
		d.sizes[i] = 1
	}
	return d
}

func (d *DisjointSet) Add() {
	d.ranks = append(d.ranks, 0)
	d.sizes = append(d.sizes, 1)
	d.parents = append(d.parents, len(d.parents))
}

func (d *DisjointSet) Get(x int) int {
	if d.parents[x] != x {
		d.parents[x] = d.Get(d.parents[x])
	}
	return d.parents[x]
}

func (d *DisjointSet) Size(x int) int {
	return d.sizes[d.Get(x)]
}

func (d *DisjointSet) Union(x int, y int) {
	x = d.Get(x)
	y = d.Get(y)
	if x == y {
		return
	}
	if d.ranks[x] == d.ranks[y] {
		d.ranks[x]++
	}
	if d.ranks[x] > d.ranks[y] {
		d.parents[y] = x
		d.sizes[x] += d.sizes[y]
	} else {
		d.parents[x] = y
		d.sizes[y] += d.sizes[x]
	}
}
