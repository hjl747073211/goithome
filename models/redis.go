package models

import (
	"github.com/astaxie/goredis"
)

const (
	URL_QUEUE="url_queue"
	URL_VISIT_SET="url_visit_set"
)

var (
	client goredis.Client
)


func ConnectRedis(addr string){
	client.Addr=addr
}
//向队列存数据
func PutinQueue(url string){
	client.Lpush(URL_QUEUE,[]byte(url))
}
//从队列取数据
func PopformQueue() string{
	res,err:=client.Rpop(URL_QUEUE)
	if err!=nil{
		panic(err)
	}
	return string(res)
}
//存储访问过的url
func AddToSet(url string ){
	client.Sadd(URL_VISIT_SET,[]byte(url))
}
//判断是否访问过
func IsVisit(url string)bool{
	isvisit,err:=client.Sismember(URL_VISIT_SET,[]byte(url))
	if err!=nil{
		return false
	}
	return isvisit
} 

func GetQueueLength() int{
	length,err:=client.Llen(URL_QUEUE)
	if err!=nil{
		return 0
	}
	return length
}


