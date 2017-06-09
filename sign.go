package birthsign

func GetLabel(code int, gender string) string {

	genderAddition := `o`

	if gender == `F` {
		genderAddition = `a`
	}

	switch code {
	case 1:
		return "Arian" + genderAddition
	case 2:
		return "Taurin" + genderAddition
	case 3:
		return "Geminian" + genderAddition
	case 4:
		return "Cancerian" + genderAddition
	case 5:
		return "Leonin" + genderAddition
	case 6:
		return "Virginian" + genderAddition
	case 7:
		return "Librian" + genderAddition
	case 8:
		return "Escorpian" + genderAddition
	case 9:
		return "Sagitarian" + genderAddition
	case 10:
		return "Capricornian" + genderAddition
	case 11:
		return "Aquarian" + genderAddition
	case 12:
		return "Piscian" + genderAddition
	}

	return ``
}
