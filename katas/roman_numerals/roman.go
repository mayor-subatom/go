package roman_numerals

//TODO: refactor
func ToRoman(value int) string {

	switch {

	case value == 0:
		return ""

	case value/1000 > 0:
		return "M" + ToRoman(value-1000)

	case value/900 == 1:
		return "CM" + ToRoman(value-900)

	case value/500 == 1:
		return "D" + ToRoman(value-500)

	case value/400 == 1:
		return "CD" + ToRoman(value-400)

	case value/100 > 0:
		return "C" + ToRoman(value-100)

	case value/90 == 1:
		return "XC" + ToRoman(value-90)

	case value/50 == 1:
		return "L" + ToRoman(value-50)

	case value/40 == 1:
		return "XL" + ToRoman(value-40)

	case value/10 > 0:
		return "X" + ToRoman(value-10)

	case value == 9:
		return "IX" + ToRoman(value-9)

	case value/5 == 1:
		return "V" + ToRoman(value-5)

	case value == 4:
		return "IV" + ToRoman(value-4)

	default:
		return "I" + ToRoman(value-1)

	}
}
