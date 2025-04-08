package cli

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"fracetel/internal/infra"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var rootCmd = &cobra.Command{
	Use: "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Handler was called")
	},
}

func Run() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt)
		<-sigChan
		cancel()
	}()

	cfg := infra.LoadConfigFromEnv()

	mongoClient, _ := mongo.Connect(options.Client().ApplyURI(cfg.DBUrl))
	defer func() {
		if err := mongoClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	mongoDB := mongoClient.Database(cfg.DBName)

	tracksCollection := mongoDB.Collection("tracks")

	loadTracksCmd := newLoadTracksCmd(tracksCollection)

	rootCmd.AddCommand(loadTracksCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error while executing Zero '%s'\n", err)
		os.Exit(1)
	}
}
