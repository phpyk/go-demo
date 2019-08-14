package main

import (
	"bytes"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)
var waitGroup = new(sync.WaitGroup)


//const HOST = "https://www.v2ex.com/recent?p=" 这个页面需要登录才能访问，还是不爬了吧，免得被封号
const HOST = "https://www.v2ex.com/?p="

func main() {
	now := time.Now()
	for p := 1; p <= 5; p++ {
		waitGroup.Add(1)
		go start(p)
	}
	//等待所有协程操作完成
	waitGroup.Wait();
	fmt.Printf("下载总时间:%v\n", time.Now().Sub(now))
}

func start(page int) {
	c := colly.NewCollector(
		colly.Debugger(&debug.LogDebugger{}),
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.100 Safari/537.36"),
		colly.AllowURLRevisit(),
		)
	c.OnRequest(func(request *colly.Request) {
		fmt.Println("visiting:",request.URL.String())
	})

	c.OnResponse(func(response *colly.Response) {
		fmt.Printf("response_status: %v",response.StatusCode)
	})

	c.OnHTML(".item_title > a", func(e *colly.HTMLElement) {
		//记录帖子主题
		link := e.Attr("href")
		row := time.Now().Format("2006-01-02 15:04:05")+"\t"+e.Text+"\t"+link+"\n"
		//fmt.Println(row)
		err := AppendToFile(row,page)
		if err != nil {
			log.Fatal("write error:",err)
		}
	})

	c.OnHTML("#Main > .box > div > table > tbody > tr > td:nth-child(1) > a > img", func(e *colly.HTMLElement) {
		//保存头像
		imgAddr := "https:"+e.Attr("src")
		imgName := GetRandomString()+".png"
		fmt.Println("file:",imgName)
		err := SaveImage(imgName,imgAddr,page)
		if err != nil {
			log.Fatal("save image error:",err)
		}
	})

	err := c.Visit(HOST+strconv.Itoa(page))
	if err != nil {
		fmt.Println("visit error",err)
	}
	waitGroup.Done()
}

func AppendToFile(data string,page int) error {
	logfile := fmt.Sprintf("src/spider/spider_%d.log",page)
	CheckFileExistsAndCreate(logfile)

	file,err := os.OpenFile(logfile,os.O_WRONLY|os.O_APPEND,0644)
	if err != nil {
		log.Fatal("Open file error:",err)
	}
	_,error := file.WriteString(data)
	return error
}

func SaveImage(imgName,imgAddr string,page int) error {
	dirname := fmt.Sprintf("src/spider/images-%d/",page)
	CheckDirExistsAndCreate(dirname)
	resp,_ := http.Get(imgAddr)

	body,_ := ioutil.ReadAll(resp.Body)

	out,_ := os.Create(dirname+imgName)

	_,err := io.Copy(out,bytes.NewReader(body))
	if err != nil {
		return err
	}
	return nil
}

func GetRandomString() string {
	seed := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	random := make([]byte,6)
	max := len(seed)
	for i := range random {
		random[i] = seed[rand.Intn(max)]
	}
	return string(random)
}

func CheckFileExistsAndCreate(filename string) (r bool) {
	_,err := os.Stat(filename)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		f,er := os.Create(filename)
		if er != nil {
			log.Fatal("创建日志文件失败：",er)
		}
		defer f.Close()
		return false
	}
	return true
}

func CheckDirExistsAndCreate(dirname string) (r bool) {
	d,err := os.Stat(dirname)
	if err != nil {
		er := os.Mkdir(dirname,os.ModePerm)
		if er != nil {
			log.Fatal("创建文件夹失败：",dirname,er)
			return false
		}
		return true
	}
	if d.IsDir() {
		return true
	}
	return true
}