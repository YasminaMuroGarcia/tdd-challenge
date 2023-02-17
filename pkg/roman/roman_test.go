package roman_test

import (
	"github.com/micbar/tdd-challenge/pkg/roman"
	"github.com/micbar/tdd-challenge/pkg/roman/mocks"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Roman", func() {
	Describe("TimeToRomanTime", func() {

		// Task: cd into pkg/roman and run "ginkgo watch"
		// Remove the leading "P" charakters from L22, 31, 36 and 41 to enable the tests one by one
		// Add error handling in clock.go to make the tests pass

		var (
			clock = &roman.RomanClock{}
		)

		PIt("reports an error when invalid time format is used", func() {
			invalid := []string{"foo", "-5:00:00", ":00:00", "111:00:00", "00:01:02:03"}
			for _, i := range invalid {
				t, err := clock.TimeToRomanTime(i)
				Expect(err).To(HaveOccurred(), "expected '"+i+"' to raise an invalid format error")
				Expect(t).To(Equal(""))
			}
		})

		PIt("fails if hour is out of range", func() {
			_, err := clock.TimeToRomanTime("24:10:05")
			Expect(err).To(HaveOccurred())
		})

		PIt("fails if minute is out of range", func() {
			_, err := clock.TimeToRomanTime("23:60:05")
			Expect(err).To(HaveOccurred())
		})

		PIt("fails if second is out of range", func() {
			_, err := clock.TimeToRomanTime("23:10:60")
			Expect(err).To(HaveOccurred())
		})

		It("converts 00:00:00", func() {
			time, err := clock.TimeToRomanTime("00:00:00")
			Expect(err).ToNot(HaveOccurred())
			Expect(time).To(Equal("::"))
		})

		It("converts 05:08:25", func() {
			time, err := clock.TimeToRomanTime("05:08:25")
			Expect(err).ToNot(HaveOccurred())
			Expect(time).To(Equal("V:VIII:XXV"))
		})

		It("converts 5:8:25", func() {
			time, err := clock.TimeToRomanTime("05:08:25")
			Expect(err).ToNot(HaveOccurred())
			Expect(time).To(Equal("V:VIII:XXV"))
		})
		It("converts 11:22:59", func() {
			time, err := clock.TimeToRomanTime("11:22:59")
			Expect(err).ToNot(HaveOccurred())
			Expect(time).To(Equal("XI:XXII:LIX"))
		})
		It("converts 23:59:11", func() {
			time, err := clock.TimeToRomanTime("23:59:11")
			Expect(err).ToNot(HaveOccurred())
			Expect(time).To(Equal("XXIII:LIX:XI"))
		})
	})

	Describe("CurrentRomanTime", func() {
		var (
			romanClock *roman.RomanClock
			cc         *mocks.ClockClient
		)

		BeforeEach(func() {
			cc = &mocks.ClockClient{}
			romanClock = roman.New(cc)
		})

		PIt("handles errors when getting the time from the ClockClient", func() {
			//TODO use the mock client to make a hypothetical API call

			_, err := romanClock.CurrentRomanTime()
			Expect(err).To(HaveOccurred())
		})

		PIt("returns the current time", func() {
			//TODO use the mock client to make a hypothetical API call

			rtime, err := romanClock.CurrentRomanTime()
			Expect(err).To(Not(HaveOccurred()))
			Expect(rtime).To(Equal("XXIII:LIX:LIX"))
		})
	})
})

// BONUS POINTS: coverage:
// ```
// ginkgo -cover; go tool cover -html=coverprofile.out
// ```
