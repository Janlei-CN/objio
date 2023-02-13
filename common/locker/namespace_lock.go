package locker

import (
	"log"
	"sync"
)

var NsMap *dsyncRwLockMap

func init() {
	NsMap = NewDsyncRwLockMap()
}

type nsLock struct {
	ref int32
	*sync.RWMutex
}

type dsyncRwLockMap struct {
	lockMap map[string]*nsLock
	mutex   sync.Mutex
}

func NewDsyncRwLockMap() *dsyncRwLockMap {
	m := dsyncRwLockMap{}

	m.lockMap = make(map[string]*nsLock)

	return &m
}

func (d *dsyncRwLockMap) Lock(version string) {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	_, found := d.lockMap[version]

	if !found {
		d.lockMap[version] = &nsLock{
			ref: 0,
		}
	}

	d.lockMap[version].Lock()
	d.lockMap[version].ref++
}

func (d *dsyncRwLockMap) UnLock(version string) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	_, found := d.lockMap[version]

	if !found {
		log.Println("not exist lock")
	}

	d.lockMap[version].Unlock()

	d.lockMap[version].ref--
	if d.lockMap[version].ref == 0 {
		// Remove from the map if there are no more references.
		delete(d.lockMap, version)
	}
}

func (d *dsyncRwLockMap) RLock(version string) {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	_, found := d.lockMap[version]

	if !found {
		d.lockMap[version] = &nsLock{
			ref: 0,
		}
	}

	d.lockMap[version].RLock()
	d.lockMap[version].ref++
}

func (d *dsyncRwLockMap) RUnLock(version string) {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	_, found := d.lockMap[version]

	if !found {
		log.Println("not exist lock")
	}

	d.lockMap[version].RUnlock()

	d.lockMap[version].ref--
	if d.lockMap[version].ref == 0 {
		// Remove from the map if there are no more references.
		delete(d.lockMap, version)
	}
}
