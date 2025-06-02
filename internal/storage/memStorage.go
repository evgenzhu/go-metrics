package storage

type MemoryStorage struct {
	Gauge   map[string]float64
	Counter map[string]int64
}

func (m *MemoryStorage) SetGauge(name string, value float64) {
	m.Gauge[name] = value
}

func (m *MemoryStorage) AddCounter(name string, delta int64) {
	m.Counter[name] += delta
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		Gauge:   make(map[string]float64),
		Counter: make(map[string]int64),
	}
}
