package main

import (
	"bytes"
	"fmt"
	"github.com/gocolly/colly"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
)

func main() {
	c := colly.NewCollector()
	c.OnRequest(func(request *colly.Request) {
		fmt.Println("visiting:",request.URL.String())
	})

	c.OnResponse(func(response *colly.Response) {
		//fmt.Printf("response: %v",response.Body)
	})

	c.OnHTML(".item_title > a", func(e *colly.HTMLElement) {
		//记录帖子主题
		link := e.Attr("href")
		row := e.Text+" -> "+link+"\n"
		err := AppendToFile(row)
		if err != nil {
			log.Fatal("write error:",err)
		}
	})
	c.OnHTML("#Main > .box > div > table > tbody > tr > td:nth-child(1) > a > img", func(e *colly.HTMLElement) {
		//保存头像
		imgAddr := "https:"+e.Attr("src")
		imgName := GetRandomString()+".png"
		fmt.Println("file:",imgName)
		err := SaveImage(imgName,imgAddr)
		if err != nil {
			log.Fatal("save image error:",err)
		}
	})

	err := c.Visit("https://www.v2ex.com")
	if err != nil {
		fmt.Println("visit error",err)
	}
}

func AppendToFile(data string) error {
	file,err := os.OpenFile("src/spider/spider.log",os.O_WRONLY|os.O_APPEND,0644)
	if err != nil {
		log.Fatal("Open file error:",err)
	}
	_,error := file.WriteString(data)
	return error
}

func SaveImage(imgName,imgAddr string) error {
	path := "src/spider/images/"
	resp,_ := http.Get(imgAddr)

	body,_ := ioutil.ReadAll(resp.Body)

	out,_ := os.Create(path+imgName)

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
