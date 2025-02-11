package main

import (
	"fmt"
	"sync"

	"context"

	"github.com/antblood/design-patterns/singleton-pattern/pkg/global"
	"github.com/antblood/design-patterns/singleton-pattern/pkg/globalconsumer"
	"go.uber.org/fx"
)

// Singleton is the type that we want to have a single instance of.
type Singleton struct {
	data int
}

// instance is a package-level variable that holds the single instance of Singleton.
var instance *Singleton
var once sync.Once

// GetInstance returns the single instance of Singleton, creating it if necessary.
func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{data: 42} // Initialize the singleton instance with some data.
	})
	return instance
}

func run() {
	// Retrieve the singleton instance.
	singleton1 := GetInstance()
	fmt.Println("Singleton 1 data:", singleton1.data)

	// Retrieve the singleton instance again.
	singleton2 := GetInstance()
	fmt.Println("Singleton 2 data:", singleton2.data)

	// Both should point to the same instance.
	if singleton1 == singleton2 {
		fmt.Println("Singleton1 and Singleton2 are the same instance.")
	}

	app := fx.New(
		fx.Options(global.Module),
		fx.Options(globalconsumer.Module),
	)

	app.Start(context.Background())
}

func main() {
	run()
}
