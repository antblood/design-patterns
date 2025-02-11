package globalconsumer

import (
	"fmt"

	"go.uber.org/fx"
)

var Module = fx.Module("globalconsumer", fx.Invoke(NewGlobalConsumer))

type GlobalConsumer struct {
	fx.In

	ServerURL string
}

func NewGlobalConsumer(global GlobalConsumer) {
	fmt.Println("SUCCESS: consuming global variable: ", global.ServerURL)
}
