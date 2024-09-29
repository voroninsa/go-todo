package common

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func UrlToDate(date string) (time.Time, error) {
	str := strings.Split(date, "/due/")[1]

	number, err := strconv.Atoi(str)
	count := 0
	for number != 0 {
		number /= 10
		count += 1
	}
	if count != 8 || err != nil {
		return time.Time{}, errors.New("incorrect date: date must be like YYYYMMDD")
	}

	year, _ := strconv.Atoi(str[:4])
	month, _ := strconv.Atoi(str[4:6])
	day, _ := strconv.Atoi(str[6:8])

	if day > 31 {
		return time.Time{}, errors.New("incorrect date: the day should not be more than 31")
	}
	if month < int(time.January) || month > int(time.December) {
		return time.Time{}, errors.New("incorrect date: month entered incorrectly")
	}

	var resDate time.Time
	resDate = resDate.AddDate(year-1, month-1, day-1)

	return resDate, nil
}

func ErrorDatabase(err error) error {
	return fmt.Errorf("error database: %w", err)
}

// Переводит слайс строк в формат "(tag1,tag2,tag3)"
func TagsToSqlQueryString(tags []string) string {
	res := "("
	for _, tag := range tags {
		res += fmt.Sprintf("'%s',", tag)
	}
	res = res[:len(res)-1]
	res += ")"

	return res
}
