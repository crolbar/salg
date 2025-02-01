package main

func MergeSort(A []int) {
	mergeSort(A, 0, len(A)-1)
}

func mergeSort(A []int, p, r int) {
	if p >= r {
		return
	}

	var q int = p + (r-p)/2

	mergeSort(A, p, q)
	mergeSort(A, q+1, r)

	rMerge(A, p, q, r)
}

func rMerge(A []int, p, q, r int) {
	nl := q - p + 1
	nr := r - q

	L := make([]int, nl)
	R := make([]int, nr)

	for i := 0; i < nl; i++ {
		L[i] = A[p+i]
	}

	for i := 0; i < nr; i++ {
		R[i] = A[q+1+i]
	}

	var (
		i = 0
		j = 0
		k = p
	)

	for i < nl && j < nr {
		if L[i] < R[j] {
			A[k] = L[i]
			triggerRefresh()
			i++
		} else {
			A[k] = R[j]
			triggerRefresh()
			j++
		}

		k++
	}

	for i < nl {
		A[k] = L[i]
		triggerRefresh()
		k++
		i++
	}

	for j < nr {
		A[k] = R[j]
		triggerRefresh()
		k++
		j++
	}
}

