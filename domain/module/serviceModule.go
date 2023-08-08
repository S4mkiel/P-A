package module

import (
	"github.com/S4mkiel/p-a/domain/service"
	"go.uber.org/fx"
)

var Service = fx.Module("service",
	fx.Provide(service.NewPaService),
)