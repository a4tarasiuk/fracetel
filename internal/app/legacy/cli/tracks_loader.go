package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func newLoadTracksCmd(collection *mongo.Collection) *cobra.Command {
	return &cobra.Command{
		Use:   "load_tracks [file to load]",
		Short: "Load F1 tracks",
		Long:  "Load F1 tracks from file and save it to the storage",
		Args:  cobra.ExactArgs(1),
		RunE:  runLoadTracksCmd(collection),
	}
}

func runLoadTracksCmd(collection *mongo.Collection) func(*cobra.Command, []string) error {
	return func(command *cobra.Command, args []string) error {
		filename := args[0]

		tracks := loadTracksFromFile(filename)

		saveTracksToDB(tracks, collection)

		return nil
	}
}

type track struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func loadTracksFromFile(filename string) []track {
	fmt.Println(filename)

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("err", err)
	}

	tracks := make([]track, 0)

	if err = json.Unmarshal(data, &tracks); err != nil {
		fmt.Printf("parse error: %s\n", err)
	}

	for _, t := range tracks {
		fmt.Printf("%+v\n", t)
	}

	return tracks
}

func saveTracksToDB(tracks []track, collection *mongo.Collection) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for _, t := range tracks {

		_, err := collection.InsertOne(ctx, t)

		if err != nil {
			log.Printf("failed to insert obj to collection: %s", err)
		}
	}
}
