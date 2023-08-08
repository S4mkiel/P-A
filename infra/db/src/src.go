package src

import (
	src "github.com/S4mkiel/p-a/infra/db/src/postgres"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

var Module = fx.Module("src", fx.Provide(NewSQLSourcers))

type Sources struct {
	PaSQL *src.PaSQLRepository
}

func NewSQLSourcers(db *gorm.DB) *Sources {
	var src = Sources{
		PaSQL: src.NewPaRepository(db),
	}
	return &src
}