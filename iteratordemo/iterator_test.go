package iterator

import (
	"testing"
)

func TestIterator(t *testing.T) {
	arr := []interface{}{1, 35, 7, 9, 10}
	a := 0
	iterator := Iterator{
		array: arr,
		idx:   &a,
	}

	for it := iterator; it.HasNext(); it.Next() {
		idx, value := it.Index(), it.Value().(int)
		if value != arr[idx] {
			t.Fatal("迭代器迭代错误!!")
		}
		t.Logf("%d\t%d\n", idx, value)
	}
}
