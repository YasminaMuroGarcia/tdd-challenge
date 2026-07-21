package roman

import (
	"context"
	"regexp"
	"strconv"
)

type RomanClock struct {
	clockClient ClockClient
}

func New(cc ClockClient) *RomanClock {
	return &RomanClock{clockClient: cc}
}

// CurrentRomanTime() returns the time in roman numbers, e.g. XII:XV:IIX (12:15:08)
func (rc *RomanClock) CurrentRomanTime() (string, error) {
	resp, err := rc.clockClient.CurrentTime(context.Background(), &TimeRequest{})
	if err != nil {
		return "", err
	}
	return rc.TimeToRomanTime(resp.GetTime())
}

// TimeToRomanTime() transcribes a time string into roman numbers, e.g. XII:XV:IIX (12:15:08)
func (rc *RomanClock) TimeToRomanTime(time string) (string, error) {
	r := regexp.MustCompile(`^(\d{1,2}):(\d{1,2}):(\d{1,2})$`)
	matches := r.FindStringSubmatch(time)

	// this is bad error handling, improve it
	hours, err := strconv.ParseInt(matches[1], 10, 64)
	minutes, err := strconv.ParseInt(matches[2], 10, 64)
	seconds, err := strconv.ParseInt(matches[3], 10, 64)

	return rc.intToRoman(int(hours)) + ":" + rc.intToRoman(int(minutes)) + ":" + rc.intToRoman(int(seconds)), err
}

func (rc *RomanClock) intToRoman(num int) string {
	values := []int{
		1000, 900, 500, 400,
		100, 90, 50, 40,
		10, 9, 5, 4, 1,
	}

	symbols := []string{
		"M", "CM", "D", "CD",
		"C", "XC", "L", "XL",
		"X", "IX", "V", "IV",
		"I"}
	roman := ""
	i := 0
	for num > 0 {
		k := num / values[i]
		for j := 0; j < k; j++ {
			roman += symbols[i]
			num -= values[i]
		}
		i++
	}
	return roman
}
