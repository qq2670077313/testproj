package utils

import (
	"bytes"

	"github.com/spf13/cobra"
)

// func ExecuteCmd(cmd *cobra.Command, args ...string) error {
// 	cmd.SetArgs(args)
// 	return cmd.Execute()
// }

func ExecuteCmd(cmd *cobra.Command, args ...string) (string, error) {
	buf := new(bytes.Buffer)
	// cmd.SetOutput(buf) // deprecated 废弃
	cmd.SetOut(buf) /// 命令成功，rlt=""  出错在err里, 没有必要用
	// cmd.SetErr(buf)
	cmd.SetArgs(args)
	err := cmd.Execute()
	return buf.String(), err
}
