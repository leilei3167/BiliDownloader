package cmd

import (
	"github.com/leilei3167/BiliDownloader/internal/parser"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"runtime"
)

var debug bool = false       //日志级别,默认false
var g int = runtime.NumCPU() //线程数

func init() {
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", debug, "日志级别")
	rootCmd.PersistentFlags().IntVarP(&g, "parallelism", "p", g, "指定并发下载线程数(1到64)")
}

var (
	rootCmd = &cobra.Command{
		Use:   os.Args[0] + " [url]",
		Short: "B站下载器",
		Run: func(cmd *cobra.Command, args []string) {
			logrus.SetLevel(logrus.DebugLevel)
			logrus.Println("running on debug mode")
			if g > 64 || g < 1 {
				logrus.Fatal("线程数只能是1-64!")
			}
			if len(args) == 0 {
				logrus.Fatal("请输入url!")
			}
			logrus.Debugf("成功运行!url:%v", args[0])
			//---------------------------------------------------------------------
			//验证url是否合法,并构建请求Req
			url := args[0]
			req, err := parser.ParseURL(url)
			if err != nil { //解析url出错退出
				logrus.Fatal(err)
			}
			//---------------------------------------------------------------------
			//开启下载器,并发任务调度

		},
	}
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		logrus.Fatal(err)
	}
}
