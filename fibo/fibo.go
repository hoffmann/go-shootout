package main

import "os"
import "fmt"
import "strconv"

func fib(n int64) int64 {
    if n < 2{return 1}
    return fib(n-2) + fib(n-1)
}

func main(){
    var num int64 = 1;
    if len(os.Args) == 2 {
        i, err := strconv.Atoi64(os.Args[1]);
        if err != nil {
            fmt.Printf("conversion failed");
        }
        num = i; 
    }
    fmt.Printf("%v\n", fib(num));
}
