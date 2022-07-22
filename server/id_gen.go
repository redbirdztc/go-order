package main

import "sync"

type idType int

const (
	idTypeMessage idType = iota
	idTypeTopic
	idTypeQueue
	idTypeSubscriber
)

// idManager maintain a map for every idType guarantee id uniqueness
type idManager struct {
	strIDMap map[idType]int
	idMutex  *sync.Mutex
}

func newIDManager() *idManager {
	return &idManager{
		make(map[idType]int),
		new(sync.Mutex),
	}
}

func (m *idManager) getID(str idType) int {
	m.idMutex.Lock()
	defer m.idMutex.Unlock()

	v, ok := m.strIDMap[str]

	if ok {
		m.strIDMap[str] = v + 1
		return v
	}

	m.strIDMap[str] = 2
	return 1
}
