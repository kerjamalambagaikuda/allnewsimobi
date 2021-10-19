package main

import (
	"fmt"

	utilConfig "github.com/kerjamalambagaikuda/allnewsimobi/util/config"
	"github.com/kerjamalambagaikuda/allnewsimobi/util/logger"
)

func main() {
	// Load Config
	config, err := utilConfig.LoadConfig(".")
	if err != nil {
		fmt.Printf("%+v\n", err)
		return
	}

	// Load Logger
	err = utilConfig.LoadLogger("testing", config)
	if err != nil {
		fmt.Printf("%+v\n", err)
		return
	}

	logger.Log.Info("RUN STATUS-SERVICE!")
}
