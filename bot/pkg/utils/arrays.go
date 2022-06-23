package utils

import "strings"

func RemoveArrayString(s []string, r string) []string {
	s2 := make([]string, 0, len(s))
	for _, v := range s {
		if v != r {
			s2 = append(s2, v)
		}
	}
	return s2
}

func ContainsString(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func Paginate(x []string, page int, size int) []string {
	skip := page * size
	if skip > len(x) {
		skip = len(x)
	}

	end := skip + size
	if end > len(x) {
		end = len(x)
	}

	return x[skip:end]
}

func NameToID(s string) string {
	o := strings.ToLower(s)
	o = strings.Replace(o, " ", "_", -1)
	return o
}

func SAtoIA(s []string) []interface{} {
	tmp := make([]interface{}, len(s))
	for i, val := range s {
		tmp[i] = val
	}
	return tmp
}
