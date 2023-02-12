package main

import "github.com/nithinsethu/bug-tracking/database"

func main() {
	pg := database.NewPostgresDB()
	pg.AutoMigrate()
}
