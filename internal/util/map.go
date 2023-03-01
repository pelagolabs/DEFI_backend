package util

import "sync"

type SyncedMap[K comparable, V any] struct {
	m sync.Map
}

func (m *SyncedMap[K, V]) Load(key K) (value V, ok bool) {
	load, ok := m.m.Load(key)
	if ok {
		return load.(V), ok
	} else {
		return value, ok
	}
}

func (m *SyncedMap[K, V]) Store(key K, value V) {
	m.m.Store(key, value)
}

func (m *SyncedMap[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	store, loaded := m.m.LoadOrStore(key, value)
	return store.(V), loaded
}

func (m *SyncedMap[K, V]) Delete(key K) {
	m.m.Delete(key)
}

func (m *SyncedMap[K, V]) LoadAndDelete(key K) (value V, loaded bool) {
	store, loaded := m.m.LoadAndDelete(key)
	return store.(V), loaded
}

func (m *SyncedMap[K, V]) Range(f func(key K, value V) bool) {
	m.m.Range(func(key, value any) bool {
		return f(key.(K), value.(V))
	})
}
