package args

import (
	"errors"
	"fmt"
)

func ValidateUserFlags(flags *[]UserFlag, inRepl bool) (error, bool) {
	if len(*flags) == 0 && inRepl {
		return nil, true
	}

	for uFlagIdx, f := range *flags {
		err, idx := getProgramFlag(f.F)

		if err != nil {
			return errors.New(fmt.Sprintf("Unrecognised flag: %v", f.F)), false
		}

		applyDefaultValues(&(*flags)[uFlagIdx], ProgramFlags[idx], inRepl)
		err = validateUserFlag((*flags)[uFlagIdx], ProgramFlags[idx], inRepl)

		if err != nil {
			return err, false
		}
	}
	return nil, true
}

func applyDefaultValues(uFlag *UserFlag, f flag, inRepl bool) {
	if uFlag.Parameter == "" {
		fmt.Println(uFlag.Parameter, f.DefaultValue)
		uFlag.Parameter = f.DefaultValue
	}
}

func validateUserFlag(uFlag UserFlag, f flag, inRepl bool) error {
	if f.ParamAccepted && uFlag.Parameter == "" && f.DefaultValue == "" {
		return errors.New(fmt.Sprintf("Parameter required for flag: %v", uFlag.F))
	}

	if !inRepl && f.ReplOnly {
		return errors.New(fmt.Sprintf("Flag '%v' only valid in REPL", uFlag.F))
	}

	return nil
}

func getProgramFlag(s string) (error, int) {
	for i := 0; i < len(ProgramFlags); i++ {
		if ProgramFlags[i].Flag == s || ProgramFlags[i].Short == s {
			return nil, i
		}
	}
	return errors.New("Unrecognised Flag"), -1
}
