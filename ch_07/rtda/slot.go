package rtda

import "ch_07/rtda/heap"

type Slot struct {
	num int32
	ref *heap.Object
}
