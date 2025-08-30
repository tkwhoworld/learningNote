package main

import (
	"fmt"
	"flag"
	"os"
	"bufio"
	"io"
	"strconv"
)

var i *string = flag.String("i","unsort.dat","file contains values for sort")
var o *string = flag.String("o","output.dat","file to receive sorted values")
var algorithm *string = flag.String("a","qsort","sort algorithm")

func main(){
	flag.Parse()
	fmt.Println(*i)
	fmt.Println(*o)
	fmt.Println(*algorithm)
	values,err := readInputFile(*i)
	if err!= nil {
		fmt.Println("fialed to open file")
	}
	fmt.Println("read values", values)
	error1 := writeValues(values,*o)
	if error1 != nil {
		fmt.Println("Faild to output file",values)
	}
}

func readInputFile(inputFile string)(values []int, err error) {
	file,err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Faild to open the input file",file)
	}
	defer file.Close()
	values = make([]int,0)
	br := bufio.NewReader(file)
	for {
		line,isPrefix,err1 := br.ReadLine()
		if err1 !=nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}
		if isPrefix {
			fmt.Println("A too long line ,seems unexpected")
			return 
		}
		str := string(line)
		fmt.Println("fialed to open file",str)
		value,err1 := strconv.Atoi(str)
		if err1 != nil{
			err = err1
			return 
		}
		values = append(values,value)
	}
	return 
}

func writeByte(values []byte, outfile string) error {
	file,err := os.Create(outfile)
	if err != nil {
		fmt.Println("create file failed",outfile)
	}
	defer file.Close()
	n,err := file.Write(values)
	fmt.Println("the len is:",n)

	return err
} 

func writeValues(values []int,outfile string) error {
	file,err := os.Create(outfile)
	if err != nil {
			fmt.Println("create file failed",outfile)
	}
	defer file.Close()
	for _,value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str+"\n")
	}

	return nil
}