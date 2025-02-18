package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	configPkg "social_web_server/config"
	"social_web_server/database"
	"social_web_server/enums"
	"social_web_server/repository"
	"social_web_server/routes"
	"social_web_server/services"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var port string

func loadEnv() error{
	err := godotenv.Load()
	if err != nil {
		return err
	}

	return nil
}

func CreateS3Client() (*s3.Client,error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			os.Getenv("AWS_ACCESS_KEY_ID"),
			os.Getenv("AWS_SECRET_ACCESS_KEY"),
			"",
		)),
		config.WithRegion(os.Getenv("AWS_REGION")),
	)

	if err!=nil{
		return nil,err
	}
	return s3.NewFromConfig(cfg),nil
}

var startCMD = &cobra.Command{
	Use: "start",
	Run: func(cmd *cobra.Command, args []string) {
		if env == "" {
			log.Fatal("ENV must be supplied")
		}
		configPrj, err := configPkg.Loadconfig(env)

		if err != nil {
			log.Fatal(err.Error())
		}
		var repo repository.Repositories
		pgxpool, err := pgxpool.New(context.Background(), configPrj.GetDSN())

		if err != nil {
			log.Fatal(err.Error())
		}else{
			queries:= database.New(pgxpool)
			repo=*repository.InitialiseRepositories(enums.Postgres,queries,nil)
		}

		err=loadEnv()

		if err!=nil{
			log.Fatal(err)
		}

		client,err := CreateS3Client()
		presigner := s3.NewPresignClient(client)

		if err!=nil{
			log.Fatal(err)
		}

		s:=&services.ServiceStruct{}

		service:=s.InitialiseService(repo,"mera-desh",client,presigner)

		r := routes.InitRoutes(service)

		addr := fmt.Sprintf(":%s", port)
		fmt.Printf("server running on port %s \n", port)

		if err := http.ListenAndServe(addr, r); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	startCMD.Flags().StringVarP(&port, "port", "p", "8080", "Port to run the server on")
	AddCommand(startCMD)
}
