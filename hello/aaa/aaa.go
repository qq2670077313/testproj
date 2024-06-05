package aaa

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// Greet ... Greet GitHub Actions
func Greet() string {
	leak()

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

// NewAccountNonce ... represents the account nonce command
func NewAccountNonce() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "use",
		Short: "short",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true

			if args[0] == "0" {
				return errors.New("args is empty") // 不带堆栈
			} else if args[0] == "1" {
				return errors.Wrap(errors.New("arg=1"), "wrong arg") // 不带堆栈
			}
			fmt.Printf("cmd is ok, %v\n", args[0])
			return nil
		},
	}
	return cmd
}

var Aliases map[string]string

func init() {
	Aliases = map[string]string{"a": "1", "b": "2", "c": "3", "d": "1", "e": "1"}
}

func CheckMap(alias, addr string) {
	aliases := AliasMap()
	for aliases[addr] != "" { // 去除重复
		delete(Aliases, aliases[addr])
		aliases = AliasMap()
	}
	Aliases[alias] = addr
	fmt.Println("===", Aliases)
}

func AliasMap() map[string]string {
	aliases := make(map[string]string)
	for name, addr := range Aliases {
		aliases[addr] = name
	}
	return aliases
}

func leak() {
	ch := make(chan struct{})
	go func() {
		ch <- struct{}{}
	}()
}
