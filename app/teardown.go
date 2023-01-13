package app

import (
	"github.com/bmdavis419/go-backend-template/database"
)

func Teardown() {
	// defer closing database
	defer database.CloseMongoDB()
}
