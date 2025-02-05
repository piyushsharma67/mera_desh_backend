package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"social_web_server/db"
	"social_web_server/routes"

	"github.com/spf13/cobra"
)

var port string

var startCMD = &cobra.Command{
	Use: "start",
	Run: func(cmd *cobra.Command, args []string) {
		if env ==""{
			log.Fatal("ENV must be supplied")
			os.Exit(1)
		}
		db.ConnectDB(env)
		r:=routes.InitRoutes()

		addr := fmt.Sprintf(":%s", port)
		fmt.Printf("server running on port %s \n",port)
		
		if err:=http.ListenAndServe(addr,r);err!=nil{
			log.Fatal(err)
		}
	},
}


func init() {
    startCMD.Flags().StringVarP(&port, "port", "p", "8080", "Port to run the server on")
	AddCommand(startCMD)
}