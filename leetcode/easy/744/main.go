package main

func nextGreatestLetter(letters []byte, target byte) byte {
	left, right := 0, len(letters)-1

	if target >= letters[right] {
		return letters[left]
	}

	for left <= right {
		middle := (left + right) / 2
		middleByte := letters[middle]

		if middleByte > target {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}

	return letters[left]
}

func main() {

}
