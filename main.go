package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptrace"
	"os"
	"time"

	_ "unsafe"

	"github.com/huof6829/testproj/twitter"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

//go:linkname Uint32 runtime.fastrand
func Uint32() uint32

func main() {

	twitter.Twittermain()

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

	/*
		n := 0
		p := make([]byte, 17)
		// var p []byte
		for n < len(p) {
			val := Uint32()
			fmt.Printf("val=%X\n", val)
			i := 0
			for ; i < 4; i++ {
				if n+i > len(p)-1 {
					break
				}
				p[n+i] = byte(val) /// val >> i*8 不能用变量
				val >>= 8
				fmt.Printf("p[n+i]=%X, i=%d, n=%d\n", p[n+i], i, n)
			}
			n = n + i

			// p[n] = byte(val >> 0)
			// fmt.Printf("p[n] =%X\n", p[n])
			// p[n+1] = byte(val >> 8)
			// p[n+2] = byte(val >> 16)
			// p[n+3] = byte(val >> 24)
			// n = n + 4
		}
		fmt.Printf("%v\n", p)
		fmt.Println(hex.EncodeToString(p))
	*/

	// err := defer1(2)
	// fmt.Println(err)

}

func for_range_error() {
	var output []*int
	nums := []int{1, 2, 3}

	for _, num := range nums {
		output = append(output, &num)
	}

	fmt.Println("Value: ", *output[0], *output[1], *output[2])
	fmt.Println("Address: ", output[0], output[1], output[2])
}

func getnums() []int {
	nums := []int{1, 2, 3}
	fmt.Printf("getnum: %p\n", &nums)
	return nums
}

func defer1(a int) (err error) {
	if a == 1 {
		return fmt.Errorf("11111")
	}
	defer func() error {

		// var err1 error
		if err != nil {
			return func() error {
				fmt.Println("-----")
				return fmt.Errorf("3333")
			}()
		}
		// err1 = err
		err = fmt.Errorf("0000")
		return err
		// err = err1
		// }
		// return err
		// return fmt.Errorf("0000")
		// }
	}()

	if a == 2 {
		return fmt.Errorf("2222")
		// return err
	}
	return nil
}

func tt() {
	// defer func() {
	// 	if r := recover(); r != nil { // 不能放到函数里，捕获不到，oom-sigkill  ctrl+c关闭-sigstop， 都捕获不到
	// 		fmt.Println(r)
	// 	}
	// }()

	// outofmemory2()

	// 自带
	// selflogger, close := initLog()
	// defer close()

	// st, err := disableEscapeHtml(v)  /// 本身字符串带 \ 才能行， 网页数据
	// data, err := json.MarshalIndent(v, "", " ") //  正确，没有 \"   \n  \t   // zap 有
	// if err != nil {
	// 	logger.Fatal("err", zap.Error(err))
	// 	sugerlogger.Error(err)
	// 	// selflogger.Panicf("get %s info occur error: %v", "memory", err)
	// }
	// logger.Info("mem info", zap.ByteString("data", data)) // 用 NewConsoleEncoder，带 \
	// sugerlogger.Infof("sugar mem info: %s", string(data)) // 用 NewConsoleEncoder，换行

	// logger.Info("logger", zap.Reflect("reflect", v)) /// 用 NewJSONEncoder
	// logger.Info("logger", zap.Any("any", v))
	// sugerlogger.Warnw("suger", "infow", v)
	// logger.Warn("logger", zap.Any("memory", v))
	// sugerlogger.Warn(string(debug.Stack()))

	// 自带
	// selflogger.Printf("mem info: %s", st)

	// fmt.Println(hello.Greet())

	// rootCmd := aaa.NewAccountNonce()
	// if _, err := utils.ExecuteCmd(rootCmd, "0"); err != nil {
	// 	os.Exit(1)
	// }
}

// 去除json中的转义字符
func disableEscapeHtml(data interface{}) (string, error) {
	bf := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(false)
	if err := jsonEncoder.Encode(data); err != nil {
		return "", err
	}
	return bf.String(), nil
}

func initLog() (*log.Logger, func()) {
	file, err := os.Create("test.log")
	if err != nil {
		log.Panic("fail to create test.log file!")
	}
	logger := log.New(file, "", log.Llongfile)
	logger.SetFlags(log.LstdFlags)
	return logger, func() { file.Close() }
}

func TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

