package main

// https://tylercipriani.com/blog/2020/01/12/unix-permissions-for-dummies/

import (
	"fmt"
	"strconv"
)

// I think you can write this function more elegant by utilizig bit operations

func parseUnixPermissions(octal string) (string, error) {
	var special int
	if len(octal) > 4 || len(octal) < 3 {
		return "", fmt.Errorf("invalid format")
	}
	if len(octal) == 4 {
		special, _ = strconv.Atoi(string(octal[0]))
		octal = octal[1:]
	}

	other, err := strconv.ParseInt(string(octal[2]), 10, 0)
	if err != nil {
		return "", err
	}
	group, err := strconv.ParseInt(string(octal[1]), 10, 0)
	if err != nil {
		return "", err
	}
	user, err := strconv.ParseInt(string(octal[0]), 10, 0)
	if err != nil {
		return "", err
	}

	if user > 7 || user < 0 {
		return "", fmt.Errorf("invalid format")
	}

	if group > 7 || group < 0 {
		return "", fmt.Errorf("invalid format")
	}

	if other > 7 || other < 0 {
		return "", fmt.Errorf("invalid format")
	}

	octalToStr := map[int64]string{
		0: "---",
		1: "--x",
		2: "-w-",
		3: "-wx",
		4: "r--",
		5: "r-x",
		6: "rw-",
		7: "rwx",
	}

	u := octalToStr[user]
	g := octalToStr[group]
	o := octalToStr[other]

	// sticky
	if special == 1 {
		if o[2] == 'x' {
			o = o[:2] + "t"
		} else {
			o = o[:2] + "T"
		}
	}
	// setgid
	if special == 2 {
		if g[2] == 'x' {
			g = g[:2] + "s"
		} else {
			g = g[:2] + "S"
		}
	}
	// sticky+setgid
	if special == 3 {
		if g[2] == 'x' {
			g = g[:2] + "s"
		} else {
			g = g[:2] + "S"
		}
		if o[2] == 'x' {
			o = o[:2] + "t"
		} else {
			o = o[:2] + "T"
		}
	}
	// setuid
	if special == 4 {
		if u[2] == 'x' {
			u = u[:2] + "s"
		} else {
			u = u[:2] + "S"
		}
	}
	// setuid+sticky
	if special == 5 {
		if u[2] == 'x' {
			u = u[:2] + "s"
		} else {
			u = u[:2] + "S"
		}
		if o[2] == 'x' {
			o = o[:2] + "t"
		} else {
			o = o[:2] + "T"
		}
	}
	// setuid+setgid
	if special == 6 {
		if u[2] == 'x' {
			u = u[:2] + "s"
		} else {
			u = u[:2] + "S"
		}
		if g[2] == 'x' {
			g = g[:2] + "s"
		} else {
			g = g[:2] + "S"
		}
	}
	// sticky+setuid+setgid
	if special == 7 {
		if u[2] == 'x' {
			u = u[:2] + "s"
		} else {
			u = u[:2] + "S"
		}
		if g[2] == 'x' {
			g = g[:2] + "s"
		} else {
			g = g[:2] + "S"
		}
		if o[2] == 'x' {
			o = o[:2] + "t"
		} else {
			o = o[:2] + "T"
		}
	}

	return u + g + o, nil
}
