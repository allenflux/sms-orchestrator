package consts

type MessageReadStatus = int

const (
	MessageReadStatusNoRead MessageReadStatus = 1 //未读  = iota
	MessageReadStatusRead   MessageReadStatus = 2 //已读
)
