package singleton

import (
	"sync"
	"testing"
)

func IncAge1() {
	p := GetInstance()
	p.IncAge()
}

func IncAge2() {
	p := GetInstance()
	p.IncAge()
}

func TestSingleton(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(200)

	for i := 0; i < 100; i++ {
		go func() {
			wg.Done()
			IncAge1()
		}()

		go func() {
			wg.Done()
			IncAge2()
		}()
	}

	wg.Wait()

	p := GetInstance()
	if p.Age() != 200 {
		t.Fatalf("期待age为200，实际为%d", p.Age())
	}
	t.Log("测试成功!!")
}
