package tos

import (
	"fmt"
	"os"
	"time"
	"bufio"
	"strings"
)

type tos struct {

}

type access struct {
	pv int
	ip_map map[string]int
}

func GetEnvInfo() {
	host, _ := os.Hostname()
	fmt.Println("the host name is:",host)
	fmt.Println("the page size is:",os.Getuid())
	fmt.Println("the page size is:",os.Getppid())
}

func GetOpenFunc() {
	var str = [6]string{"[","]","\"","\"","\"","\""}
	for i:= 0 ; i<len(str); i++ {
		fmt.Println("hello os:",str[i])
	}
	var acc *access = new(access)
	file, err := os.Open("./tos/access.log-20200429") // For read access.
	if err != nil {
		fmt.Println("hello os:",err)
	}
	acc.ip_map = map[string]int{}
	defer file.Close()
	if err != nil {
		fmt.Println("hello os:",err)
	}
	var arr []string;
	arr = make([]string, 5)
	time.Sleep(2 * time.Second)
	r := bufio.NewReader(file)
	for {
		line, err := r.ReadString('\n');
		if err != nil{
			fmt.Println("end",err)
			break
		}
		arr_str := strings.Split(line, " ")
		_, ok := acc.ip_map[arr_str[0]]
		if ok {
			acc.ip_map[arr_str[0]] = acc.ip_map[arr_str[0]]+1
		} else {
			acc.ip_map[arr_str[0]] = 1
		}
		acc.pv = acc.pv+1
		arr = append(arr,line)
	}
	fmt.Println("file lien len is:",acc.pv)
	fmt.Println("file lien len is:",arr[10])
}