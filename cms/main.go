package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"save-tamal/cms/handler"
	"strings"

	"github.com/gorilla/schema"
	"github.com/gorilla/sessions"
	"github.com/spf13/afero"
	"github.com/spf13/viper"
	"github.com/yookoala/realpath"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	collgrpc "save-tamal/proto/collection"
	commgrpc "save-tamal/proto/comments"
	usergrpc "save-tamal/proto/users"
)

func main() {
	config := viper.NewWithOptions(
		viper.EnvKeyReplacer(
			strings.NewReplacer(".", "_"),
		),
	)
	config.SetConfigFile("cms/env/config.yaml")
	config.SetConfigType("yaml")
	config.AutomaticEnv()
	if err := config.ReadInConfig(); err != nil {
		log.Printf("Error loading configuration: %v", err)
	}

	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)
	store := sessions.NewCookieStore([]byte(config.GetString("session.secret")))
	conn, err := grpc.NewClient(
		fmt.Sprintf("%s:%s", config.GetString("tamal.host"), config.GetString("tamal.port")),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal("Connection failed", err)
	}

	wd, err := os.Getwd()
	if err != nil {
		log.Printf("Error loading configuration: %v", err)
	}
	assetPath, err := realpath.Realpath(filepath.Join(wd, "cms/assets"))
	if err != nil {
		log.Printf("Error loading configuration: %v", err)
	}
	asst := afero.NewIOFS(afero.NewBasePathFs(afero.NewOsFs(), assetPath))

	uc := usergrpc.NewUserServiceClient(conn)
	cc := collgrpc.NewCollectionServiceClient(conn)
	cmc := commgrpc.NewCommentServiceClient(conn)
	r := handler.GetHandler(decoder, store, asst, uc, cc, cmc)

	host, port := config.GetString("server.host"), config.GetString("server.port")
	log.Println("Server  starting...")
	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), r); err != nil {
		log.Fatal(err)
	}
}
