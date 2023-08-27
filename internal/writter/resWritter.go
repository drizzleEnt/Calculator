package writter

import "fmt"

func WriteOut(a int, isArabic bool) {

	//nums := map[int]string{1: "I", 2: "II", 3: "III", 4: "IV", 5: "V", 6: "VI", 7: "VII", 8: "VIII", 9: "IX", 10: "X"}

	if isArabic {
		fmt.Println(a)

	} else {
		WriteRim(a) //fmt.Println(nums[a])
	}
}

func WriteRim(a int) {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	numerals := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	var result string

	for i := 0; i < len(values); i++ {
		for a >= values[i] {
			a -= values[i]
			result += numerals[i]
		}
	}

	fmt.Println(result)
}
