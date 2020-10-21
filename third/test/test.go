package test

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var addr = "localhost:8443"
//服务端超时时间
//var serverTimeout = time.Second*2

//客户端超时时间
var clientTimeout = time.Second*2

//服务端处理时间
var handlerTime   = time.Millisecond*500


type handler struct {}


func (h *handler)ServeHTTP(w http.ResponseWriter, req *http.Request)  {
	time.Sleep(handlerTime)
	n,err := w.Write([]byte("req ok"))
	if err != nil {
		log.Println("write err:",err)
		return
	}
	log.Println("write ok!",n,"bytes")
	w.WriteHeader(http.StatusOK)
}

//https server
func HttpsServer(clientCrt,serverCrt,serverKey string)  {
	//证书读取配置
	caCert,err := ioutil.ReadFile(clientCrt)
	if err != nil {
		log.Fatal(err)
	}
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(caCert)
	cfg      := &tls.Config{
		ClientAuth: tls.RequireAndVerifyClientCert,
		ClientCAs: certPool,
	}

	//开始监听
	httpSvr := &http.Server{
		Addr:addr,
		Handler:&handler{},
		TLSConfig:cfg,
	}
	log.Printf("listen in https://%s , pid=%d ... \n",addr,os.Getpid())
	log.Fatal(httpSvr.ListenAndServeTLS(serverCrt,serverKey))

}


//https client
func HttpsClient(serverCrt,clientCrt,clientKey string) {
	//客户端证书配置
	caCert,err := ioutil.ReadFile(serverCrt)
	if err != nil {
		log.Fatal(err)
	}
	certPool   := x509.NewCertPool()
	certPool.AppendCertsFromPEM(caCert)
	cert,err := tls.LoadX509KeyPair(clientCrt,clientKey)
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{
		Transport:&http.Transport{
			TLSClientConfig:&tls.Config{
				RootCAs:certPool,
				Certificates:[]tls.Certificate{cert},
			},
		},
		Timeout:clientTimeout,
	}


	//带超时的context
	ctx,cf := context.WithCancel(context.Background())
	errCh  := make(chan error)

	//处理请求
	go func(ctx context.Context,cancelFunc context.CancelFunc,client2 *http.Client) {
		defer cancelFunc()

		req,err  := http.NewRequestWithContext(ctx,"GET",fmt.Sprintf("https://%s",addr),nil)
		if err != nil {
			log.Println("req err:",err)
			return
		}



		var input string
		for {
			fmt.Scanln(&input)
			if input == "exit" {
				errCh <- errors.New("exit")
				return
			}
			log.Printf("req https://%s timeout=%s input=%s \n",addr,clientTimeout,input)

			//开始请求
			resp,err      := client2.Do(req)
			if err != nil {
				log.Println("req err:",err,"resp:",resp)
				continue
			}
			defer resp.Body.Close()

			//获取数据
			data,err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Println("read body err:",err)
				continue
			}
			log.Printf("received status:%s body:%s \n",resp.Status,string(data))
		}
	}(ctx,cf,client)

	Err := <-errCh
	log.Println("main got err:",Err)
}