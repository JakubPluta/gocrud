package config

import (
	"github.com/JakubPluta/gocrud/prisma/db"
	"github.com/rs/zerolog/log"
)

func ConnectDB() (*db.PrismaClient, error) {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		return nil, err
	}
	log.Info().Msg("Connected to Prisma")
	return client, nil
}
