package main

import "fmt"
import "os"
import "strconv"

func main(){

    num:= 1;
    if len(os.Args) == 2 {
        i, err := strconv.Atoi(os.Args[1]);
        if err != nil {
            fmt.Printf("conversion failed");
        }
        num = i;
    }
    flags :=[8192+1]bool{};
    count := 0;

    for ; num>0; num--{
        count = 0;
        for i :=2; i< 8192; i++{
            flags[i] = true;
        }

        for i:=2; i <= 8192; i++{
            if flags[i] {
                // remove all multiples of prime i
                for k:=i+i; k <= 8192; k+=i {
                    flags[k] = false;
                }
                count++;
            }
        }
    }
    fmt.Printf("Count: %d\n", count);
}
