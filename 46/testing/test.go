package testing

import (
	"fmt"
	"project_auth_jwt/config"
	"project_auth_jwt/database"
	"project_auth_jwt/helper"
	"project_auth_jwt/testing/generator"
)

// Fungsi error handler
func handlerError(err error) {
	fmt.Printf("Error: %v\n", err)
}

func MainTesting() {
	// Buat instance generator
	gen := generator.NewInfraGenerator(nil)

	// Setup fungsi inisialisasi
	initFunc := map[string]interface{}{
		"config": func(args ...interface{}) (interface{}, error) {
			cfg, err := config.ReadConfig()
			if err != nil {
				return nil, err
			}
			return cfg, nil
		},
		"logger": func(args ...interface{}) (interface{}, error) {
			log, err := helper.InitZapLogger()
			if err != nil {
				return nil, err
			}
			return log, nil
		},
		"database": func(args ...interface{}) (interface{}, error) {
			if len(args) < 1 {
				return nil, fmt.Errorf("missing config parameter for database initialization")
			}
			cfg, err := config.ReadConfig()
			if err != nil {
				return nil, fmt.Errorf(err.Error())
			}
			db, err := database.InitDB(cfg)
			if err != nil {
				return nil, err
			}
			return db, nil
		},
	}

	// Tambahkan path modul
	path := map[string]string{
		"repository": "project_auth_jwt/testing/repository",
		"service":    "project_auth_jwt/testing/service",
	}

	// Setup modules
	gen.SetupModules(initFunc, path)

	// Bangun ServiceContext
	// ctx, err := gen.BuildServiceContext()
	// if err != nil {
	// 	handlerError(err)
	// 	return
	// }
	// fmt.Printf("ServiceContext initialized: %+v\n", ctx)

	// Generate infra file
	outputPath := "infra.go"
	err := gen.GenerateInfraFile(outputPath)
	if err != nil {
		handlerError(err)
		return
	}
	fmt.Printf("Infra file generated at: %s\n", outputPath)
}
