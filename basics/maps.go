package basics

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getLogFile(filePath string) []byte {
	data, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	return data
}

func getDictionary() map[string]string {
	dict := map[string]string{}

	data, err := os.ReadFile("../tmp/dictionary.json")
	if err != nil {
		panic(err)
	}

	temp := []map[string]string{}

	if err := json.Unmarshal(data, &temp); err != nil {
		panic(err)
	}

	for _, m := range temp {
		for k, v := range m {
			dict[k] = v
		}
	}

	return dict
}

func Dictionary(word string) string {
	w, ok := getDictionary()[word]

	if !ok {
		return fmt.Sprintf("%s not found.", word)
	}

	return w
}

func LogParser(filePath string) (map[string]int, error) {
	sum := map[string]int{}

	file := getLogFile(filePath)

	fields := strings.Split(string(file), "\n")

	for i, v := range fields {
		f := strings.Fields(v)
		if len(f) != 2 {
			fmt.Printf("wrong input: %v (line #%d)\n", f, i+1)

			return nil, errors.New(fmt.Sprintf("wrong input: %v (line #%d)", f, i+1))
		}

		domain := f[0]
		visits, err := strconv.Atoi(f[1])
		if visits < 0 || err != nil {
			fmt.Printf("wrong input: %v (line #%d)\n", f[1], i+1)

			return nil, errors.New(fmt.Sprintf("wrong input: %v (line #%d)", f, i+1))
		}

		sum[domain] += visits
	}

	return sum, nil
}
