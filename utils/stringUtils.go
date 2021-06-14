package utils

// helper funtion to search for a string in string array
func Contains(arry []string, str string) (bool, int) {
	for idx, s := range arry {
		if s == str {
			return true, idx
		}
	}

	return false, -1
}
