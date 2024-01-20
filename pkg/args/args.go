package args

import (
	"fmt"
	"net/url"
	"req"
	"strings"
)

type UserFlag struct {
	F         string
	Parameter string
}

type Flag struct {
	Flag          string
	Short         string
	ParamRequired bool
	DefaultValue  any
}

var Flags []UserFlag
var ProgramFlags []Flag

func init() {
	setProgramFlags()
}

func ParseArgs(cmdLineArgs []string) {
	for i := 1; i < len(cmdLineArgs); i++ {
		var f UserFlag

		//if first arg is url, set url
		if u, err := url.Parse(cmdLineArgs[i]); err == nil && i == 1 {
			req.R.SetUrl(u)
		}

		if b, r := isShort(cmdLineArgs[i]); b == true {
			f.F = string(r)
			Flags = append(Flags, f)
		} else if b, flag := isFlag(cmdLineArgs[i]); b == true {
			f.F = flag
			Flags = append(Flags, f)
			continue
		} else if len(Flags) >= 1 {
			Flags[len(Flags)-1].Parameter = strings.Trim(cmdLineArgs[i], " ")
		}

	}
	fmt.Println(Flags)
}

func isFlag(s string) (bool, string) {
	if len(s) > 2 && strings.HasPrefix(s, "--") {
		return true, strings.Trim(strings.ToLower(s[2:]), " ")
	} else {
		return false, ""
	}
}

func isShort(s string) (bool, string) {
	if len(s) == 2 && strings.HasPrefix(s, "-") {
		return true, string(s[1])
	} else {
		return false, ""
	}
}
