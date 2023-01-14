package api

import (
	"github.com/diwaandrew/kredit/controller/auth"
	"github.com/diwaandrew/kredit/controller/listreport"
	"github.com/diwaandrew/kredit/controller/skalaangsuran"
	"github.com/gin-contrib/cors"
)

func (s *server) SetupRouter() {
	s.Router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "GET", "DELETE", "PUT"},
		AllowHeaders: []string{"*"},
	}))

	authRepo := auth.NewRepository(s.DB)
	authService := auth.NewService(authRepo)
	authHandler := auth.NewHandler(authService)
	s.Router.GET("/getcust", authHandler.GetNasabah)

	skalaRepo := skalaangsuran.NewRepository(s.DB)
	skalahService := skalaangsuran.NewService(skalaRepo)
	skalahHandler := skalaangsuran.NewHandler(skalahService)
	s.Router.GET("/angsuran", skalahHandler.GenerateSkalaAngsuran)

	reportRepo := listreport.NewRepository(s.DB)
	reportService := listreport.NewService(reportRepo)
	reportHandler := listreport.NewHandler(reportService)
	s.Router.GET("/listreport", reportHandler.GetListReport)
	s.Router.PUT("/updateflag", reportHandler.UpdateCustomer)
	s.Router.GET("/branch", reportHandler.GetBranch)
	s.Router.GET("/company", reportHandler.GetCompany)
	s.Router.GET("/search", reportHandler.SearchListReport)

}
