package utils

func CheckValue(value string) string {
	if value == "" {
		return "<empty>"
	}
	return value
}