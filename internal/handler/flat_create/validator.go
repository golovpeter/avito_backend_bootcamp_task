package flat_create

func validateInParams(in CreateFlatIn) (bool, string, error) {
	if in.HouseID <= 0 || in.Rooms <= 0 || in.Number <= 0 || in.Price <= 0 {
		return false, "invalid input", nil
	}

	return true, "", nil
}
