package database

import (
	"time"

	"github.com/wader/gormstore/v2"
)

var Store *gormstore.Store

func InitializeStore() {
	// initialize and setup cleanup
	Store = gormstore.New(DB, []byte("secret")) // db cleanup every hour
	// close quit channel to stop cleanup
	quit := make(chan struct{})
	go Store.PeriodicCleanup(1*time.Hour, quit)
}
