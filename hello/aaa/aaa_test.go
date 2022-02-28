package aaa

import (
	"fmt"
	"testing"

	"github.com/huof6829/testproj/utils"
	"github.com/stretchr/testify/require"
)

func TestNewAccountNonce(t *testing.T) {
	require := require.New(t)
	cmd := NewAccountNonce()
	// rlt, err := utils.ExecuteCmd(cmd)
	// // fmt.Println(rlt)
	// require.Error(err)

	// rlt, err = utils.ExecuteCmd(cmd, "0", "1")
	// // fmt.Println(rlt)
	// require.Error(err)

	// rlt, err = utils.ExecuteCmd(cmd, "0")
	// // fmt.Println(rlt)
	// require.Error(err)

	rlt, err := utils.ExecuteCmd(cmd, "1")
	fmt.Printf("rlt=%s, err=%v\n", rlt, err)
	require.Error(err)

	rlt, err = utils.ExecuteCmd(cmd, "aaa")
	fmt.Printf("rlt=%s, err=%v\n", rlt, err)
	require.NoError(err)

	// err := utils.ExecuteCmd(cmd)
	// require.Error(err)
	// err = utils.ExecuteCmd(cmd, "0", "1")
	// require.Error(err)
	// err = utils.ExecuteCmd(cmd, "0")
	// require.Error(err)
	// err = utils.ExecuteCmd(cmd, "1")
	// require.Error(err)
	// err = utils.ExecuteCmd(cmd, "aaa")
	// require.NoError(err)
}
