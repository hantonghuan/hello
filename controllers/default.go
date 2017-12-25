package controllers

import (
	"github.com/astaxie/beego"
	"strings"
	"strconv"
	"github.com/garyburd/redigo/redis"
	"net"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	req := c.Ctx.Request
	addr := string(req.RemoteAddr)
	ip := strings.Split(addr, ":")[0]
	
	passwd := "hjkl4f80-fds"
	zkey := "z_visit_record"

	conn, _ := redis.Dial("tcp", "kdluck.com:6379")
	conn.Do("Auth", passwd)
	if online() {
		conn.Do("Zincrby", zkey, 1, ip)
	}
	uv, _ := redis.Int(conn.Do("Zcard", zkey))

	ip_map, _ := redis.StringMap(conn.Do("Zrange", zkey, 0, -1, "withscores"))
	var pv int
	for _, value := range ip_map {
		score, _ := strconv.Atoi(value)
		pv += score
	}

	c.Data["uv"] = uv
	c.Data["pv"] = pv
	c.Data["img"] = "http://ws1.sinaimg.cn/large/5febaf6ely1fmt6h7c0b0j20jg0p8kce.jpg"
	c.TplName = "index.tpl"
}

// 判断是否为线上机
func online() bool {
	addrs, _ := net.InterfaceAddrs()
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				if strings.Contains(ipnet.IP.String(), "172.16.0.15") {
					return true
				}
			}
		}
	}
	return false
}