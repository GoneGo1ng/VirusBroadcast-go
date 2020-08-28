package util

import (
	"fmt"
	"testing"
)

func TestStdGaussian(t *testing.T) {
	for i := 0; i < 100; i++ {
		g := StdGaussian(4, 14)
		fmt.Println(g)
	}
}
