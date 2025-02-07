package main

func InsertionSort(A []int) {
	for i := 1; i < len(A); i++ {
		n := A[i]
		j := i - 1

		for j >= 0 {
			if A[j] < n {
				break
			}

			A[j + 1] = A[j]
			j--
			triggerRefresh()
		}

		A[j + 1] = n
	}
}
