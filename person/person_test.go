package person

import (
	"fmt"
	"testing"
	"time"
)

func TestPerson_Update(t *testing.T) {
	ticker := time.NewTicker(10 * time.Millisecond)
	for _ = range ticker.C {
		pp := GetInstance()

		start := time.Now()
		for _, p := range pp.Persons {
			p.Update()
		}
		cost := time.Since(start)
		fmt.Printf("cost=[%s]\n", cost)
	}
}
