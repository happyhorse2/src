package downloadimage

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "sync"
)

var waitGroup = new(sync.WaitGroup)

func Download(i int, url string){
    res, err := http.Get(url)
    if err != nil || res.StatusCode != 200 {
	fmt.Printf("下载图片%d失败\n", i)
	waitGroup.Done()    
    }
    data, err2 := ioutil.ReadAll(res.Body)  
    if err2 != nil {  
        fmt.Printf("下载图片%d失败\n", i)
	waitGroup.Done()
    }
    err3 := ioutil.WriteFile(fmt.Sprintf("/home/mayansong/image/1_%d.jpg", i), data, 0777)
    if err3 != nil{
        fmt.Printf("下载图片%d失败\n", i)
	waitGroup.Done()
    }
    waitGroup.Done()
}

func CurrentDownload(urlslice []string){
    for i,url:= range urlslice{
        waitGroup.Add(1) 
	fmt.Printf("开始下载图片%d\n", i+1)
        go Download(i+1, url)
    }
    waitGroup.Wait()
    fmt.Printf("全部下载完成\n")
}
