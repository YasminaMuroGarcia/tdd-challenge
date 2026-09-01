package roman_test

import (
	"errors"

	"github.com/micbar/tdd-challenge/pkg/roman"
	"github.com/micbar/tdd-challenge/pkg/roman/mocks"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ = Describe("Roman", func() {
	Describe("TimeToRomanTime", func() {

		// Task: cd into pkg/roman and run "ginkgo watch"
		// Remove the leading "P" charakters from L22, 31, 36 and 41 to enable the tests one by one
		// Add error handling in clock.go to make the tests pass

		var (
			clock = &roman.RomanClock{}
		)

		It("reports an error when invalid time format is used", func() {
			invalid := []string{"foo", "-5:00:00", ":00:00", "111:00:00", "00:01:02:03"}
			for _, i := range invalid {
				t, err := clock.TimeToRomanTime(i)
				Expect(err).To(HaveOccurred(), "expected '"+i+"' to raise an invalid format error")
				Expect(t).To(Equal(""))
			}
		})

		It("fails if hour is out of range", func() {
			_, err := clock.TimeToRomanTime("24:10:05")
			Expect(err).To(HaveOccurred())
		})

		It("fails if minute is out of range", func() {
			_, err := clock.TimeToRomanTime("23:60:05")
			Expect(err).To(HaveOccurred())
		})

		It("fails if second is out of range", func() {
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
		It("returns an error when the gRPC call fails", func() {
			cc.On("CurrentTime", mock.Anything, mock.Anything).
				Return(nil, errors.New("gRPC error"))

			rtime, err := romanClock.CurrentRomanTime()
			Expect(err).To(HaveOccurred())
			Expect(rtime).To(Equal(""))
			cc.AssertExpectations(GinkgoT())
		})
		It("returns an error when the gRPC response is nil", func() {
			cc.On("CurrentTime", mock.Anything, mock.Anything).
				Return(nil, nil)

			rtime, err := romanClock.CurrentRomanTime()

			Expect(err).To(HaveOccurred())
			Expect(rtime).To(Equal(""))

			cc.AssertExpectations(GinkgoT())
		})
		It("returns an error when the gRPC service is unavailable", func() {
			grpcErr := status.Error(codes.Unavailable, "service unavailable")

			cc.On("CurrentTime", mock.Anything, mock.Anything).
				Return(nil, grpcErr)

			rtime, err := romanClock.CurrentRomanTime()

			Expect(err).To(HaveOccurred())
			Expect(rtime).To(Equal(""))

			cc.AssertExpectations(GinkgoT())
		})
		It("returns an error when the gRPC call returns an invalid time format", func() {
			cc.On("CurrentTime", mock.Anything, mock.Anything).
				Return(&roman.TimeResponse{Time: "invalid"}, nil)

			rtime, err := romanClock.CurrentRomanTime()
			Expect(err).To(HaveOccurred())
			Expect(rtime).To(Equal(""))
			cc.AssertExpectations(GinkgoT())
		})
		It("returns an error when the gRPC call returns a time with out of range hour", func() {
			cc.On("CurrentTime", mock.Anything, mock.Anything).
				Return(&roman.TimeResponse{Time: "24:00:00"}, nil)

			rtime, err := romanClock.CurrentRomanTime()
			Expect(err).To(HaveOccurred())
			Expect(rtime).To(Equal(""))
			cc.AssertExpectations(GinkgoT())
		})
		It("returns an error when the gRPC call returns a time with out of range minute", func() {
			cc.On("CurrentTime", mock.Anything, mock.Anything).
				Return(&roman.TimeResponse{Time: "23:60:00"}, nil)

			rtime, err := romanClock.CurrentRomanTime()
			Expect(err).To(HaveOccurred())
			Expect(rtime).To(Equal(""))
			cc.AssertExpectations(GinkgoT())
		})
		It("returns an error when the gRPC call returns a time with out of range second", func() {
			cc.On("CurrentTime", mock.Anything, mock.Anything).
				Return(&roman.TimeResponse{Time: "23:59:60"}, nil)

			rtime, err := romanClock.CurrentRomanTime()
			Expect(err).To(HaveOccurred())
			Expect(rtime).To(Equal(""))
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
