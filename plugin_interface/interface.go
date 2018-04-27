package raven_interface

type ReplyType int
type ReplyValueType int

const (
    REPLY_GROUP ReplyType = iota
    REPLY_SINGLE
)

const (
    REPLY_VALUE_STR ReplyValueType = iota
    REPLY_VALUE_FLOAT
    REPLY_VALUE_INT
)

type Reply struct {
    Key string
    ValueType ReplyValueType
    Value interface{}
}

type PluginReply struct {
    ReplyType ReplyType
    Reply interface{}
}

