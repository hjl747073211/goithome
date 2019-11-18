package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"regexp"  //正则
	// "bytes"
	// "github.com/HttpRequest"
	// "fmt"
	// "time"
)

var (
	db orm.Ormer
)


type Ask struct {
	Id int
	Title string
	CreateTime string
	Author string
	Content string `orm:"type(text)"`
	TalkUrl string
}




func init(){
	//  orm.Debug = true //开启调试模式，可以打印SQL语句
	orm.RegisterDataBase("default","mysql","root:root@tcp(127.0.0.1:3306)/test?charset=utf8&loc=Local")
	 orm.RegisterModel(new(Ask))
	 orm.RunSyncdb("default",false,true)
	 db=orm.NewOrm()
}

func AddAsk(ask *Ask)(int64,error){
	id,err:=db.Insert(ask)
	return id,err
}


//获取文章标题
func GetAskTitle(askHtml string) string{
	if askHtml==""{
		return ""
	}
	
	reg:=regexp.MustCompile(`<div class="post_title"><h1>(.*?)</h1>`)
	result:=reg.FindAllStringSubmatch(askHtml,-1)
	if len(result)==0{
		return ""
	}

	return string(result[0][1])
}

//获取时间
func GetAskTime(askHtml string) string{
	if askHtml==""{
		return ""
	}
	reg:=regexp.MustCompile(`<span id="pubtime_baidu">(.*?)</span>`)
	result:=reg.FindAllStringSubmatch(askHtml,-1)
	if len(result)==0{
		return ""
	}
	return string(result[0][1])
}


//获取作者
func GetAskAuthor(askHtml string) string{
	if askHtml==""{
		return ""
	}
	reg:=regexp.MustCompile(`<span id="author_baidu">作者：<strong>(.*?)</strong></span>`)
	result:=reg.FindAllStringSubmatch(askHtml,-1)
	if len(result)==0{
		return ""
	}
	return string(result[0][1])
}


//获取内容
func GetAskContent(askHtml string) string{
	if askHtml==""{
		return ""
	}
	reg:=regexp.MustCompile(`<div class="post_content" id="paragraph">(.*?)</div>`)
	result:=reg.FindAllStringSubmatch(askHtml,-1)
	if len(result)==0{
		return ""
	}
	return string(result[0][1])
}



//获取评论，这里他的评论是个新网页，我这里只拿到他的url，没继续深入爬取
func GetAskTalk(askHtml string) string{
	if askHtml==""{
		return ""
	}
	reg:=regexp.MustCompile(`<iframe.*?data=\"([^\"]*)\".*?></iframe>`)
	result:=reg.FindAllStringSubmatch(askHtml,-1)
	if len(result)==0{
		return ""
	}
	aid:= string(result[0][1])
	url:="http://dyn.ithome.com/comment/"+aid
	return url
}


//获取文章页面的其他文章url
func GetAskUrl(askHtml string) []string{

	reg:=regexp.MustCompile(`<li><a target="_blank" href=\"([^\"]*)\">.*?</a></span></li>`)
	result:=reg.FindAllStringSubmatch(askHtml,-1)
	
	var askUrls []string
	for _,v:=range result{
	 	askUrls = append(askUrls,v[1])
	}
	return askUrls
}

