package main

import (
	"belajar/internal/configs"
	memberships "belajar/internal/handlers/memperships"
	"belajar/pkg/internalsql"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	var (
		cfg *configs.Config
	)
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

	membershipHandler := memberships.NewHandler(r)
	membershipHandler.RegisterRoute()
	r.Run(cfg.Service.Port)
}
