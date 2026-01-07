package commonval

import "strconv"

func ParseStringToInt(value string) int {
	parsedValue, err := strconv.Atoi(value)
	if err != nil {
		return 0 // Default to 0 if parsing fails
	}
	return parsedValue
}
