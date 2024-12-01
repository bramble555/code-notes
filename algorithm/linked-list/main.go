package main

import (
	"fmt"
	"gonotes/algorithm/linked-list/le206"
	"gonotes/algorithm/linked-list/le24"
	"gonotes/algorithm/linked-list/le707"
)

func main() {
	// 测试le707
	obj := le707.Constructor()
	obj.AddAtHead(7)
	fmt.Println(obj.Size)
	obj.AddAtHead(2)
	fmt.Println(obj.Size)
	obj.AddAtHead(1)
	fmt.Println(obj.Size)
	obj.AddAtIndex(3, 0)
	fmt.Println("o val", obj.Get(3))
	fmt.Println(obj.Size)
	obj.DeleteAtIndex(2)
	fmt.Println(obj.Size)
	obj.AddAtHead(6)
	fmt.Println(obj.Size)
	obj.AddAtTail(4)
	fmt.Println(obj.Size)
	fmt.Println(obj.Get(4))
	fmt.Println()
	// 测试le206
	a := &le206.ListNode{0, nil}
	b := &le206.ListNode{1, nil}
	c := &le206.ListNode{2, nil}
	d := &le206.ListNode{3, nil}
	a.Next = b
	b.Next = c
	c.Next = d
	// fmt.Printf("%v\n", le206.ReverseList(a).Val)
	// fmt.Println(le206.ReverseList1(d).Val)
	f := &le24.ListNode{0, nil}
	g := &le24.ListNode{1, nil}
	h := &le24.ListNode{2, nil}
	// i := &le24.ListNode{3, nil}
	f.Next = g
	g.Next = h
	// h.Next = i
	// fmt.Println(le24.SwapPairs(f).Val)
	// dakkda := &le24.ListNode{0, nil}
	// var dkadka *le24.ListNode = nil
	fmt.Println(le24.SwapPairs(f))
	fmt.Println(le24.SwapPairs(f).Next)
	fmt.Println(le24.SwapPairs(f).Next.Next)
	// fmt.Println(le24.SwapPairs(f).Next.Next.Next)

}
