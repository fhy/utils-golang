package utils

import "fmt"

type ClientInfo struct {
	Ip        string
	Platform  string
	Os        string
	UserAgent string
	SessionId string
	UserId    int64
	RequestId int64
}

func (c *ClientInfo) LogFormatShort() string {
	return fmt.Sprintf("%s/%s/%s", c.RequestId, c.UserId, c.SessionId)
}

func (c *ClientInfo) LogFormatShortAndIp() string {
	return fmt.Sprintf("%s/%s/%s/%s", c.RequestId, c.UserId, c.SessionId, c.Ip)
}

func (c *ClientInfo) LogFormatLong() string {
	return fmt.Sprintf("%s/%s/%s/%s/%s/%s/%s", c.RequestId, c.UserId, c.SessionId, c.Ip, c.Platform, c.Os, c.UserAgent)
}
