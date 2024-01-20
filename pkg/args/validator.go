package args

import (
	"errors"
	"fmt"
	"log"
)

func ValidateUserFlags() []UserFlag {
	for _, f := range Flags {
		fmt.Println(f.F)
		i, err := getProgramFlag(f.F)

		if err != nil {
			log.Fatal("Invalid flag given")
		}

		err = validateUserFlag(f, ProgramFlags[i])

		if err != nil {
			log.Fatal("Flag validation failed")
		}
	}

	return Flags
}

func validateUserFlag(uFlag UserFlag, flag Flag) error {
	if flag.ParamRequired && uFlag.Parameter == "" {
		return errors.New("User has not provided a parameter for a flag")
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
