package main

import (
	//"database/sql"
	httpartist "github.com/aakashloyar/beats/track/internal/adapters/in/http/artist"
	httptrack "github.com/aakashloyar/beats/track/internal/adapters/in/http/track"
	postgres "github.com/aakashloyar/beats/track/internal/adapters/out/postgres"
	"github.com/aakashloyar/beats/track/internal/application/ports/out/system"
	artistsvc "github.com/aakashloyar/beats/track/internal/application/service/artist"
	tracksvc "github.com/aakashloyar/beats/track/internal/application/service/track"
	"log"
	"net/http"
)

func main() {

	config := postgres.Config{
		Host:     "",
		Port:     123,
		User:     "",
		Password: "",
		DBName:   "",
		SSLMode:  "",
	}

	clock := system.SystemClock{}
	idGen := system.UUIDGenerator{}

	db, err := config.NewDB()

	if err != nil {
		log.Fatalf("failed to open DB: %v", err)
	}

	trackRepo := postgres.NewTrackRepository(db)
	artistRepo := postgres.NewArtistRepository(db)

	createtrackService := tracksvc.NewCreateTrackService(trackRepo, idGen, clock)
	gettrackService := tracksvc.NewGetTrackService(trackRepo)
	listtracksService := tracksvc.NewListTracksService(trackRepo)

	trackHandler := httptrack.NewHandler(createtrackService, gettrackService, listtracksService)

	createartistService := artistsvc.NewCreateTrackService(artistRepo, idGen, clock)
	getartistService := artistsvc.NewGetArtistService(artistRepo)

	artistHandler := httpartist.NewHandler(createartistService, getartistService)

	mux := http.NewServeMux()
	httptrack.RegisterRoutes(mux, trackHandler)
	httpartist.RegisterRoutes(mux, artistHandler)
	//artist.ResisterRoutes(mux,artistHandler)

	http.ListenAndServe(":8080", mux)
}
