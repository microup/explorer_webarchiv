package utils

func IsValueExists(target string, list []string) bool {
	for _, value := range list {
		if value == target {
			return true
		}
	}
	return false
}