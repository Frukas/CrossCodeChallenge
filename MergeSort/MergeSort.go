package mergesort

func MergeSort(originalSlice []float32) {
	sliceLength := len(originalSlice)
	if sliceLength < 2 {
		return
	}

	midIndex := sliceLength / 2
	leftHalf := make([]float32, midIndex)
	rightHalf := make([]float32, (sliceLength - midIndex))

	copy(leftHalf, originalSlice[(sliceLength-midIndex):])
	copy(rightHalf, originalSlice[:sliceLength-midIndex])

	MergeSort(leftHalf)
	MergeSort(rightHalf)

	merge(originalSlice, leftHalf, rightHalf)

}

func merge(inSlice []float32, rightHalf []float32, leftHalf []float32) {
	//sliceLength := len(inSlice)
	rightLength := len(rightHalf)
	leftLength := len(leftHalf)

	i := 0
	j := 0
	k := 0

	for i < leftLength && j < rightLength {
		if leftHalf[i] <= rightHalf[j] {
			inSlice[k] = leftHalf[i]
			i++
		} else {
			inSlice[k] = rightHalf[j]
			j++
		}
		k++
	}

	for i < leftLength {
		inSlice[k] = leftHalf[i]
		i++
		k++
	}

	for j < rightLength {
		inSlice[k] = rightHalf[j]
		j++
		k++
	}
}
