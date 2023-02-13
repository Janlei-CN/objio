package heartbeat

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestChooseRandomDataServers(t *testing.T) {

	perm := rand.Perm(5)
	fmt.Println(perm)
}
