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



/*
 * Complete the 'customSorting' function below.
 *
 * The function is expected to return a STRING_ARRAY.
 * The function accepts STRING_ARRAY strArr as parameter.
 */

func customSorting(strArr []string) []string {
oddArray := []string{""}
    evenArray := []string{""}
    resultArray := []string{""}
    for _, value := range strArr {
        if len(value)%2 != 0 {
            oddArray = append(oddArray, value)
        } else {
            evenArray = append(evenArray, value)
        }
    }
    oddArray = oddArray[1:]
    evenArray = evenArray[1:]
    //gaurav starts
    odd_map := make(map[int][]string)

    for _, value := range oddArray {
        odd_map[len(value)] = append(odd_map[len(value)], value)
    }
    keys := make([]int, 0, len(odd_map))

    for k := range odd_map {
        keys = append(keys, k)
    }
    sort.Ints(keys)

    for _, k := range keys {
        sort.Strings(odd_map[k])
        //fmt.Println(k, odd_map[k])
        for _, v := range odd_map[k] {
            resultArray = append(resultArray, v)
        }
    }
    //gaurav ends
    //gaurav starts
    even_map := make(map[int][]string)

    for _, value := range evenArray {
        even_map[len(value)] = append(even_map[len(value)], value)
    }
    keys = make([]int, 0, len(even_map))

    for k := range even_map {
        keys = append(keys, k)
    }
    sort.Sort(sort.Reverse(sort.IntSlice(keys)))

    for _, k := range keys {
        sort.Strings(even_map[k])
        //fmt.Println(k, even_map[k])
        for _, v := range even_map[k] {
            resultArray = append(resultArray, v)
        }
    }
    //gaurav ends
    resultArray = resultArray[1:]
    return resultArray
}
func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    strArrCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)

    var strArr []string

    for i := 0; i < int(strArrCount); i++ {
        strArrItem := readLine(reader)
        strArr = append(strArr, strArrItem)
    }

    result := customSorting(strArr)

    for i, resultItem := range result {
        fmt.Fprintf(writer, "%s", resultItem)

        if i != len(result) - 1 {
            fmt.Fprintf(writer, "\n")
        }
    }

    fmt.Fprintf(writer, "\n")

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