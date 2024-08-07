package update_flat_status

func validateInParams(in UpdateFlatStatusIn) (bool, string, error) {
	if in.ID <= 0 || in.Status == "" {
		return false, "invalid input", nil
	}

	return true, "", nil
}
