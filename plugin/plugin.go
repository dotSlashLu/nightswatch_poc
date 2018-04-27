package main

import . "github.com/dotSlashLu/nightswatch_poc/plugin/interface"
import "fmt"

func TestCB(report func(string)) {
    fmt.Println("t")
    report("called back from plugin")
}

func TestChan(ch chan *PluginReply) {
    reply := Reply{"reply", REPLY_VALUE_STR, "cb from plugin"}
    ch <- &PluginReply{REPLY_SINGLE, reply}

    ch <- &PluginReply{REPLY_GROUP, []Reply{reply, reply}}
    close(ch)
}

