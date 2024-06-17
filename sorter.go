package g-sorter

func Sort[T int | uint | int32 | int64 | uint32 | uint64 | float32](numbers []T) {
	nLen := len(numbers)
	if nLen < 1000 {
		quickSort(numbers, 0, nLen-1)
	} else {
		copy(numbers, mergeSort(numbers))
	}
}

func mergeSort[T int | uint | int32 | int64 | uint32 | uint64 | float32](numbers []T) []T {
	if len(numbers) <= 1 {
		return numbers
	}

	mid := len(numbers) / 2
	left := mergeSort(numbers[:mid])
	right := mergeSort(numbers[mid:])

	return merge(left, right)
}

func merge[T int | uint | int32 | int64 | uint32 | uint64 | float32](left, right []T) []T {
	result := make([]T, 0, len(left)+len(right))

	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

func quickSort[T int | uint | int32 | int64 | uint32 | uint64 | float32](numbers []T, low, high int) {
	if low < high {
		pivot := partition(numbers, low, high)
		quickSort(numbers, low, pivot-1)
		quickSort(numbers, pivot+1, high)
	}
}

func partition[T int | uint | int32 | int64 | uint32 | uint64 | float32](numbers []T, low, high int) int {
	pivot := low - 1
	for i := low; i < high; i++ {
		if numbers[i] < numbers[high] {
			pivot++
			numbers[i], numbers[pivot] = numbers[pivot], numbers[i]
		}
	}
	pivot++
	numbers[pivot], numbers[high] = numbers[high], numbers[pivot]
	return pivot
}
