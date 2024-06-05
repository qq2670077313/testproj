package aaa

import (
	"testing"

	"go.uber.org/goleak"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}

func TestGreet(t *testing.T) {
	Greet()
}

func TestLeak(t *testing.T) {
	leak()
}

// func TestLeakWithGoleak(t *testing.T) {
//  defer goleak.VerifyNone(t)
// 	leak()
// }

// func emptyRun(*cobra.Command, []string) {}
// func executeCommand(root *cobra.Command, args ...string) (output string, err error) {
// 	_, output, err = executeCommandC(root, args...)
// 	return output, err
// }
// func executeCommandC(root *cobra.Command, args ...string) (c *cobra.Command, output string, err error) {
// 	buf := new(bytes.Buffer)
// 	root.SetOut(buf)
// 	root.SetErr(buf)
// 	root.SetArgs(args)
// 	c, err = root.ExecuteC()
// 	return c, buf.String(), err
// }
// func checkStringContains(t *testing.T, got, expected string) {
// 	if !strings.Contains(got, expected) {
// 		t.Errorf("Expected to contain: \n %v\nGot:\n %v\n", expected, got)
// 	}
// }
// func TestSuggestions(t *testing.T) {
// 	c := &cobra.Command{Use: "c", Run: emptyRun}
// 	c.Flags().BoolP("boola", "a", false, "a boolean flag")

// 	output, err := executeCommand(c, "c", "-a", "--unknown", "flag")
// 	if err == nil {
// 		t.Error("expected unknown flag error")
// 	}
// 	t.Log(output)
// 	t.Log(err.Error())
// 	checkStringContains(t, output, "unknown flag: --unknown")
// }

// func TestCheckMap(t *testing.T) {
// 	CheckMap("aaa", "1")
// }

// func TestNewAccountNonce(t *testing.T) {
// 	require := require.New(t)
// 	cmd := NewAccountNonce()
// 	// rlt, err := utils.ExecuteCmd(cmd)
// 	// // fmt.Println(rlt)
// 	// require.Error(err)

// 	// rlt, err = utils.ExecuteCmd(cmd, "0", "1")
// 	// // fmt.Println(rlt)
// 	// require.Error(err)

// 	// rlt, err = utils.ExecuteCmd(cmd, "0")
// 	// // fmt.Println(rlt)
// 	// require.Error(err)

// 	rlt, err := utils.ExecuteCmd(cmd, "1")
// 	fmt.Printf("rlt=%s, err=%v\n", rlt, err)
// 	require.Error(err)

// 	rlt, err = utils.ExecuteCmd(cmd, "aaa")
// 	fmt.Printf("rlt=%s, err=%v\n", rlt, err)
// 	require.NoError(err)

// 	// err := utils.ExecuteCmd(cmd)
// 	// require.Error(err)
// 	// err = utils.ExecuteCmd(cmd, "0", "1")
// 	// require.Error(err)
// 	// err = utils.ExecuteCmd(cmd, "0")
// 	// require.Error(err)
// 	// err = utils.ExecuteCmd(cmd, "1")
// 	// require.Error(err)
// 	// err = utils.ExecuteCmd(cmd, "aaa")
// 	// require.NoError(err)
// }
