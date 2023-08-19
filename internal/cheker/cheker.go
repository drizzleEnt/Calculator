package cheker

import "strconv"

func CheckIsArabic(numstr string) (int, bool) {
	num, err := strconv.Atoi(numstr)

	if err != nil {
		return 0, false
	}
	return num, true
}

func CheckIsRim(numstr string) (int, bool) {
	rim := map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10}
	if num, ok := rim[numstr]; ok {
		return num, ok
	}
	return 0, false
}
