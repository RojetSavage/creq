package args

import (
	"errors"
	"fmt"
)

func ValidateUserFlags(flags []UserFlag, inRepl bool) (error, bool) {
	if len(flags) == 0 && inRepl {
		return nil, true
	}

	for i, f := range flags {
		ok := getProgramFlag(f.F)

		if !ok {
			return errors.New(fmt.Sprintf("Unrecognised flag: %v", f.F)), false
		}

		err := validateUserFlag(f, ProgramFlags[i], inRepl)

		if err != nil {
			return errors.New(fmt.Sprintf("Parameter not provided for flag: %v", f.F)), false
		}
	}
	return nil, true
}

func validateUserFlag(uFlag UserFlag, f flag, inRepl bool) error {
	if f.ParamRequired && uFlag.Parameter == "" {
		return errors.New(fmt.Sprintf("Parameter not provided for flag: %v", uFlag.F))
	}

	if !inRepl && f.ReplOnly {
		return errors.New(fmt.Sprintf("Flag '%v' only valid in REPL", uFlag.F))
	}

	return nil
}

func getProgramFlag(s string) bool {
	for i := 0; i < len(ProgramFlags); i++ {
		if ProgramFlags[i].Flag == s || ProgramFlags[i].Short == s {
			return true
		}
	}
	return false
}
