package app

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB gorm.DB

func InitializeDB() {
	var connString = fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
	)

	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{})
	if err != nil {
		log.Fatalln("Cannot connect database", err)
	}

	db.AutoMigrate(&Spin{})
	db.AutoMigrate(&Result{})
	DB = *db
}

type Spin struct {
	ID        string         `gorm:"type:varchar(36);primarykey;" json:"id"`
	Name      string         `gorm:"size:191" json:"name"`
	Options   string         `json:"options"`
	Comment   string         `json:"comment"`
	InputHint string         `gorm:"size:191" json:"input_hint"`
	Password  string         `gorm:"size:191" json:"-"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type Result struct {
	ID             string         `gorm:"type:varchar(36);primarykey;" json:"id"`
	Spin           *Spin          `json:"spin,omitempty"`
	SpinID         string         `gorm:"size:36" json:"spin_id"`
	Input          string         `gorm:"size:191" json:"input"`
	SelectedOption string         `gorm:"size:191" json:"selected_option"`
	CreatedAt      time.Time      `json:"created_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
}
