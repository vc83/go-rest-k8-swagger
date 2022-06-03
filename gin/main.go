package main

import (
	"encoding/json"
	"os"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	docs "gin/docs"
)

type QueryParam struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type SuccessResponseStruct struct {
	Message QueryParam `json:"message"`
}

type FailureResponseStruct struct {
	Message string `json:"message"`
}

type ParamStruct struct {
	Params struct{} `json:"params"`
}

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags queryParam
// @Accept json
// @Produce json
// @Param   params     query   string     false    "params"
// @Success 200 {object} SuccessResponseStruct
// @Failure 400  {object} FailureResponseStruct
// @Router /ping [get]
func Ping(c *gin.Context) {

	params := c.Query("params")
	valid := json.Valid([]byte(params))
	if valid {
		queryParam := &QueryParam{}
		json.Unmarshal([]byte(params), queryParam)
		//fmt.Println(queryParam)
		c.JSON(200, gin.H{
			"message": queryParam,
		})
	} else {
		c.JSON(400, gin.H{
			"message": "incorrect query param",
		})
	}

}

func HealthHandler(c *gin.Context) {
	c.JSON(200, "")
}

func ReadinessHandler(c *gin.Context) {
	c.JSON(200, "")
}

func LoadBalancerReq(c *gin.Context) {
	hostname, error := os.Hostname()
	if error != nil {
		c.JSON(200, error)
	} else {
		c.JSON(200, hostname)
	}
}

func main() {
	r := gin.Default()

	docs.SwaggerInfo.Title = "Swagger API"
	docs.SwaggerInfo.Description = "This is a sample server rest server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}
	v1 := r.Group("/api/v1")
	{
		eg := v1.Group("/")

		eg.GET("/ping", Ping)

	}
	//health and readiness API's
	r.GET("/health", HealthHandler)
	r.GET("/readiness", ReadinessHandler)
	r.GET("/gethost", LoadBalancerReq)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// r.Use(cors.New(cors.Config{
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
	// 	AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
	// 	AllowCredentials: true,
	// 	MaxAge:           12 * time.Hour,
	// 	AllowAllOrigins:  true,
	// }))

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
