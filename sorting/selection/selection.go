package selection

func findSmallest(arr []int) int {
	smallest := arr[0]
	smallestIndex := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] < smallest {
			smallest = arr[i]
			smallestIndex = i
		}
	}
	return smallestIndex
}

func Sort(arr []int) []int {
	maxArr := len(arr)
	var newArr []int

	for i := 0; i < maxArr; i++ {
		smallest := findSmallest(arr)
		newArr = append(newArr, arr[smallest])
		arr = append(arr[:smallest], arr[smallest+1:]...)
	}

	return newArr
}
