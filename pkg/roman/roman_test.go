package roman_test

import (
	"github.com/micbar/tdd-challenge/pkg/roman"
	"github.com/micbar/tdd-challenge/pkg/roman/mocks"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
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
		It("returns invalid format when the hour part is malformed", func() {
			_, err := clock.TimeToRomanTime("2a:10:05")
			Expect(err).To(MatchError(roman.ErrInvalidFormat))
		})
		It("returns invalid format when the minute part is malformed", func() {
			_, err := clock.TimeToRomanTime("23:1b:05")
			Expect(err).To(MatchError(roman.ErrInvalidFormat))
		})
		It("returns invalid format when the second part is malformed", func() {
			_, err := clock.TimeToRomanTime("23:10:0c")
			Expect(err).To(MatchError(roman.ErrInvalidFormat))
		})
	})

	Describe("TimeResponse", func() {
		It("returns an empty string from GetTime when the receiver is nil", func() {
			var resp *roman.TimeResponse
			Expect(resp.GetTime()).To(Equal(""))
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
		It("returns the current time", func() {
			// the hypothetical gRPC call returns a typed response
			cc.On("CurrentTime", mock.Anything, mock.Anything).
				Return(&roman.TimeResponse{Time: "23:59:59"}, nil)

			rtime, err := romanClock.CurrentRomanTime()
			Expect(err).To(Not(HaveOccurred()))
			Expect(rtime).To(Equal("XXIII:LIX:LIX"))
			cc.AssertExpectations(GinkgoT())
		})
		// Add more tests to showcase the error handling of GRPC requests
		// We do not want to test the mock client but we want to see GRPC error handling.
		// Think of which kinds of errors could possibly happen.
	})
})

// BONUS POINTS: coverage:
// ```
// ginkgo -cover; go tool cover -html=coverprofile.out
// ```
