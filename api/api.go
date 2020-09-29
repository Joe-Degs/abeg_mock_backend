package api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Joe-Degs/abeg_mock_backend/api/database"
	"github.com/Joe-Degs/abeg_mock_backend/api/models"
	"github.com/Joe-Degs/abeg_mock_backend/api/router"
	"github.com/Joe-Degs/abeg_mock_backend/config"
)

func Run() {
	fmt.Printf("\tserver is running on port [::]:%d\n", config.PORT)
	migrateDB()
	listen(config.PORT)
}

// connect to database and automigrate the database models.
func migrateDB() {
	_, err := os.Stat(config.DBNAME)
	if !os.IsNotExist(err) {
		return
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	err = db.Debug().AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal(err)
	}
}

func listen(port int) {
	mux := router.New()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.PORT), mux))
}
