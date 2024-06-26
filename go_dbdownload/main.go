package main

import "dbdownload/cmd"

func main() {
	// save json files located inside a folder into mongodb
	// cmd.JsonToDB()

	// save db collection to json files
	cmd.DBToJson()

}
