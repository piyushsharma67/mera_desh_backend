package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	configPkg "social_web_server/config"
	"social_web_server/db"
	"social_web_server/routes"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/cobra"
)

var port string

var startCMD = &cobra.Command{
	Use: "start",
	Run: func(cmd *cobra.Command, args []string) {
		if env ==""{
			log.Fatal("ENV must be supplied")
		}
		config,err:=configPkg.Loadconfig(env)

		if err!=nil{
			log.Fatal(err.Error())
		}
		conn, err := pgxpool.New(context.Background(), config.GetDSN())
		db.New(conn)
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