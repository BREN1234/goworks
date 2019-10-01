package datastructure

func LinearSearch(list []int, element int) bool {
	for _, value := range list {
		if value == element {
			return true
		}
	}
	return false
}
