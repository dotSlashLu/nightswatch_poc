package main

import "plugin"
import "fmt"
import . "github.com/dotSlashLu/nightswatch_poc/plugin/interface"

func report(word string) {
    fmt.Println("callback with " + word)
}

func funcCB() {
    p, err := plugin.Open("./plugin.so")
    if err != nil {
        panic(err)
    }
    testCB, _ := p.Lookup("TestCB")
    fmt.Println("test cb")
    testCB.(func(func(string)))(report)
}

func chanCB() {
    p, err := plugin.Open("./plugin.so")
    if err != nil {
        panic(err)
    }
    ch := make(chan *PluginReply)
    // defer close(ch)
    testChan, _ := p.Lookup("TestChan")
    go testChan.(func(chan *PluginReply))(ch)
    for {
        recv, ok := <- ch
        if !ok {
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

