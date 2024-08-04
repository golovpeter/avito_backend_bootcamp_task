package common

import "regexp"

const emailValidPattern = "^\\S+@\\S+\\.\\S+$"

func ValidateUserData(email string, password string) (bool, string, error) {
	if email == "" || password == "" {
		return false, "invalid user data", nil
	}

	regExp, err := regexp.Compile(emailValidPattern)
	if err != nil {
		return false, "invalid regular expression", ErrCompileRegexp
	}

	if !regExp.MatchString(email) {
		return false, "invalid email format", nil
	}

	return true, "", nil
}
