package get_flats

func validateInParams(houseID int) (bool, string, error) {
	if houseID <= 0 {
		return false, "invalid input", nil
	}

	return true, "", nil
}
