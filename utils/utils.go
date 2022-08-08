package utils

func StringInList(s string, l []string) bool {
	for _, i := range l {
		if s == i {
			return true
		}
	}
	return false
}
