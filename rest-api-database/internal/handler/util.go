package handler

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func readContent(filename string, store interface{}) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, store)
}

func contains(in []string, val []string) bool {
	found := 0

	for _, n := range in {
		n = strings.ToLower(n)
		for _, v := range val {
			if n == strings.ToLower(v) {
				found++
				break
			}
		}
	}

	return len(val) == found
}

func containsInt(in []int, val []string) bool {
	found := 0
	for _, _n := range in {
		n := strconv.Itoa(_n)
		for _, v := range val {
			if n == v {
				found++
				break
			}
		}
	}

	return len(val) == found
}
