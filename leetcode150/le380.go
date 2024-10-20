package leetcode150

import "math/rand"

type RandomizedSet struct {
	set    map[int]struct{}
	length int
}

func Constructor() RandomizedSet {
	return RandomizedSet{make(map[int]struct{}, 0), 0}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, exists := this.set[val]; exists {
		return false
	}
	this.set[val] = struct{}{}
	this.length++
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, exists := this.set[val]; exists {
		delete(this.set, val)
		this.length--
		return true
	}
	return false
}

// GetRand 不高效，因为是 O(n) , 可以直接把值 存到一个新的切片里面，然后秒杀
func (this *RandomizedSet) GetRandom() int {
	i := rand.Uint64() % uint64(this.length)
	for k, _ := range this.set {
		if i == 0 {
			return k
		}
		i--
	}
	return 0
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
