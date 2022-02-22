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

// TestErrorExclude ... test return error
func TestErrorExclude() error {
	var err error
	i := 10
	if true {
		i++
		return nil
	}

	i++
	err = fmt.Errorf("test codecov %s", "TestErrorExclude")
	i = i + 1
	return err
}

// TestErrorExcludeReturn ... test return error
func TestErrorExcludeReturn() error {
	i := 10
	if true {
		i++
		return nil
	}

	i--
	i++
	return errors.New("retrun error")
}
