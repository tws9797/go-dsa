package essentials

import "math/rand"

type RandomizedSet struct {
	data []int
	dict map[int]int
}

func ConstructorRandomizedSet() RandomizedSet {
	return RandomizedSet{
		data: []int{},
		dict: make(map[int]int),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.dict[val]; ok {
		return false
	}

	this.dict[val] = len(this.data)
	this.data = append(this.data, val)

	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.dict[val]; !ok {
		return false
	}

	l := len(this.data)
	pos := this.dict[val]
	num := this.data[l-1]

	this.dict[num] = pos
	this.data[pos] = num
	delete(this.dict, val)
	this.data = this.data[:l-1]
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.data[rand.Intn(len(this.data))]
}
