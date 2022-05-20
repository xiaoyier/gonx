package main

import (
	//gonx "../.."
	"bufio"
	"fmt"
	"gonx"
	"os"
)

func main()  {

	file, err := os.OpenFile("example/access/access.log", os.O_RDONLY, 666)
	if err != nil {
		panic(err)
	}

	format := "[$time_local] remote:$http_x_real_ip real:$remote_addr, forwarded:$http_x_forwarded_for -> $scheme://$host \"$request\" http_code:$status bytes_sent:$body_bytes_sentB UA`$http_user_agent` request_time:$request_time"
	//reader := gonx.NewReader(file, format)
	parser := gonx.NewParser(format)
	fileReader := bufio.NewReader(file)
	count := 0
	for  {
		count++
		line, err := fileReader.ReadString('\n')
		if err != nil {
			fmt.Println(count,"-----------------",err)
			break
		}
		entry, err := parser.ParseString(line)
		if err != nil {
			fmt.Println(count,"++++++++++++++++++",err)
			continue
		}
		fmt.Printf("entry: %v\n", entry)
	}
	//for {
	//	rec, err := reader.Read()
	//	if err == io.EOF{
	//		break
	//	}
	//
	//	fmt.Printf("entry: %+v\n", rec)
	//	// Process the record... e.g.
	//}

}
