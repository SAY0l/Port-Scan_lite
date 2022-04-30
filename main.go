package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/SAY0l/tcp_scan/version_series"
)

var (
	h bool //help
	v bool //verbose

	u        string //url
	p        int    //port
	para_num int    //parallel_number
	proxy    string //proxy
)

func init() {
	flag.BoolVar(&h, "h", false, "show help")
	flag.BoolVar(&v, "v", false, "verbose")

	// 另一种绑定方式
	//q = flag.Bool("q", false, "suppress non-error messages during configuration testing")

	flag.StringVar(&u, "u", "", "scan url")
	flag.IntVar(&p, "p", 1024, "set the largest port")
	flag.IntVar(&para_num, "para_num", 100, "set parallel_num")
	flag.StringVar(&proxy, "proxy", "", "set proxy")

	// 改变默认的 Usage
	flag.Usage = usage
}

func usage() {
	fmt.Fprintf(os.Stderr, `Go_port_scan_lite version: 1.0

Main Usage: ./Go_port_scan_lite [-h help] [-v verbose] [-u url] [-p port] 

Options:
`)
	flag.PrintDefaults()
}

func main() {
	flag.Parse()
	if h {
		flag.Usage()
	} else if u == "" {
		log.Println("The URL parameter is required.\n Please use -u to enter the URL or use -h to show help")
	}

	if u != "" {
		version_series.Version_6(u, p, para_num, v)
	}
}
