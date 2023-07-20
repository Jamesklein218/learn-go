package main

import (
    "fmt"
    "os"
)

func main() {
    args := os.Args[1:]

    var first_name string;
    var last_name string;
    var middle_name string;
    var region string;

    if len(args) < 3 {
        fmt.Printf("Not enough input.");
        os.Exit(0);
    }

    first_name = args[0];
    if len(args) == 3 {
        last_name = args[1];
        region = args[2];
    } else {
        middle_name = args[1];
        last_name =  args[2];
        region = args[3];
    }

    fmt.Println("Output: ");

    switch region {
    case "US":
        fmt.Println(first_name + " " + middle_name + " " + last_name);
    case "VN":
        fmt.Println(last_name + " " + middle_name + " " + first_name);
    }

}

