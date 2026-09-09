package slice

func SlicesNumbers(originalSlice []int) []int {
	valueToInsert := 88
	// valueToInsert := 88
	// valueToInsert := []int{88}
	// insertIndex := 3
	mid := len(originalSlice) / 2
	return append(append(originalSlice[:mid], valueToInsert), originalSlice[mid:]...)
	// 	return append(append(originalSlice[:insertIndex], valueToInsert), originalSlice[insertIndex:]...)
}
