package main

import (
  "fmt"
  "os")

func main(){
    fmt.Println(os.Args);
    fmt.Println("Hello world");
    if len(os.Args)>1{
        fmt.Println("Hello world",os.Args[1]);
    }
    os.Exit(-1);
}
