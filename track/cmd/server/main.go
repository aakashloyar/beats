package main

import (
	//"database/sql"
	"log"
	"net/http"
	//httpartist "github.com/aakashloyar/beats/track/internal/adapters/in/http/artist"
	httptrack "github.com/aakashloyar/beats/track/internal/adapters/in/http/track"
	postgres   "github.com/aakashloyar/beats/track/internal/adapters/out/postgres"
	tracksvc   "github.com/aakashloyar/beats/track/internal/application/service/track"
	//artistsvc  "github.com/aakashloyar/beats/track/internal/application/service/artist"
	"github.com/aakashloyar/beats/track/internal/application/ports/out/system"

)


func main() {

	config := postgres.Config{
		Host     : "",
		Port     : 123, 
		User     : "",
		Password : "",
		DBName   : "",
		SSLMode  : "", 
	}

	clock := system.SystemClock{}
    idGen := system.UUIDGenerator{}

	db, err := config.NewDB()

	if err !=nil {
        log.Fatalf("failed to open DB: %v", err)
	}

	trackRepo := postgres.NewTrackRepository(db)
	//audio_variantsRepo := postgres.NewAudioVariantRepository(db)

	createtrackService := tracksvc.NewCreateTrackService(trackRepo,clock,idGen)

	gettrackService := tracksvc.NewGetTrackService(trackRepo)


	listtracksService := tracksvc.NewListTracksService(trackRepo)

	trackHandler := httptrack.NewHandler(createtrackService,gettrackService, listtracksService)



	mux := http.NewServeMux()
	httptrack.RegisterRoutes(mux,trackHandler)
	//artist.ResisterRoutes(mux,artistHandler)

	http.ListenAndServe(":8080", mux)
}