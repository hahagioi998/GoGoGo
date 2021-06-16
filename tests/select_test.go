/**
*   @Author: yky
*   @File: select_test
*   @Version: 1.0
*   @Date: 2021-06-16 21:48
 */
package tests

import (
	"fmt"
	"testing"
	"time"
)

func test1(ch chan string) {
	time.Sleep(time.Second * 5)
	ch <- "test1"
}
func test2(ch chan string) {
	time.Sleep(time.Second * 2)
	ch <- "test2"
}

/**
* @Description: select可以同时监听一个或多个channel，直到其中一个channel ready
* @Param:
* @return:
**/
func Test_Select1(t *testing.T) {
	// 2个管道
	output1 := make(chan string)
	output2 := make(chan string)
	// 跑2个子协程，写数据
	go test1(output1)
	go test2(output2)
	// 用select监控，注意其中一个执行以后程序就直接执行下去了，不会再去接受其他的case
	select {
	case s1 := <-output1:
		fmt.Println("s1=", s1)
	case s2 := <-output2:
		fmt.Println("s2=", s2)
	}
}
