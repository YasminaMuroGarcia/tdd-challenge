# TDD code challenge

This is an exercise to write a small dummy feature which converts times into roman numbers. The goal of the exercise is to cover all cornercases by enabling all test cases. The main point of it is to use the test cases for real "Test Driven Development".

## Prerequisites

- A working go development environment.
- A locally installed version of [Ginkgo](https://onsi.github.io/ginkgo/)

## Getting started

Open the files `clock.go` and `roman_test.go` in your editor (e.g. https://code.visualstudio.com/).

run

```sh
cd pkg/roman
ginkgo watch
```

You should see something like that

```sh
Identified 1 test suite.  Locating dependencies to a depth of 1 (this may take a while)...
Watching 1 suite:
  . [2 dependencies]
Running Suite: Roman Suite - /Users/mbarz/Development/micbar/tdd-challenge/pkg/roman
====================================================================================
Random Seed: 1676641899

Will run 5 of 11 specs
------------------------------
P [PENDING]
Roman TimeToRomanTime reports an error when invalid time format is used
/Users/mbarz/Development/micbar/tdd-challenge/pkg/roman/roman_test.go:21
------------------------------
P [PENDING]
Roman TimeToRomanTime fails if hour is out of range
/Users/mbarz/Development/micbar/tdd-challenge/pkg/roman/roman_test.go:30
------------------------------
P [PENDING]
Roman TimeToRomanTime fails if minute is out of range
/Users/mbarz/Development/micbar/tdd-challenge/pkg/roman/roman_test.go:35
------------------------------
P [PENDING]
Roman TimeToRomanTime fails if second is out of range
/Users/mbarz/Development/micbar/tdd-challenge/pkg/roman/roman_test.go:40
------------------------------
•••••
------------------------------
P [PENDING]
Roman CurrentRomanTime handles errors when getting the time from the ClockClient
/Users/mbarz/Development/micbar/tdd-challenge/pkg/roman/roman_test.go:85
------------------------------
P [PENDING]
Roman CurrentRomanTime returns the current time
/Users/mbarz/Development/micbar/tdd-challenge/pkg/roman/roman_test.go:92
------------------------------

Ran 5 of 11 Specs in 0.001 seconds
SUCCESS! -- 5 Passed | 0 Failed | 6 Pending | 0 Skipped
PASS
```

## Refactor the code in clock.go

The main task is now to refactor `clock.go`. Remove the leading "P" charakters in `roman_test.go` from L22, 31, 36 and 41 to enable the tests one by one. Ginkgo will run all tests when you save changes.

To solve the challenge, you need to add more error handling to `clock.go` to cover the corner cases.

## Showcase remote error handling

In microservice environments, many calls that have been implemented in functions are now remote calls.
That creates challenges in error handling. To showcase this, the `CurrentRomanTime` function in `clock.go` calls a remote service to get the current time. The remote call is simulated by the `ClockClient` interface. 
You need to implement error handling for this remote call in `CurrentRomanTime`.

## Send me the solution

Please send back your solution.

Be prepared to talk about it and answer some questions about the process and your experience.
