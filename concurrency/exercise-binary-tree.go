package main

type Tree struct {
    Left  *Tree
    Value int
    Right *Tree
}

func insert(t *Tree, v int) *Tree {
    if t == nil {
        return &Tree{Value: v}
    }
    if v < t.Value {
        t.Left = insert(t.Left, v)
    } else {
        t.Right = insert(t.Right, v)
    }
    return t
}

func New(k int) *Tree {
    var t *Tree
    for i := 1; i <= 10; i++ {
        t = insert(t, i*k)
    }
    return t
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *Tree, ch chan int) {
	if t == nil{
		return
	}
	Walk(t.Left,ch)
	ch <- t.Value
	Walk(t.Right,ch)
}

func Walker(t *Tree, ch chan int) {
    // initiate walking over tree t with corresponding channel
    Walk(t, ch)
    // close channel so that range can finish
    close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *Tree) bool {
	if t1 == nil || t2 == nil {return false}
	ch1,ch2 := make(chan int),make(chan int)
	go Walker(t1,ch1)
	go Walker(t2,ch2)
	for v1 := range ch1 {
		v2,ok := <- ch2  
		if !ok {
			return false
		}
		if v1 != v2 {
			return false
		}
	}
	return true
}
