package writter

import "fmt"

func WriteOut(a int, isArabic bool) {

	nums := map[int]string{1: "I", 2: "II", 3: "III", 4: "IV", 5: "V", 6: "VI", 7: "VII", 8: "VIII", 9: "IX", 10: "X"}

	if isArabic {
		fmt.Println(a)

	} else {
		fmt.Println(nums[a])
	}
}
