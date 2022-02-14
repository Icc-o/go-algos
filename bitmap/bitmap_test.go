package bitmap_test

import (
	"algos/bitmap"
	"testing"
)

func TestSet(t *testing.T) {
	bm, _ := bitmap.NewBitMap(100000)

	for i := 0; i < 5000; i++ {
		if i%32 == 5 {
			bm.Set(uint64(i))
		}
	}
	bm.Print()
}
