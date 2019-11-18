package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"gozh/models"
	"time"
)

type AskController struct {
	beego.Controller
}

func (c *AskController) GetAsk() {
	//链接redis
	models.ConnectRedis("127.0.0.1:6379")

	//根url
	sUrl:="http://localhost/test/111.html"

	models.PutinQueue(sUrl)

	var m models.Ask

	for{
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

		rsp:=httplib.Get(sUrl)
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

			_,err:=models.AddAsk(&m)
			if err!=nil{
				panic(err)
			}
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
		time.Sleep(time.Second)
	}



	









}
