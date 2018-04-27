package main

import "plugin"
import "fmt"
import . "github.com/dotslashlu/nightswatch_poc/plugin_interface"

func report(word string) {
    fmt.Println("callback with " + word)
}

func main() {
    p, err := plugin.Open("./plugin.so")
    if err != nil {
        panic(err)
    }
    testCB, _ := p.Lookup("TestCB")
    fmt.Println("test cb")
    testCB.(func(func(string)))(report)

    fmt.Println("test chan")
    ch := make(chan *PluginReply)
    // defer close(ch)
    testChan, _ := p.Lookup("TestChan")
    go testChan.(func(chan *PluginReply))(ch)
    for {
        recv, ok := <- ch
        fmt.Println(ok)
        if !ok {
            fmt.Println(err)
            break
        }
        switch recv.ReplyType {
            case REPLY_GROUP:
                fmt.Printf("recv from ch of group: %v\n", recv)
            case REPLY_SINGLE:
                fmt.Printf("recv from ch of single: %v\n", recv)
        }
    }
}
