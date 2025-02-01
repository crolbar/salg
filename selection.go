package main

func SelectionSort(A []int) {
	for i := 0; i < len(A); i++ {
		l := i
		for j := i + 1; j < len(A); j++ {
			if (A[j] < A[l]) {
				l = j
			}
		}

		if l == i {
			continue
		}

		tmp := A[i]
		A[i] = A[l]
		A[l] = tmp
		triggerRefresh()
	}
}
