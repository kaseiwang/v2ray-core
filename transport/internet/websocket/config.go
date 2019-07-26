// +build !confonly

package websocket

import (
	"net/http"
	"time"

	"v2ray.com/core/common"
	"v2ray.com/core/transport/internet"
)

const protocolName = "websocket"

func (c *Config) GetNormalizedPath() string {
	path := c.Path
	if path == "" {
		return "/"
	}
	if path[0] != '/' {
		return "/" + path
	}
	return path
}

func (c *Config) GetRequestHeader() http.Header {
	header := http.Header{}
	for _, h := range c.Header {
		header.Add(h.Key, h.Value)
	}
	return header
}

func (c *Config) GetHeartBeatInterval() time.Duration {
	return time.Duration(c.HeartBeatInterval) * time.Millisecond
}

func (c *Config) GetHeartBeatTimeout() time.Duration {
	return time.Duration(c.HeartBeatTimeout) * time.Millisecond
}

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
