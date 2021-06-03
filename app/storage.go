package app

type Storage interface {
	AddOne()
	Count() uint
}

type MemoryStorage struct {
	count uint
}

func NewMemoryStorage() Storage {
	return &MemoryStorage{
		count: 0,
	}
}

func (p *MemoryStorage) AddOne() {
	p.count++
}

func (p *MemoryStorage) Count() uint {
	return p.count
}
