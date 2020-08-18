package main
import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"fmt"
	"strconv"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)
type Yaml struct {
	Mysql string  `yaml:"mysql"`
}


type Person struct {
	UserId   string    `db:"id"`
	Username string `db:"username"`
}
type Place struct {
	Country string `db:"country"`
	City    string `db:"city"`
	TelCode int    `db:"telcode"`
}
var Db *sqlx.DB
var Config Yaml
func init() {
	conf := new(Yaml)
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	Config = *conf
	database, err := sqlx.Open("mysql", Config.Mysql)
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	Db = database
	//defer Db.Close()  // 注意这行代码要写在上面err判断的下面
}
func main() {
	// Echo instance
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// Routes
	e.GET("/", hello)
	// Start server
	e.Logger.Fatal(e.Start(":6001"))
}

// Handler
func hello(c echo.Context) error {
	var person []Person
	err := Db.Select(&person, "select id, username from user")
	if err != nil {
		fmt.Println("exec failed, ", err)
	}
	return c.String(http.StatusOK, "person num:"+strconv.Itoa(len(person)))
}