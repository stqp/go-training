package main

import "errors"

func main() {}

func max(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, errors.New("")
	}
	res := vals[0]
	for _, v := range vals {
		if res < v {
			res = v
		}
	}
	return res, nil
}

func min(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, errors.New("")
	}
	res := vals[0]
	for _, v := range vals {
		if res > v {
			res = v
		}
	}
	return res, nil
}

func max2(val int, vals ...int) (int, error) {
	a, err := max(vals...)
	if err != nil {
		return 0, err
	}
	return max(val, a)
}

func min2(val int, vals ...int) (int, error) {
	a, err := min(vals...)
	if err != nil {
		return 0, err
	}
	return min(val, a)
}
