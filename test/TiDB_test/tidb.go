package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Product struct {
	ID    uint   `gorm:"primaryKey;default:auto_random()"`
	Code  string `gorm:"unique"`
	Price uint
}

func main() {
	dsn := "kC45oJcMFwADSBQ.root:TgZ5I3eZEGkevd6x@tcp(gateway01.eu-central-1.prod.aws.tidbcloud.com:4000)/test?charset=utf8mb4&parseTime=True&loc=Local&tls=true"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("connect failed: %v", err)
	}

	if err := db.AutoMigrate(&Product{}); err != nil {
		log.Fatal(err)
	}

	p := &Product{Code: "D42", Price: 100}
	if err := db.Create(p).Error; err != nil {
		log.Fatal(err)
	}
	fmt.Printf("inserted ID=%d\n", p.ID)

	var out Product
	db.First(&out, "code = ?", "D42")
	fmt.Printf("read back: %+v\n", out)
}