package priorityqueue

type BinaryIntHeap struct {
	elements [][2]int
}

func (a *BinaryIntHeap) push(p [2]int) {
	up := func() {
		//push元素的下标
		pIndex := a.Len() - 1
		for {
			parentIndex := (pIndex - 1) / 2
			if pIndex == parentIndex || a.Less(pIndex, parentIndex) {
				break
			} else {
				a.Swap(pIndex, parentIndex)
				pIndex = parentIndex
			}
		}
	}
	a.elements = append(a.elements, p)
	up()
}

func (a *BinaryIntHeap) pop() [2]int {
	down := func() {
		idx := 0
		for {
			cLeft := 2*idx + 1
			if cLeft < 0 || cLeft >= a.Len() {
				break
			}
			cswap := cLeft
			cRight := 2*idx + 2
			if cRight < a.Len() && a.Less(cLeft, cRight) {
				cswap = cRight
			}
			if a.Less(cswap, idx) {
				break
			}
			a.Swap(idx, cswap)
			idx = cswap
		}
	}
	a.Swap(0, a.Len()-1)
	var val [2]int
	a.elements, val = a.elements[:a.Len()-1], a.elements[a.Len()-1]
	down()
	return val
}

func (a *BinaryIntHeap) Len() int {
	return len(a.elements)
}

func (a *BinaryIntHeap) Less(i, j int) bool {
	return a.elements[i][1] < a.elements[j][1]
}

// Swap swaps the elements with indexes i and j.
func (a *BinaryIntHeap) Swap(i, j int) {
	a.elements[i], a.elements[j] = a.elements[j], a.elements[i]
}
