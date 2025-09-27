package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type ServerParams struct {
	LoadAverage int64
	RAMSize     int64
	RAMExp      int64
	DiskSize    int64
	DiskExp     int64
	NetSpeed    int64
	NetExp      int64
}

func main() {
	for {
		resp, err := http.Get("http://srv.msk01.gigacorp.local/_stats")
		if err != nil {
			panic(err)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		err = resp.Body.Close()
		if err != nil {
			panic(err)
		}

		serverParams, err := convSliceToParams(string(body))
		if err != nil {
			panic(err)
		}
		//fmt.Println(serverParams)
		checkValue(serverParams)

	}
}

func checkValue(serverParams ServerParams) {
	if serverParams.LoadAverage > 30 {
		fmt.Printf("Load Average is too high: %d\n", serverParams.LoadAverage)
	}
	if (serverParams.RAMExp*100)/serverParams.RAMSize > 80 {
		fmt.Printf("Memory usage too high: %d%%\n", (serverParams.RAMExp*100)/serverParams.RAMSize)
	}
	if (serverParams.DiskExp*100)/serverParams.DiskSize > 90 {
		fmt.Printf("Free disk space is too low: %d Mb left\n", (serverParams.DiskSize-serverParams.DiskExp)/(1024*1024))
	}
	if (serverParams.NetExp*100)/serverParams.NetSpeed > 90 {
		fmt.Printf("Network bandwidth usage high: %d Mbit/s available\n", (serverParams.NetSpeed-serverParams.NetExp)/(1024*1024))
	}
}

func convSliceToParams(stringParams string) (ServerParams, error) {
	listParams := strings.Split(stringParams, ",")
	var digitList []int
	for _, param := range listParams {
		num, err := strconv.Atoi(strings.TrimSpace(param))
		if err != nil {
			return ServerParams{LoadAverage: -12}, err
		}
		digitList = append(digitList, num)
	}

	serverParams := ServerParams{
		LoadAverage: int64(digitList[0]),
		RAMSize:     int64(digitList[1]),
		RAMExp:      int64(digitList[2]),
		DiskSize:    int64(digitList[3]),
		DiskExp:     int64(digitList[4]),
		NetSpeed:    int64(digitList[5]),
		NetExp:      int64(digitList[6]),
	}

	return serverParams, nil
}
