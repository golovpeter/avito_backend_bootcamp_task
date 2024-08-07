package house_create

func validateInParams(in CreateHouseIn) (bool, string, error) {
	if in.Year <= 0 || in.Developer == "" || in.Address == "" {
		return false, "invalid input", nil
	}

	return true, "", nil
}
