package main

func BubbleSort(A []int) {
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A)-1-i; j++ {
			if A[j] > A[j+1] {
				tmp := A[j]
				A[j] = A[j+1]
				A[j+1] = tmp
			}
		}
		triggerRefresh()
	}
}
