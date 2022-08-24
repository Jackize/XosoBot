package Find

var (
	day, month, year int
)

func numberOfDay() int {
	if month == 4 || month == 6 || month == 9 || month == 11 {
		return 30
	} else {
		if month == 2 {
			if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
				return 29
			} else {
				return 28
			}
		}
	}
	return 31
}

func checkMonth(month int) bool {
	if month > 0 && month < 13 {
		return true
	}
	return false
}
