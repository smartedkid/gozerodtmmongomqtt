/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"2112a-6/month/wsyx/cobra/cmd"
	"2112a-6/month/wsyx/cobra/models"
)

func main() {
	models.MysqlClient()
	models.Connect()
	cmd.Execute()

}
