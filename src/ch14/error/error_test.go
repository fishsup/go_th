package error_test

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

func GetFibList(n int) ([]int, error) {
	if n < 1 || n > 100 {
		// errors.New("something..")
		return nil, fmt.Errorf("The parameter %d cannot be less than 1 or greater than 100", n)

	}
	if n == 1 {
		return []int{1}, nil
	}
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-1]+fibList[i-2])
	}
	return fibList, nil
}

var ErrLessThanOne = errors.New("the parameter cannot be less than 1")
var ErrGreaterThanHundred = errors.New("the parameter cannot be greater than 100")

func GetFibList1(n int) ([]int, error) {
	if n < 1 {
		return nil, ErrLessThanOne

	}
	if n > 100 {
		return nil, ErrGreaterThanHundred
	}
	if n == 1 {
		return []int{1}, nil
	}
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-1]+fibList[i-2])
	}
	return fibList, nil
}

// 及早失败, 避免嵌套
func GetFibList2(str string) ([]int, error) {
	if i, err := strconv.Atoi(str); err == nil {
		if list, err := GetFibList1(i); err == nil {
			return list, nil
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}

func GetFibList3(str string) ([]int, error) {
	var err error
	i, err := strconv.Atoi(str)
	if err != nil {
		return nil, err
	}
	list, err := GetFibList1(i)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func TestGetFibList(t *testing.T) {
	if v, err := GetFibList1(-101); err != nil {
		t.Error(err)
	} else {
		t.Log(v)
	}

	if v, err := GetFibList1(101); err != nil {
		t.Error(err)
	} else {
		t.Log(v)
	}

	t.Log(GetFibList(1))

	t.Log(GetFibList(30))

	t.Log(GetFibList3("12"))
	t.Log(GetFibList3("-12"))
	t.Log(GetFibList3("sdfjk"))
}
