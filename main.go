package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"net/url"
)

//http.Get
//http.Post
//http.PostForm
//https client.Get()
//new request  client.Do()

const(
	host= "http://localhost:7878"
	hostHttps = "https://localhost:7878"
)

func getRoot(){
	path := "/"
	resp , err := http.Get(fmt.Sprintf("%s%s", host, path))
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	fmt.Println(resp.StatusCode)
	if resp.StatusCode == http.StatusOK{
		defer resp.Body.Close()
		data ,err := ioutil.ReadAll(resp.Body)
		if err != nil{
			fmt.Println(err.Error())
			return
		}
		fmt.Println(string(data))
	}
}

func postForm(){
	path := "/form"
	resp , err := http.PostForm(fmt.Sprintf("%s%s", host, path),url.Values{"name":{"chenchao"},"pass":{"123"}})
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	fmt.Println(resp.StatusCode)
	if resp.StatusCode == http.StatusOK{
		defer resp.Body.Close()
		data ,err := ioutil.ReadAll(resp.Body)
		if err != nil{
			fmt.Println(err.Error())
			return
		}
		fmt.Println(string(data))
	}
}

func getHttps(){
	path := "/"
	//tr := &http.Transport{
	//	TLSClientConfig:&tls.Config{InsecureSkipVerify:true},
	//}
	client := &http.Client{
		//Transport:tr,
	}

	req ,err := http.NewRequest("GET",fmt.Sprintf("%s%s", host, path),nil)
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	req.Header.Add("If-None-Match","/")
	resp, err:= client.Do(req)
	//resp, err:=client.Get(fmt.Sprintf("%s%s", hostHttps, path))
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	fmt.Println(resp.StatusCode)
	if resp.StatusCode == http.StatusOK{
		defer resp.Body.Close()
		data ,err := ioutil.ReadAll(resp.Body)
		if err != nil{
			fmt.Println(err.Error())
			return
		}
		fmt.Println(string(data))
	}
}

func main() {
	getHttps()
	//postForm()
	//getRoot()
}
