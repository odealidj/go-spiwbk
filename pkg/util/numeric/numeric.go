package numeric

import "regexp"

func IsNumeric(s string) bool {
	re := regexp.MustCompile(`^[0-9]+(\.[0-9]+)?$`)
	return re.Match([]byte(s))
}
