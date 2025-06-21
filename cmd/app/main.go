package main

import (
	"fmt"

	"github.com/Muhandis2000/online-school/internal/config"
	"github.com/Muhandis2000/online-school/internal/controllers"
	"github.com/Muhandis2000/online-school/internal/db"
	"github.com/Muhandis2000/online-school/internal/routes"
	"github.com/Muhandis2000/online-school/internal/services"
	"github.com/Muhandis2000/online-school/internal/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Загрузка конфигурации
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(fmt.Sprintf("Failed to load config: %v", err))
	}

	// 2. Инициализация логирования
	logger, err := utils.NewLogger(cfg.Log.Directory, cfg.Log.Filename)
	if err != nil {
		panic(fmt.Sprintf("Failed to initialize logger: %v", err))
	}

	// 3. Подключение к базе данных
	dbConn, err := db.NewDB(cfg)
	if err != nil {
		logger.Error.Printf("Failed to connect to database: %v", err)
		panic(err)
	}
	defer dbConn.Close()

	// 4. Инициализация базы данных (миграции)
	if err := db.InitializeDB(dbConn); err != nil {
		logger.Error.Printf("Failed to initialize database: %v", err)
		panic(err)
	}

	// 5. Создание сервисов
	authService := services.NewAuthService(dbConn)
	courseService := services.NewCourseService(dbConn)

	// 6. Создание контроллеров
	authCtrl := controllers.NewAuthController(authService)
	courseCtrl := controllers.NewCourseController(courseService)

	// 7. Настройка маршрутов
	r := gin.Default()
	routes.SetupRoutes(r, authCtrl, courseCtrl)

	// 8. Запуск сервера
	logger.Info.Printf("Server starting on port %s", cfg.Server.Port)
	if err := r.Run(":" + cfg.Server.Port); err != nil {
		logger.Error.Printf("Failed to start server: %v", err)
	}
}
