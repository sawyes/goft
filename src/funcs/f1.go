package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
)

func main() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	
	seedUrl := "https://www.phpip.com"
	resp, err := client.Get(seedUrl)
	defer resp.Body.Close()
	
	if err != nil {
		fmt.Errorf(seedUrl," 请求失败")
		panic(err)
	}
	
	//fmt.Println(resp.TLS.PeerCertificates[0])
	certInfo:=resp.TLS.PeerCertificates[0]
	fmt.Println("过期时间:",certInfo.NotAfter)
	fmt.Println("组织信息:",certInfo.Subject)
	
	
	
}