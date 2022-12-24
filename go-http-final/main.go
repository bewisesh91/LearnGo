package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

var (
	errorGroup errgroup.Group
)

var trace = []string{"InitialState"}

func main() {
	mapi := &http.Server{
		Addr:           "0.0.0.0:8080",
		Handler:        index(),
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	errorGroup.Go(func() error {
		return mapi.ListenAndServe()
	})

	stopSig := make(chan os.Signal)                         // chan 선언
	signal.Notify(stopSig, syscall.SIGINT, syscall.SIGTERM) // 해당 chan 핸들링 선언, SIGINT, SIGTERM에 대한 메세지 notify
	<-stopSig                                               // 메세지 등록

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // 해당 context 타임아웃 설정, 5초후 server stop
	defer cancel()
	if err := mapi.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	select {
	case <-ctx.Done(): // catching ctx.Done(), timeout of 5 seconds.
		fmt.Println("timeout 5 seconds.")
	}
	fmt.Println("Server stop")
}

// CORS 설정, cross domain을 위해 사용
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		//허용할 header 타입에 대해 열거
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, X-Forwarded-For, Authorization, accept, origin, Cache-Control, X-Requested-With")
		//허용할 method에 대해 열거
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func liteAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c == nil {
			c.Abort() // 미들웨어에서 사용, 이후 요청 중지
			return
		}
		//http 헤더내 "Authorization" 폼의 데이터를 조회
		auth := c.GetHeader("Authorization")
		//실제 인증기능이 올수있다. 현재는 단순히 출력 기능만 처리
		if len(trace) == 1 {
			trace = append(trace, auth)
		}
		c.Next() // 다음 요청 진행
	}
}

func index() *gin.Engine {
	r := gin.New() //gin 선언

	r.Use(gin.Logger())   //gin 내부 log, logger 미들웨어 사용 선언
	r.Use(gin.Recovery()) //gin 내부 recover, recovery 미들웨어 사용 - 패닉복구
	r.Use(CORS())         //crossdomain 미들웨어 사용 등록

	//원하는 만큼 미들웨어 추가가능
	r.GET("/index", indexPage(), checkIndexPage) // 함수, 함수를 할당한 변수 모두 가능

	route := r.Group("/index/login", liteAuth())
	{
		signin := route.Group("signin")
		signin.POST("/post", signIn)

		account := route.Group("account")
		account.GET("/get", getAccountInfo)
		account.PUT("/put", putAccountInfo)
		account.DELETE("/delete", deleteAccountInfo)
	}
	return r
}

func indexPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c == nil {
			c.Abort()
			return
		}
		c.Next()
	}
}

var checkIndexPage = func(c *gin.Context) {
	if c == nil {
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, trace)
	c.Next()
}

func signIn(c *gin.Context) {
	{
		c.JSON(http.StatusOK, gin.H{
			"Login Info": trace,
		})
	}
}

func getAccountInfo(c *gin.Context) {
	var count = len(trace)
	if count != 2 {
		c.JSON(http.StatusForbidden, gin.H{
			"result": "Access Denied",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"trace": trace,
	})
}

func putAccountInfo(c *gin.Context) {
	trace[1] = "Authorization Changed"
	{
		c.JSON(http.StatusOK, gin.H{
			"trace": trace,
		})
	}
}

func deleteAccountInfo(c *gin.Context) {
	trace = append(trace[:1])
}
