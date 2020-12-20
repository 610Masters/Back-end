package main

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/610masters/Backend/dal/db"
	"github.com/610masters/Backend/dal/model"
)

func main() {
	articles := make([]model.Article, 4)
	users := make([]model.User, 4)
	titles := []string{"Golang学习笔记 2.安装Go1.15版本", "区块链：Golang 开发相关", "深入Golang Runtime之Golang GC的过去,当前与未来","Golang使用WebSocket通信"}
	author := []string{"NBody攻城狮", "跨链技术践行者", "雪东~","大鹏1987"}
	times := []string{"2020-08-21 09:33:59", "2019-05-17 18:44:03", " 2019-09-17 15:52:45","2018-10-27 12:52:03"}
	for i := 0; i < 4; i++ {
		f, err := os.Open(strconv.FormatInt(int64(i), 10) + ".html")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		content, _ := ioutil.ReadAll(f)
		a1 := model.Article{int64(i + 1), titles[i], author[i], nil, times[i], string(content), nil}
		articles = append(articles, a1)
		u := model.User{author[i], "123"}
		users = append(users, u)
	}
	db.PutUsers(users)
	db.PutArticles(articles)
}