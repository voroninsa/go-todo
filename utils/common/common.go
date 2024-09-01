package common

import (
	"errors"
	"flag"
	"strconv"
	"strings"
	"time"
)

func UrlToDate(date string) (int, time.Month, int, error) {
	str := strings.Split(date, "/due/")[1]

	number, err := strconv.Atoi(str)
	count := 0
	for number != 0 {
		number /= 10
		count += 1
	}
	if count != 8 || err != nil {
		return 0, 0, 0, errors.New("incorrect date: date must be like YYYYMMDD")
	}

	year, _ := strconv.Atoi(str[:4])
	month, _ := strconv.Atoi(str[4:6])
	day, _ := strconv.Atoi(str[6:8])

	if day > 31 {
		return 0, 0, 0, errors.New("incorrect date: the day should not be more than 31")
	}
	if month < int(time.January) || month > int(time.December) {
		return 0, 0, 0, errors.New("incorrect date: month entered incorrectly")
	}

	return year, time.Month(month), day, nil
}

func ParseFlags() map[string]interface{} {
	config := flag.String("config", "../config.json", "enter destination of project configuration file")
	flag.Parse()

	flags := make(map[string]interface{})
	flags["config"] = *config

	return flags
}
