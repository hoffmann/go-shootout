package main

import "fmt"
import "os"
import "strconv"

func ack(m int, n int) int{
    if m==0{return n+1;}
    if n==0{return ack(m-1, 1);}
    return ack(m-1, ack(m, n-1));
}

func main(){

    num:= 1;
    if len(os.Args) == 2 {
        i, err := strconv.Atoi(os.Args[1]);
        if err != nil {
            fmt.Printf("conversion failed");
        }
        num = i;
    }
    fmt.Printf("Ack(3,%d): %d\n" ,num, ack(3, num));
}
