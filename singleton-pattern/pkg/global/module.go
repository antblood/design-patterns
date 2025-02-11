package global

import (
	"fmt"

	"go.uber.org/fx"
)

var Module = fx.Module("global", fx.Provide(NewGlobal))

type Global struct {
	fx.Out

	ServerURL string
}

func NewGlobal() Global {
	fmt.Println("SUCCESS: creating global variable")
	return Global{
		ServerURL: "https://example.com",
	}
}
