package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	// import 안된다면 go mod vendor
)

type UriParameter struct {
	Name string `uri:"name"`
	Age  string `uri:"age"`
}

func ping(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "pong\n")
}

func get(c *gin.Context) {
	var uri UriParameter
	if err := c.ShouldBindUri(&uri); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"name":   uri.Name,
		"age":    uri.Age,
		"result": "OK",
	})
}

func post(c *gin.Context) {
	name := c.PostForm("name")
	age := c.PostForm("age")
	c.JSON(200, gin.H{
		"name":   name,
		"age":    age,
		"result": "OK",
	})
}

func main() {
	// Step 1) : API 동작 맛보기
	// 0.0.0.0:8080/ping 이라는 주소로 접속하면 ping 함수가 호출된다.
	// http.HandleFunc("/ping", ping)
	// http.ListenAndServe("0.0.0.0:8080", nil)

	// Step 2) : Gin-gonic 사용해보기
	// go get "github.com/gin-gonic/gin"
	// 0.0.0.0:8080/ping 이라는 주소로 접속하면 func(c *gin.Context) 함수가 호출된다.
	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"result": "OK", // result로 "OK"를 반환한다.
	// 	})
	// })
	// r.Run()

	// Step 3) : GET, POST 구핸해보기
	// r := gin.Default()
	// r.GET("/:name/:age", get)
	// r.POST("/:name/:age", post)
	// r.Run()

	// Step 4) : Route Group 설정하기
	// router := gin.Default()
	// login := router.Group("/login")
	// {
	// 	login.GET("/get", get)    // 0.0.0.0:8080/login/get
	// 	login.POST("/post", post) // 0.0.0.0:8080/login/post
	// }
	// join := router.Group("/join")
	// {
	// 	join.GET("/get", get)    // 0.0.0.0:8080/join/get
	// 	join.POST("/post", post) // 0.0.0.0:8080/join/post
	// }
	// router.Run(":8080") // Listen and Serve on 0.0.0.0:8080
}
