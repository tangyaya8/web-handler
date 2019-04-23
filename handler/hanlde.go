package handler

import (
	"strings"
	. "webHandler/config"
	. "webHandler/route"
)

func DealWithUrl(path string) (pathSlice []string) {
	pathSlice = strings.Split(
		strings.TrimRight(strings.TrimLeft(path,
			UnixSeparator),
			UnixSeparator),
		UnixSeparator)
	return
}

func ExtractURL(path []string) (string, []Param, error) {
	params := make([]Param, 0)
	var realUrl string
	for _, value := range path {
		if strings.HasPrefix(value, RestSeparator) {
			params = append(params,
				Param{
					Name: strings.Trim(value, RestSeparator),
				})
		} else {
			realUrl += UnixSeparator + value
		}
	}
	return realUrl, params, nil
}
