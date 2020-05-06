package singleton

import (
	"sync"
)

var (
	once sync.Once
	p    *pet
)

func init() {
	once.Do(
		func() {
			p = &pet{}
		})
}

func GetInstance() *pet {
	return p
}

type pet struct {
	name string
	age  int
	mu   sync.Mutex
}

func (p *pet) SetName(name string) {
	p.mu.Lock()
	p.mu.Unlock()
	p.name = name
}

func (p *pet) Name() string {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.name
}

func (p *pet) IncAge() {
	p.mu.Lock()
	p.mu.Unlock()
	p.age++
}

func (p *pet) Age() int {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.age
}
