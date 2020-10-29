package main
import (
	"fmt"
	"github.com/Unknwon/goconfig"
	"github.com/gin-gonic/gin"
	"net/http"
)
func main() {
	cfg, err := goconfig.LoadConfigFile("conf.ini")
	if err != nil {
		fmt.Println(err)
	}
	viewstr, err := cfg.GetValue("static", "view")
	if err != nil {
		fmt.Println(err)
	}
	port, err := cfg.GetValue("static", "port")
	if err != nil {
		fmt.Println(err)
	}
	router :=gin.Default()
	router.StaticFS("/", http.Dir(viewstr))
	router.Run(port)
}
