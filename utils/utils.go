package utils

import (
	"bytes"
	_ "unsafe"

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

//go:linkname Uint32 runtime.fastrand
func Uint32() uint32

func Read(p []byte) (n int) {

	// var token []byte
	// token = append(token, []byte("1")...)
	// token := make([]byte, 16)
	// l, err := rand.Read(token)
	// fmt.Println(l)
	// fmt.Println(err)
	// fmt.Println(token)
	// fmt.Println(hex.EncodeToString(token))

	// val := Uint32()
	// m := byte(val)
	// fmt.Printf("val=%X\n", val)
	// fmt.Printf("m=%X\n", m)

	// p := make([]byte, 17)
	// // var p []byte
	for n < len(p) {
		val := Uint32()
		// fmt.Printf("val=%X\n", val)

		i := 0
		for ; i < 4; i++ {
			if n+i > len(p)-1 {
				break
			}
			p[n+i] = byte(val) /// val >> i*8 不能用变量
			val >>= 8
			// fmt.Printf("p[n+i]=%X, i=%d, n=%d\n", p[n+i], i, n)
		}
		n = n + i

		// p[n+1] = byte(val >> 8)
		// p[n+2] = byte(val >> 16)
		// p[n+3] = byte(val >> 24)
		// n = n + 4
	}
	return
	// fmt.Printf("%v\n", p)
	// fmt.Println(hex.EncodeToString(p))
}

func ReadSingle(p []byte) (n int) {
	for ; n < len(p); n++ {
		p[n] = byte(Uint32())
	}
	return
}
