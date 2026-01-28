package main

import (
	httpartist "github.com/aakashloyar/beats/track/internal/adapters/in/http/artist"
	httptrack "github.com/aakashloyar/beats/track/internal/adapters/in/http/track"
	httpalbum "github.com/aakashloyar/beats/track/internal/adapters/in/http/album"
	postgres "github.com/aakashloyar/beats/track/internal/adapters/out/postgres"
	"github.com/aakashloyar/beats/track/internal/application/ports/out/system"
	artistsvc "github.com/aakashloyar/beats/track/internal/application/service/artist"
	tracksvc "github.com/aakashloyar/beats/track/internal/application/service/track"
	albumsvc "github.com/aakashloyar/beats/track/internal/application/service/album"
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
	albumRepo := postgres.NewAlbumRepository(db)

	createtrackService := tracksvc.NewCreateTrackService(trackRepo, idGen, clock)
	gettrackService := tracksvc.NewGetTrackService(trackRepo)
	listtracksService := tracksvc.NewListTracksService(trackRepo)

	trackHandler := httptrack.NewHandler(createtrackService, gettrackService, listtracksService)

	createartistService := artistsvc.NewCreateTrackService(artistRepo, idGen, clock)
	getartistService := artistsvc.NewGetArtistService(artistRepo)

	artistHandler := httpartist.NewHandler(createartistService, getartistService)

	createablumService := albumsvc.NewCreateAlbumService(albumRepo, idGen, clock)
	getalbumService := albumsvc.NewGetAlbumService(albumRepo)
	listalbumsService := albumsvc.NewListAlbumsService(albumRepo)

	albumHandler := httpalbum.NewHandler(createablumService, getalbumService, listalbumsService)

	mux := http.NewServeMux()
	httptrack.RegisterRoutes(mux, trackHandler)
	httpartist.RegisterRoutes(mux, artistHandler)
	httpalbum.RegisterRoutes(mux,albumHandler)

	http.ListenAndServe(":8080", mux)
}
