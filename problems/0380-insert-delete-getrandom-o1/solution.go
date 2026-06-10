// https://leetcode.com/problems/insert-delete-getrandom-o1/description/
package insertdeletegetrandomo1

import "math/rand"

type RandomizedSet struct {
	index map[int]int
	array []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		index: make(map[int]int),
		array: make([]int, 0),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.index[val]; ok {
		return false
	}
	this.array = append(this.array, val)
	this.index[val] = len(this.array) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	idx, ok := this.index[val]
	if !ok {
		return false
	}

	if idx != len(this.array)-1 {
		tmp := this.array[len(this.array)-1]
		this.array[idx] = tmp
		this.index[tmp] = idx
	}

	this.array = this.array[:len(this.array)-1]
	delete(this.index, val)

	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.array[rand.Intn(len(this.array))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
