package main

import (
	"fmt"
	"reflect"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)

}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	if t1 == nil || t2 == nil {
		return false
	}
	ch1 := make(chan int, 10)
	// 用来收集ch1结果
	res1 := make([]int, 0)
	Walk(t1, ch1)
	close(ch1)
	for v := range ch1 {
		res1 = append(res1, v)
	}

	ch2 := make(chan int, 10)
	// 用来收集ch2结果
	res2 := make([]int, 0)
	Walk(t2, ch2)
	close(ch2)
	for v := range ch2 {
		res2 = append(res2, v)
	}
	return reflect.DeepEqual(res1, res2)
}

func main() {
	tree1 := tree.Tree{nil, 3, nil}
	tree21 := tree.Tree{nil, 1, nil}
	tree22 := tree.Tree{nil, 8, nil}
	tree31 := tree.Tree{nil, 1, nil}
	tree32 := tree.Tree{nil, 2, nil}
	tree33 := tree.Tree{nil, 5, nil}
	tree34 := tree.Tree{nil, 13, nil}
	tree1.Left = &tree21
	tree1.Right = &tree22
	tree21.Left = &tree31
	tree21.Right = &tree32
	tree22.Left = &tree33
	tree22.Right = &tree34

	ch := make(chan int, 7)
	Walk(&tree1, ch)
	fmt.Println("tree1先序遍历结果为")
	for i := 0; i < 7; i++ {
		value := <-ch
		fmt.Printf("%d\t", value) // This should print 1 to 10
	}
	fmt.Println()
	close(ch)
	tree11 := tree1
	fmt.Println(Same(&tree1, &tree11))

}
