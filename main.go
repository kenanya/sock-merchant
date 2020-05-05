package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
	"strings"
	"sort"
)

// Complete the sockMerchant function below.
func sockMerchant(n int32, ar []int32) int32 {

	sort.Slice(ar, func(i, j int) bool{ return ar[i] < ar[j] })
	canBeSold := int32(0)
	if n == 1 { 	return canBeSold	}
	
	for ii:=int32(0); ii<n; ii++ {		
		if ar[ii] == ar[ii+1] {
			canBeSold++			
			ii++
		}
		
		if ii+1 == n-1 {	ii++	}
	} 
	return canBeSold
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := int32(nTemp)

    arTemp := strings.Split(readLine(reader), " ")

    var ar []int32

    for i := 0; i < int(n); i++ {
        arItemTemp, err := strconv.ParseInt(arTemp[i], 10, 64)
        checkError(err)
        arItem := int32(arItemTemp)
        ar = append(ar, arItem)
    }

    result := sockMerchant(n, ar)

    fmt.Fprintf(writer, "%d\n", result)

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
