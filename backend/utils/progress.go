package utils

func CalculateProgress(theory, examples bool, testScore int) int {
	progress := 0

	if theory {
		progress += 20
	}
	if examples {
		progress += 30
	}

	if theory && examples {
		progress += int(float64(testScore) / 20 * 50)
	} else if testScore > 0 {
		progress += 50
	}

	if progress > 100 {
		progress = 100
	}

	return progress
}
