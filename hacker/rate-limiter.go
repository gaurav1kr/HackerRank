package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)



/*
 * Complete the 'ModuloFibonacciSequence' function below.
 *
 * The function accepts following parameters:
 *  1. chan bool requestChan
 *  2. chan int resultChan
 */

func ModuloFibonacciSequence(quit chan bool, c chan int) {
    x, y := 1, 2
    for {
        select {
        case c <- x%1000000000:
            x, y = y%1000000000, (x%1000000000)+(y%1000000000)
        case <-quit:
            return
        }
    }
    close(c)
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    skipTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    skip := int32(skipTemp)

    totalTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    total := int32(totalTemp)

    resultChan := make(chan int)
    requestChan := make(chan bool)
    //go ModuloFibonacciSequence(requestChan, resultChan)
    go func(){
    for i := int32(0); i < skip + total; i++ {
        
        if i < skip {
            fmt.Errorf("%v",<-resultChan % 1000000000)
        } else {
        fmt.Println(<-resultChan % 1000000000)
        }
    }
    requestChan <- true
    }()
    ModuloFibonacciSequence(requestChan, resultChan)
    close(resultChan)
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




/*sample input :-
0
6


###sample output :- 
1
2
3
5
8
13*/