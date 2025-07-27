package main

func ConvertToRomain(num int) string {
	switch num {
	case 5:
		return "V"
	case 4:
		return "IV"
	case 3:
		return "III"
	case 2:
		return "II"
	case 1:
		return "I"
	}

	return ""
}
