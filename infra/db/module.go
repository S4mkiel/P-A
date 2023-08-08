package db

import (
	"github.com/S4mkiel/p-a/domain/repository"
	"github.com/S4mkiel/p-a/infra/db/src"
	"go.uber.org/fx"
)

var Module = fx.Module("database",
	PostgresModule,
	src.Module,
	fx.Provide(func(s *src.Sources) repository.PaRepository { return s.PaSQL }),
)