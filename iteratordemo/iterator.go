package iterator

/**
* Iterator接口
 */
type IIterator interface {
	HasNext() bool
	Next()
	Index() int
	Value() interface{}
}

type Iterator struct {
	array []interface{}
	idx   *int
}

func (it *Iterator) HasNext() bool {
	return *(it.idx)+1 <= len(it.array)
}

func (it *Iterator) Next() {
	*(it.idx)++
}

func (it *Iterator) Index() int {
	return *(it.idx)
}

func (it *Iterator) Value() interface{} {
	return it.array[*(it.idx)]
}
