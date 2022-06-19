package validators

import (
	"regexp"
	"strconv"
	"strings"

	"hushclan.com/pkg/utils"
)

var Regions = [2]string{
	"EU",
	"NA",
}

var Sexes = [3]string{
	"Female",
	"Male",
	"Mixed",
}

var MaxMembers = 10 // 3sub + 5pla + 1coa + 1man
var MaxSubs = 3
var MaxPlayers = 5
var MaxCoaches = 1

func ValidateHexWithHashtag(s string) (int, bool) {
	if len(s) != 7 {
		return 0, false
	}
	if s[0:1] != "#" {
		return 0, false
	}
	hex := s[1:7]
	n, err := strconv.ParseUint(hex, 16, 64)
	if err != nil {
		return 0, false
	}
	return int(n), true
}

func ValidateTeamName(s string) bool {
	if len(s) > 24 {
		return false
	}
	r, _ := regexp.Compile(`^[a-zA-Z0-9\s]+$`)
	return r.MatchString(s)
}

func ValidateRegion(s string) (string, bool) {
	region := strings.ToUpper(s)
	if utils.ContainsString(Regions[:], region) {
		return region, true
	}
	return "", false
}

func ValidateSex(s string) (string, bool) {
	sex := strings.ToLower(s)
	sex = strings.Title(sex)
	if utils.ContainsString(Sexes[:], sex) {
		return sex, true
	}
	return "", false
}
