package version_series

import (
	"fmt"
	"log"
	"net"
	"sort"
	"sync"
)

func Version_1() {
	_, err := net.Dial("tcp", "scanme.nmap.org:80")
	if err == nil {
		fmt.Println("connect success!")
	}
} //just 1 port

func Version_2() {
	for i := 1; i <= 1024; i++ {
		fmt.Println(i)
		address := fmt.Sprintf("scanme.nmap.org:%d", i)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			continue
		}
		conn.Close()
		fmt.Printf("%d is open\n", i)
	}
} //too slow ,always waiting

func Version_3() {
	for i := 1; i <= 1024; i++ {
		go func(j int) {
			fmt.Println(j)
			address := fmt.Sprintf("scanme.nmap.org:%d", j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				return
			} //shutdown this goroutine
			conn.Close()
			fmt.Printf("%d is open", j)
		}(i)
	}
} //too fast, when main goroutine has done, everything shutdown

func Version_4() {
	var wg sync.WaitGroup
	for i := 1; i <= 1024; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			address := fmt.Sprintf("127.0.0.1:%d", j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("%d is open\n", j)
		}(i)
	}
	wg.Wait()
	fmt.Println("done!!!")
} //When the number of ports is too large, such as 65535, too many threads will be created, which will lead to network and system restrictions

func Version_5_test() {
	ports := make(chan int, 100)
	wg := sync.WaitGroup{}
	for i := 0; i < cap(ports); i++ {
		go worker(ports, &wg)
	}
	for i := 1; i <= 1024; i++ {
		wg.Add(1)
		ports <- i
	}
	wg.Wait()
	close(ports)
}

func worker(ports chan int, wg *sync.WaitGroup) {
	for p := range ports {
		fmt.Println(p)
		wg.Done()
	}
} //Use the channel to specify the thread pool, where up to 100 workers work in parallel. But still use waitgroup

func worker6(ports, results chan int, url string, v bool) {
	for p := range ports {
		if v {
			log.Println("now scanning ....",p)
		}

		address := fmt.Sprintf("%s:%d", url, p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

func Version_6(url string, p, para_num int, v bool) {
	ports := make(chan int, para_num)
	results := make(chan int)
	var openports []int

	for i := 0; i < cap(ports); i++ {
		go worker6(ports, results, url, v)
	}

	go func() {
		for i := 1; i <= p; i++ {
			ports <- i
		}
	}()

	for i := 0; i < p; i++ {
		port := <-results
		if port != 0 {
			fmt.Println("finding:",port)
			openports = append(openports, port)
		}
	}

	close(ports)
	close(results)
	sort.Ints(openports)
	log.Printf("\ntask is finished :)\n\n")
	for _, port := range openports {
		fmt.Printf("%d is open\n", port)
	}
} //The principle of channel forced blocking synchronization is used to replace waitgroup. Use buffered channel to ensure that the maximum number of parallelism is 100
