package main

import (
	"belajar/internal/configs"
	memberships "belajar/internal/handlers/memperships"
	membershipRepo "belajar/internal/repository/memberships"
	membershipSvc "belajar/internal/service/memberships"
	"belajar/pkg/internalsql"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	var (
		cfg *configs.Config
	)
	//this is config la
	err := configs.Init(
		configs.WithConfigFolder(
			[]string{"./internal/configs"},
		),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)

	if err != nil {
		log.Fatal("gagal inisiasi config", err)
	}

	cfg = configs.Get()
	log.Println("config", cfg)

	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatal("Gagal inisialisasi database", err)
	}

	membershipRepo := membershipRepo.NewRepository(db)
	membershipSvc := membershipSvc.NewService(membershipRepo)
	membershipHandler := memberships.NewHandler(r, membershipSvc)
	membershipHandler.RegisterRoute()

	r.Run(cfg.Service.Port)
}
