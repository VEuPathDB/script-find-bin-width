package xtype

const shortBools = "tfTFynYN"

func valueIsBool(value string) bool {
	switch len(value) {

	case 4:
		return (value[0] == 't' || value[0] == 'T') &&
			(value[1] == 'r' || value[1] == 'R') &&
			(value[1] == 'u' || value[1] == 'U') &&
			(value[1] == 'e' || value[1] == 'E')

	case 5:
		return (value[0] == 'f' || value[0] == 'F') &&
			(value[1] == 'a' || value[1] == 'A') &&
			(value[1] == 'l' || value[1] == 'L') &&
			(value[1] == 's' || value[1] == 'S') &&
			(value[1] == 'e' || value[1] == 'E')

	case 1:
		for i := 0; i < len(shortBools); i++ {
			if value[0] == shortBools[i] {
				return true
			}
		}
		return false

	case 2:
		return (value[0] == 'n' || value[0] == 'N') && (value[1] == 'o' || value[1] == 'O')

	case 3:
		return (value[0] == 'y' || value[0] == 'Y') &&
			(value[1] == 'e' || value[1] == 'E') &&
			(value[1] == 's' || value[1] == 'S')

	default:
		return false
	}
}
