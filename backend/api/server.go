package api

import (
	"os"
	"time"

	"github.com/diwaandrew/kredit/controller/auth"
	"github.com/diwaandrew/kredit/controller/skalaangsuran"
	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
	"gorm.io/gorm"
)

type server struct {
	Router *gin.Engine
	DB     *gorm.DB
}

func MakeServer(db *gorm.DB) *server {
	s := &server{
		Router: gin.Default(),
		DB:     db,
	}

	run1 := gocron.NewScheduler(time.UTC)
	run1.Every(1).Minute().Do(func() {
		auto1 := auth.NewRepository(db)
		auto1.GetNasabah()
	})
	run1.StartAsync()

	run2 := gocron.NewScheduler(time.UTC)
	run2.Every(6).Minute().Do(func() {
		auto2 := skalaangsuran.NewRepository(db)
		auto2.GenerateSkalaAngsuran()
	})
	run2.StartAsync()

	return s
}

func (s *server) RunServer() {
	s.SetupRouter()
	port := os.Getenv("PORT")
	if err := s.Router.Run(":" + port); err != nil {
		panic(err)
	}
}
