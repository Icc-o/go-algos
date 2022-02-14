package bitmap

import "fmt"

type BitMap struct {
	data []uint32
	cap  uint64
}

func NewBitMap(c uint64) (*BitMap, error) {
	var bm = BitMap{
		data: make([]uint32, c, c),
		cap:  c,
	}
	return &bm, nil
}

func (bm *BitMap) Set(val uint64) {
	index := val / 32
	// index := val >> 5
	offset := val % 32
	bm.data[index] |= 1 << offset
}

func (bm *BitMap) Exist(val uint64) bool {
	index := val / 32
	// index := val >> 5
	offset := val % 32

	return (bm.data[index] & 1 << offset) != 0
}

func (bm *BitMap) Print() {
	for i := 0; i < 10; i++ {
		fmt.Println(bm.data[i])
	}
}
