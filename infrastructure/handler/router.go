package handler

import (
	"github.com/MikelSot/tribal-training-back/infrastructure/auth"
	"github.com/MikelSot/tribal-training-back/infrastructure/search"
	"github.com/MikelSot/tribal-training-back/model"
)

func InitRoutes(spec model.Config) {
	// A
	auth.NewRouter(spec)

	// S
	search.NewRouter(spec)
}
