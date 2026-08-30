package loan

type Rules struct {
	MinAge    int
	MinIncome int
}

func Check(age, income int, hasBadCredit bool, rules Rules) bool {
	if age < rules.MinAge {
		return false
	}
	if income < rules.MinIncome {
		return false
	}
	if hasBadCredit {
		return false
	} else {
		return true
	}
}
