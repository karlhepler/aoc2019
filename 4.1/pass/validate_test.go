package pass_test

import (
	"testing"

	"github.com/karlhepler/aoc2019/4.1/pass"
)

func TestCheckLength(t *testing.T) {
	var length = 6
	tcs := []struct {
		bs       []byte
		expected bool
	}{
		{nil, false},
		{[]byte(""), false},
		{[]byte("1"), false},
		{[]byte("12"), false},
		{[]byte("1234"), false},
		{[]byte("12345"), false},
		{[]byte("123456"), true},
		{[]byte("1234567"), false},
	}

	for _, tc := range tcs {
		result := pass.CheckLength(tc.bs, length)
		if result != tc.expected {
			t.Errorf("Expected %v; Received %v", tc.expected, result)
		}
	}
}

func TestCheckRange(t *testing.T) {
	var lower = 10
	var upper = 20
	tcs := []struct {
		bs       []byte
		expected bool
	}{
		{nil, false},
		{[]byte(""), false},
		{[]byte("0"), false},
		{[]byte("1"), false},
		{[]byte("10"), false},
		{[]byte("11"), true},
		{[]byte("15"), true},
		{[]byte("19"), true},
		{[]byte("20"), false},
		{[]byte("21"), false},
	}

	for _, tc := range tcs {
		result := pass.CheckRange(tc.bs, lower, upper)
		if result != tc.expected {
			t.Errorf("Expected %v; Received %v", tc.expected, result)
		}
	}
}

func TestCheckAdjacent(t *testing.T) {
	var numAdjacent = 3
	tcs := []struct {
		bs       []byte
		expected bool
	}{
		{nil, false},
		{[]byte(""), false},
		{[]byte("1234"), false},
		{[]byte("1134"), false},
		{[]byte("1114"), true},
		{[]byte("12345332356"), false},
		{[]byte("1235222968"), true},
	}

	for _, tc := range tcs {
		result := pass.CheckAdjacent(tc.bs, numAdjacent)
		if result != tc.expected {
			t.Errorf("Expected %v; Received %v", tc.expected, result)
		}
	}
}

func TestCheckSort(t *testing.T) {
	tcs := []struct {
		bs       []byte
		ob       pass.OrderBy
		expected bool
	}{
		{nil, pass.ASC, true},
		{nil, pass.DESC, true},
		{[]byte(""), pass.ASC, true},
		{[]byte(""), pass.DESC, true},
		{[]byte("1"), pass.ASC, true},
		{[]byte("1"), pass.DESC, true},
		{[]byte("12"), pass.ASC, true},
		{[]byte("12"), pass.DESC, false},
		{[]byte("21"), pass.ASC, false},
		{[]byte("21"), pass.DESC, true},
		{[]byte("1357"), pass.ASC, true},
		{[]byte("2468"), pass.DESC, false},
		{[]byte("7531"), pass.ASC, false},
		{[]byte("8642"), pass.DESC, true},
	}

	for _, tc := range tcs {
		result := pass.CheckSort(tc.bs, tc.ob)
		if result != tc.expected {
			t.Errorf("Expected %v; Received %v", tc.expected, result)
		}
	}
}

func TestValidatorValid(t *testing.T) {
	tcs := []struct {
		bs       []byte
		expected bool
	}{
		{[]byte("111111"), true},
		{[]byte("223450"), false},
		{[]byte("123789"), false},
	}

	v := pass.Validator{
		Length:      6,
		LowerBound:  0,
		UpperBound:  111112,
		NumAdjacent: 2,
		OrderBy:     pass.ASC,
	}

	for _, tc := range tcs {
		result := v.Valid(tc.bs)
		if result != tc.expected {
			t.Errorf("Expected %v; Received %v", tc.expected, result)
		}
	}
}
