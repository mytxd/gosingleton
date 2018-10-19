# gosingleton
A simple singleton library for Golang project

# Usage

## install
```
    go get github.com/mytxd/gosingleton
```

```
    package main

    import (
        "fmt"

        singleton "github.com/mytxd/gosingleton"
    )

    // define the working type
    type Person struct {
        Name string
        Age  int8
    }

    func main() {
        // register the type with a name: "person"
        singleton.RegisterType("person", singleton.GetTypeInfo(Person{}))
        
        // get the zero value of specified type
        vf := singleton.GetInstance("person")
        v, ok := vf.(*Person)
        if !ok {
            fmt.Printf("Error when check type")
            return
        }
        fmt.Printf("%v\n", v)
        v.Name = "fuck"
        v.Age = 100

        fmt.Printf("%v\n", v)
        vh := singleton.GetInstance("person")
        v, ok = vh.(*Person)
        if !ok {
            fmt.Printf("Error when check type")
            return
        }
        fmt.Printf("%v\n", v)
    }
```