// 在性能很好但不是很关键的上下文中，使用SugaredLogger。
// 在每一微秒和每一次内存分配都很重要的上下文中，使用Logger。
func initZap() (*zap.Logger, *zap.SugaredLogger, func()) {
	fileName := "zap.log"
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Panic("fail to create zap.log file!")
	}
	hook := lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   1 << 30, //1G
		LocalTime: true,
		Compress:  true,
	}

	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = TimeEncoder
	// encoder.EncodeTime = zapcore.ISO8601TimeEncoder // 设置
	// core := zapcore.NewCore(zapcore.NewJSONEncoder(encoder), syncWriter, zap.NewAtomicLevelAt(zap.DebugLevel))  // log内容 带 \ NewConsoleEncoder  没有 转义字符
	// core := zapcore.NewCore(zapcore.NewJSONEncoder(NewEncoderConfig()),

	core := zapcore.NewCore(zapcore.NewConsoleEncoder(encoder),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
		zap.NewAtomicLevelAt(zap.DebugLevel))
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.WarnLevel))

	zap.ReplaceGlobals(logger)

	sugarLogger := logger.Sugar()

	// func NewEncoderConfig() zapcore.EncoderConfig {
	// 	return zapcore.EncoderConfig{
	// 		// Keys can be anything except the empty string.
	// 		TimeKey:        "T",                           // json时时间键
	// 		LevelKey:       "L",                           // json时日志等级键
	// 		NameKey:        "N",                           // json时日志记录器名
	// 		CallerKey:      "C",                           // json时日志文件信息键
	// 		MessageKey:     "M",                           // json时日志消息键
	// 		StacktraceKey:  "S",                           // json时堆栈键
	// 		LineEnding:     zapcore.DefaultLineEnding,     // 友好日志换行符
	// 		EncodeLevel:    zapcore.CapitalLevelEncoder,   // 友好日志等级名大小写（info INFO）
	// 		EncodeTime:     TimeEncoder,                   // 友好日志时日期格式化
	// 		EncodeDuration: zapcore.StringDurationEncoder, // 时间序列化
	// 		EncodeCaller:   zapcore.ShortCallerEncoder,    // 日志文件信息（包/文件.go:行号）
	// 	}
	// }

	// encoderConfig := zapcore.EncoderConfig{
	// 	TimeKey:        "time",
	// 	LevelKey:       "level",
	// 	NameKey:        "logger",
	// 	CallerKey:      "caller",
	// 	MessageKey:     "msg",
	// 	StacktraceKey:  "stacktrace",
	// 	LineEnding:     zapcore.DefaultLineEnding,
	// 	EncodeLevel:    zapcore.LowercaseLevelEncoder, // 小写编码器
	// 	EncodeTime:     zapcore.ISO8601TimeEncoder,    // ISO8601 UTC 时间格式   // TimeEncoder 标准格式
	// 	EncodeDuration: zapcore.SecondsDurationEncoder,
	// 	EncodeCaller:   zapcore.FullCallerEncoder, // 全路径编码器
	// }
	// config := zap.Config{
	// 	Level:            zap.NewAtomicLevelAt(zap.DebugLevel),                // 日志级别
	// 	Development:      true,                                                // 开发模式，堆栈跟踪
	// 	Encoding:         "json",                                              // 输出格式 console 或 json
	// 	EncoderConfig:    encoderConfig,                                       // 编码器配置
	// 	InitialFields:    map[string]interface{}{"serviceName": "spikeProxy"}, // 初始化字段，如：添加一个服务器名称
	// 	OutputPaths:      []string{"stdout", "./zap.log"},                     // 输出到指定文件 stdout（标准输出，正常颜色） stderr（错误输出，红色）  多路输出
	// 	ErrorOutputPaths: []string{"stderr"},
	// }
	// // 构建日志
	// logger, err := config.Build()
	// if err != nil {
	// 	panic(fmt.Sprintf("log 初始化失败: %v", err))
	// }
	// sugarLogger := logger.Sugar()
	return logger, sugarLogger, func() { file.Close() }
}

func outofmemory1() {
	tracer := &httptrace.ClientTrace{
		ConnectStart: func(network, addr string) {
			log.Println("traced")
		},
	}
	req, err := http.NewRequest("GET", "http://localhost:9999/hello", nil)
	if err != nil {
		log.Printf("error creating request %v", err)
	}

	// Adding the same trace twice causes a stack overflow.
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), tracer))
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), tracer))
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), tracer))
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), tracer))

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Printf("request error: %v", err)
	}
	if res != nil && res.Body != nil {
		res.Body.Close()
	}
}

func outofmemory2() {
	for {
		b := make([]byte, 100000000000) // 输出 signal: killed
		_ = b
	}
}
func outofmemory3() {
	/// out of memory
	bytes := make([]byte, 1)
	for i := 0; i < (10 * 1024 * 1024 * 1024); i++ { // 输出 signal: killed
		bytes = append(bytes, 1)
	}
}
