package main

// creating a struct for the map
type FullMap struct {
	idx map[string]map[string]int
}

// creating a new map
func newMap() *FullMap {
	return &FullMap{make(map[string]map[string]int)}
}

// update the map
func (m *FullMap) updateMap(i map[string]map[string]int) {
	m.idx = i
}

// get the map
func (m *FullMap) getMap() map[string]map[string]int {
	return m.idx
}
