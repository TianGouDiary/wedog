package httpServer

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wedog/pkg/config"
	"wedog/pkg/models"
)

type Server struct {
	conf config.Config
	M    models.Model
}

func NewServer(conf config.Config) *Server {
	if len(conf.Addr) == 0 {
		conf.Addr = "0.0.0.0:80"
	}
	m := models.NewWeDog(
		conf.User,
		conf.Password,
		conf.Ip,
		conf.Port,
		conf.Db)
	return &Server{conf, m}
}

func (s *Server) Run() {
	err := s.M.Open()
	if err != nil {
		panic(err)
	}
	defer s.M.Close()

	engine := gin.New()
	engine.GET("jkl201906112027.info/wedog/getLetter", s.Get)
	engine.POST("jkl201906112027.info/wedog/getLetter", s.Post)
	err = engine.Run(s.conf.Addr)
	if err != nil {
		panic(err)
	}
}

func (s *Server) Get(c *gin.Context) {
	// TODO 获取数据库数量count 1到 count随机取一个，从数据库获取数据
	m, err := s.M.GetRandomOne()
	if err != nil {
		ErrorReturn(c, err)
		return
	}
	c.JSON(http.StatusOK, m)
	return
}

type ErrorMsg struct {
	ErrorCode int
	ErrMsg    string
}

const (
	ErrorPara = 0
)

func ErrorReturn(c *gin.Context, err error) {
	c.JSON(http.StatusOK, ErrorMsg{
		ErrorCode: ErrorPara,
		ErrMsg:    err.Error(),
	})
	return
}

func (s *Server) Post(c *gin.Context) {

}
