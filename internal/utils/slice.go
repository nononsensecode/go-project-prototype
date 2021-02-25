package utils

// ItemExists checks whether there is an item in a particular index in a slice
func ItemExists(slice []string, itemIndex int) bool {
	if len(slice) > itemIndex {
		return true
	}

	return false
}