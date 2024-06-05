package world

import (
	"errors"
	"fmt"
)

// Greet ... Greet GitHub Actions
func Greet() string {
	i := 0
	if i == 1 {
		// notest
		i = 2
		return "111"
	}
	i = 3
	return "Hello GitHub Actions"
}

// Het ... make a test
func WorldHet() string { // notest
	return "Hello GitHub Actions"
}

// Get ... make a test
func Get() string {
	i := 10
	if true {
		i++
	} else {
		// notest
		i--
		i++
	}
	i = i + 1
	return "Hello GitHub Actions"
}

// ErrorExclude ... test return error
func ErrorExclude(i int) error {
	var err error
	if i == 0 {
		fmt.Println("i==0")
		err = fmt.Errorf("test codecov %s", "i==0")
		return err
	} else if i < 0 {
		fmt.Println("i<0")
		return errors.New("i<0")
	}
	fmt.Println("i>0")
	return nil
}

// ErrorExcludeReturn ... test return error
func ErrorExcludeReturn(i int) error {
	err := ErrorExclude(i)
	if err != nil {
		return err
	}
	return nil // test parameter is 0.  excluded by courtney
}

// ErrorExcludeReturnTwo ... test return error
func ErrorExcludeReturnTwo(i int) error {
	err := ErrorExclude(i)
	if err != nil {
		return err // excluded by courtney
	}
	return nil
}
