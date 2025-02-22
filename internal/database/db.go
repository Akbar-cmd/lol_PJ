package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

var DB *gorm.DB

func InitDB() {
	dsn := "postgres://postgres:12345@localhost:5432/main?sslmode=disable"
	var err error
	for i := 0; i < 10; i++ { // Пробуем 10 раз с задержкой
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Println("База данных успешно подключена")
			return
		}

		log.Println("Не удалось подключиться к базе. Повтор через 3 секунды...")
		time.Sleep(3 * time.Second)
	}
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
}
