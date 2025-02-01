package main

func QuickSort(A []int) {
	qs(A, 0, len(A)-1)
}

func qs(A []int, l, r int) {
	if l >= r {
		return
	}

	pi := pivot(A, l, r)

	qs(A, l, pi-1)
	qs(A, pi+1, r)
}

func pivot(A []int, l, r int) int {
	pivot := A[r]

	i := l - 1

	for j := l; j < r; j++ {
		if A[j] >= pivot {
			continue
		}

		i++
		tmp := A[j]
		A[j] = A[i]
		A[i] = tmp
		triggerRefresh()
	}

	i++
	A[r] = A[i]
	A[i] = pivot

	// triggerRefresh()

	return i
}
