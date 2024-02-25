package xtype

import "find-bin-width/pkg/xutil"

const (
	yearDivider   = 4
	monthDivider  = 7
	dayDivider    = 10
	hourDivider   = 13
	minuteDivider = 16
	secondDivider = 19
)

func valueIsDate(value string) bool {
	l := len(value)

	switch l {
	case 10:
		return validateShortDate(value)
	case 19, 20, 25:
		result := validateShortDateTime(value)
		if result == vsdtResultOK {
			return true
		} else if result == vsdtResultInvalid || value[secondDivider] != '.' || l < secondDivider+2 {
			return false
		}
	}

	return validateTimestamp(value)
}

func validateShortDate(value string) bool {
	if value[yearDivider] != '-' || value[monthDivider] != '-' {
		return false
	}

	year, ok := dateYearIsValid(value)

	if !ok {
		return false
	}

	month, ok := dateMonthIsValid(value)

	if !ok {
		return false
	}

	return dateDayIsValid(value, year, month)
}

const (
	vsdtResultOK uint8 = iota
	vsdtResultMaybeISO
	vsdtResultInvalid
)

func validateShortDateTime(value string) uint8 {
	ok := (value[dayDivider] == ' ' || value[dayDivider] == 'T') &&
		validateShortDate(value) &&
		dateHourIsValid(value, dayDivider) &&
		dateMinSecIsValid(value, hourDivider) &&
		dateMinSecIsValid(value, minuteDivider)

	if !ok {
		return vsdtResultInvalid
	}

	l := len(value)

	if l == 19 || (l == 20 && value[secondDivider] == 'Z') {
		return vsdtResultOK
	}

	if l != 25 || (value[secondDivider] != '-' && value[secondDivider] != '+' || value[22] != ':') {
		return vsdtResultMaybeISO
	}

	if dateHourIsValid(value, secondDivider) && dateMinSecIsValid(value, 22) {
		return vsdtResultOK
	}

	return vsdtResultInvalid
}

func validateTimestamp(value string) bool {
	// If we're here then we know that the length of the value is > 20 and the
	// 19th character is a '.'
	l := len(value)

	if l > 36 {
		return false
	}

	i := 20
	for ; i < 36; i++ {
		if !xutil.CharIsNumeric(value[i]) {
			break
		}
	}

	if i == 20 {
		return false
	}

	if i+1 == l && value[i] == 'Z' {
		return true
	}

	if l != i+6 || (value[i] != '-' && value[i] != '+') || value[i+3] != ':' {
		return false
	}

	return dateHourIsValid(value, i) && dateMinSecIsValid(value, i+3)
}

func dateMinSecIsValid(value string, offset int) bool {
	return xutil.CharIsNumberNoMoreThan(value[offset+1], '5') && xutil.CharIsNumeric(value[offset+2])
}

func dateHourIsValid(value string, offset int) bool {
	a := value[offset+1]
	b := value[offset+2]

	if a < '0' || b < '0' {
		return false
	}

	switch a {
	case 0, 1:
		return xutil.CharIsNumeric(b)
	case 2:
		return xutil.CharIsNumberNoMoreThan(b, '4')
	default:
		return false
	}
}

func dateDayIsValid(value string, year uint16, month uint8) bool {
	a := value[8]
	b := value[9]

	if a < '0' || b < '0' {
		return false
	}

	switch month {

	case 1, 3, 5, 7, 8, 10, 12:
		if a > '3' || (a < '3' && b > '9') || (a == '3' && b > '1') {
			return false
		}

	case 4, 6, 9, 11:
		if a > '3' || (a < '3' && b > '9') || (a == '3' && b != '0') {
			return false
		}

	case 2:
		if a > '2' {
			return false
		}

		if a < '2' {
			return b <= '9'
		}

		if year%4 == 0 {
			if b > '9' {
				return false
			}
		} else if b > '8' {
			return false
		}

	default:
		panic("illegal state")
	}

	return true
}

func dateMonthIsValid(value string) (out uint8, ok bool) {
	switch value[5] {
	case '0':
		if !xutil.CharIsNumeric(value[6]) {
			return
		}

		return value[6] - '0', true

	case '1':
		if !xutil.CharIsNumberNoMoreThan(value[6], '2') {
			return
		}

		return 10 + (value[6] - '0'), true

	default:
		return
	}
}

func dateYearIsValid(value string) (out uint16, ok bool) {
	for i := 0; i < yearDivider; i++ {
		if !xutil.CharIsNumeric(value[i]) {
			return
		}
		out = out*10 + uint16(value[i]-'0')
	}

	return out, true
}
