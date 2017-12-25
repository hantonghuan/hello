package controllers

import (
	"github.com/astaxie/beego"
	"strings"
	"strconv"
	"github.com/garyburd/redigo/redis"
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

	conn, _ := redis.Dial("tcp", "localhost:6379")
	conn.Do("Auth", passwd)
	conn.Do("Zincrby", zkey, 1, ip)
	uv, _ := redis.Int(conn.Do("Zcard", zkey))

	ip_map, _ := redis.StringMap(conn.Do("Zrange", zkey, 0, -1, "withscores"))
	var pv int
	for _, value := range ip_map {
		score, _ := strconv.Atoi(value)
		pv += score
	}

	c.Data["uv"] = uv
	c.Data["pv"] = pv
	c.TplName = "index.tpl"
}
