package dyarray

type ArrayList struct {
	data []interface{}
	size int
}

const INIT_CAP = 1

func NewArrayListWithCap(cap int) *ArrayList {
	return &ArrayList{
		data: make([]interface{}, cap),
		size: 0,
	}
}

func NewArrayList() *ArrayList {
	return NewArrayListWithCap(INIT_CAP)
}

func (this *ArrayList) Size() int {
	return this.size
}

func (this *ArrayList) IsEmpty() bool {
	return this.size == 0
}

func (this *ArrayList) Add(e interface{}) {
	if this.size == len(this.data) {
		this.resize(2 * this.size)
	}
	this.data[this.size] = e
	this.size++
}

func (this *ArrayList) resize(newCap int) {
	newData := make([]interface{}, newCap)
	for i := 0; i < this.size; i++ {
		newData[i] = this.data[i]
	}
	this.data = newData
}

func (this *ArrayList) CheckIndex(index int) {
	if index < 0 || index >= this.size {
		panic("index out of range")
	}
}

func (this *ArrayList) Get(index int) interface{} {
	this.CheckIndex(index)
	return this.data[index]
}

func (this *ArrayList) GetLast() interface{} {
	return this.Get(this.size - 1)
}

func (this *ArrayList) GetFirst() interface{} {
	return this.Get(0)
}

func (this *ArrayList) AddLast(e interface{}) {
	this.Add(e)
}

func (this *ArrayList) AddByIndex(index int, e interface{}) {
	this.CheckIndex(index)
	if this.size == len(this.data) {
		this.resize(2 * this.size)
	}
	for i := this.size - 1; i >= index; i-- {
		this.data[i+1] = this.data[i]
	}
	this.data[index] = e
	this.size++
}

func (this *ArrayList) AddFirst(e interface{}) {
	this.AddByIndex(0, e)
}

func (this *ArrayList) RemoeByIndex(index int) interface{} {
	this.CheckIndex(index)
	e := this.data[index]
	for i := index + 1; i < this.size; i++ {
		this.data[i-1] = this.data[i]
	}
	return e
}

func (this *ArrayList) RemoveLast() interface{} {
	return this.RemoeByIndex(this.size - 1)
}
func (this *ArrayList) RemoveFirst() interface{} {
	return this.RemoeByIndex(0)
}
func (this *ArrayList) Remove(e interface{}) bool {
	for i := 0; i < this.size; i++ {
		if this.data[i] == e {
			this.RemoeByIndex(i)
			return true
		}
	}
	return false
}
func (this *ArrayList) Contains(e interface{}) bool {
	for i := 0; i < this.size; i++ {
		if this.data[i] == e {
			return true
		}
	}
	return false
}
