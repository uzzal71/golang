package main

import (
	"fmt"
	"maps"
)

func main() {
    // creating maps
    m := make(map[string]string)
    // setting an element
    m["name"] = "golan"
    m["area"] = "backend"

    // get and element
    // fmt.Println(m["name"], m["area"])
    // fmt.Println(m["phone"])

    m2 := make(map[string]int)
    m2["age"] = 30
    m2["price"] = 64
    fmt.Println(m2["age"])
    fmt.Println(m2["phone"])
    fmt.Println(len(m2))
    delete(m2, "price")
    clear(m2)
    fmt.Println(m2)

    st := map[string]string{"name": "Rahim ali", "roll": "458796"}
    st2 := map[string]string{"name": "Rahim ali", "roll": "458796"}
    fmt.Println(st)

    v, ok := st["roll"]
    if ok {
        fmt.Println("all ok", v)
    } else {
        fmt.Println("not ok")
    }

    fmt.Println(maps.Equal(st, st2))
}