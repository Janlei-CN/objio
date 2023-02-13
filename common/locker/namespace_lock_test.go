package locker

import (
	"testing"
)

func Test_dsyncRwLockMap_Lock(t *testing.T) {
	lockMap := NewDsyncRwLockMap()

	lockMap.Lock("test")
}
