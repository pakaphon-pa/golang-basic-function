package romanNumber

var roman = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

var numbers = []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
var romans = []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}

func RomanToInt(s string) int {
	if s == "" {
		return 0
	}

	num, lastint, total := 0, 0, 0

	for i := 0; i < len(s); i++ {
		char := s[len(s)-(i+1) : len(s)-i]
		num = roman[char]

		if num < lastint {
			total = total - num
		} else {
			total = total + num
		}

		lastint = num
	}

	return total
}

func IntToRoman(num int) string {
	var result string = ""
	var index = len(romans) - 1

	for num > 0 {
		for numbers[index] <= num {
			result += romans[index]
			num -= numbers[index]
		}
		index -= 1
	}

	return result
}
