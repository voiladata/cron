package cron

type EntryHeap []*Entry

func (h *EntryHeap) Less(i, j int) bool {
	if (*h)[i].Next.IsZero() {
		return false
	}
	if (*h)[j].Next.IsZero() {
		return true
	}
	return (*h)[i].Next.Before((*h)[j].Next)
}

func (h *EntryHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *EntryHeap) Len() int {
	return len(*h)
}

func (h *EntryHeap) Pop() (v interface{}) {
	if len(*h) == 0 {
		return
	}
	*h, v = (*h)[:len(*h)-1], (*h)[len(*h)-1]
	return
}

func (h *EntryHeap) Push(v interface{}) {
	*h = append(*h, v.(*Entry))
}

func (h *EntryHeap) Peek() *Entry {
	if len(*h) == 0 {
		return nil
	}
	return (*h)[0]
}
