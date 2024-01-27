package args

import (
	"errors"
)

func ValidateUserFlags(flags []UserFlag, inRepl bool) bool {
	for _, f := range flags {
		i, err := getProgramFlag(f.F)

		if err != nil {
			panic("Bad flag")
		}

		err = validateUserFlag(f, ProgramFlags[i], inRepl)

		if err != nil {
			panic("User has not provided a required parameter")
		}
	}
	return true
}

func validateUserFlag(uFlag UserFlag, f flag, inRepl bool) error {
	if f.ParamRequired && uFlag.Parameter == "" {
		return errors.New("User has not provided a required parameter")
	}

	if !inRepl && f.ReplOnly {
		return errors.New("Using REPL command in cmd line mode.")
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
