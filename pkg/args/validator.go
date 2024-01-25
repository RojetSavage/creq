package args

import (
	"errors"
)

func ValidateUserFlags() []UserFlag {
	for _, f := range Flags {
		i, err := getProgramFlag(f.F)

		if err != nil {
			panic("Bad flag")
		}

		err = validateUserFlag(f, ProgramFlags[i])

		if err != nil {
			panic("User has not provided a required parameter")
		}
	}

	return Flags
}

func validateUserFlag(uFlag UserFlag, f flag) error {
	if f.ParamRequired && uFlag.Parameter == "" {
		return errors.New("User has not provided a required parameter")
	}
	return nil
}

func getProgramFlag(s string) (int, error) {
	for i := 0; i < len(ProgramFlags); i++ {
		if ProgramFlags[i].Flag == s || ProgramFlags[i].Short == s {
			return i, nil
		}
	}
	return -1, errors.New("Invalid Flag Given")
}
