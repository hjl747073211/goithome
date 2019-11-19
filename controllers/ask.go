package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"gozh/models"
	// "time"
	"net/http"
	"net/url"
	"fmt"
)

type AskController struct {
	beego.Controller
}
type Userinfo struct {
	Username string
	Password string
}

var User Userinfo


func (c *AskController) GetAsk() {
	//链接redis
	models.ConnectRedis("127.0.0.1:6379")

	//根url
	// sUrl:="http://localhost/test/111.html"
	sUrl:="https://www.ithome.com/0/457/806.htm"

	models.PutinQueue(sUrl)

	

	for{
		var m models.Ask
		length:=models.GetQueueLength()
		if length==0{
			//如果url队列为空，则退出
			break
		}
		//从队列获取url，获取url对应的html
		sUrl = models.PopformQueue()
		//判断surl是否被访问过
		if models.IsVisit(sUrl){
			continue
		}

		rsp:=httplib.Get(sUrl).SetProxy(func(request *http.Request) (*url.URL, error) {
			u := new(url.URL)
			u.Scheme = "http"
			u.Host = "10.191.131.21:3128" 
			u.User = url.UserPassword("F7688609","w98ZavXy")
			return u, nil
		})

		sHtml,err:=rsp.String()
	
		if err!=nil{
			panic(err)
		}
		//如果获取到title，代表获取成功
		if models.GetAskTitle(sHtml)!=""{
			m.Title   = models.GetAskTitle(sHtml)
			m.CreateTime=models.GetAskTime(sHtml)
			m.Author  = models.GetAskAuthor(sHtml)
			m.Content  = models.GetAskContent(sHtml)
			m.TalkUrl = models.GetAskTalk(sHtml)

			fmt.Println(m)
			id,_:=models.AddAsk(&m)
			fmt.Println(id)
			// if err!=nil{
			// 	panic(err)
			// }
		}
		//提取该页面的所有相关链接url
		urls:=models.GetAskUrl(sHtml)

		//将url循环插入队列中
		for _,url :=range urls{
			models.PutinQueue(url)
		}

		//sUrl 放进访问过的set中
		models.AddToSet(sUrl)



		//防止爬取太快被封，休息1秒
		// time.Sleep(time.Second)
	}



	









}
