package roman_numerals

import (
	"testing"
)

func TestConvertOne(t *testing.T) {
	one := 1
	expected := "I"

	result := ToRoman(one)

	assertThatNumeralIsCorrect(one, result, expected, t)
}

func TestConvert1000(t *testing.T) {
	thousand := 1000
	expected := "M"

	result := ToRoman(thousand)

	assertThatNumeralIsCorrect(thousand, result, expected, t)
}

func TestConvert2000(t *testing.T) {
	value := 2000
	expected := "MM"

	result := ToRoman(value)

	assertThatNumeralIsCorrect(value, result, expected, t)
}

func TestConvert500(t *testing.T) {
	value := 500
	expected := "D"

	result := ToRoman(value)

	assertThatNumeralIsCorrect(value, result, expected, t)
}

func TestConvert100(t *testing.T) {
	value := 100
	expected := "C"

	result := ToRoman(value)

	assertThatNumeralIsCorrect(value, result, expected, t)
}

func TestConvert200(t *testing.T) {
	value := 200
	expected := "CC"

	result := ToRoman(value)

	assertThatNumeralIsCorrect(value, result, expected, t)
}

func TestConvert50(t *testing.T) {
	value := 50
	expected := "L"

	result := ToRoman(value)

	assertThatNumeralIsCorrect(value, result, expected, t)
}

func TestConvert10(t *testing.T) {
	value := 10
	expected := "X"

	result := ToRoman(value)

	assertThatNumeralIsCorrect(value, result, expected, t)
}

func TestConvert5(t *testing.T) {
	value := 5
	expected := "V"

	result := ToRoman(value)

	assertThatNumeralIsCorrect(value, result, expected, t)
}

func TestConvert4(t *testing.T) {
	value := 4
	expected := "IV"

	result := ToRoman(value)

	assertThatNumeralIsCorrect(value, result, expected, t)
}

func TestConvert9(t *testing.T) {
	value := 9
	expected := "IX"

	result := ToRoman(value)

	assertThatNumeralIsCorrect(value, result, expected, t)
}

func TestConvert900(t *testing.T) {
	value := 900
	expected := "CM"

	result := ToRoman(value)

	assertThatNumeralIsCorrect(value, result, expected, t)
}

func TestConvert400(t *testing.T) {
	value := 400
	expected := "CD"

	result := ToRoman(value)

	assertThatNumeralIsCorrect(value, result, expected, t)
}

var cases = []struct {
	value    int
	expected string
}{
	{12, "XII"},
	{76, "LXXVI"},
	{1776, "MDCCLXXVI"},
	{94, "XCIV"},
	{4, "IV"},
	{34, "XXXIV"},
	{267, "CCLXVII"},
	{764, "DCCLXIV"},
	{987, "CMLXXXVII"},
	{1983, "MCMLXXXIII"},
	{2014, "MMXIV"},
	{4000, "MMMM"},
	{4999, "MMMMCMXCIX"},
}

func TestManyCases(t *testing.T) {
	for _, test := range cases {
		t.Run(test.expected, func(t *testing.T) {
			result := ToRoman(test.value)
			assertThatNumeralIsCorrect(test.value, result, test.expected, t)
		})

	}
}

func assertThatNumeralIsCorrect(number int, numeral string, expectedNumeral string, t *testing.T) {
	if numeral != expectedNumeral {
		t.Errorf("roman numeral for %d is %s , not %s", number, expectedNumeral, numeral)
	}
}
