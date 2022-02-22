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
func Het() string { // notest
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

var i int

// ErrorExclude ... test return error
func ErrorExclude() error {
	var err error
	if i == 0 {
		err = fmt.Errorf("test codecov %s", "TestErrorExclude")
		return err
	} else if i < 0 {
		return errors.New("test return")
	}
	return nil
}

// ErrorExcludeReturn ... test return error
func ErrorExcludeReturn() error {
	err := ErrorExclude()
	if err != nil {
		return err
	}
	return nil
}
