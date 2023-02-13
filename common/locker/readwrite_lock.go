package locker

type RWLocker interface {
	Lock(version string)
	UnLock(version string)
	RLock(version string)
	RUnLock(version string)
}
