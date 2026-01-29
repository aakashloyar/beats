package main

import (
	httpalbum "github.com/aakashloyar/beats/track/internal/adapters/in/http/album"
	httpartist "github.com/aakashloyar/beats/track/internal/adapters/in/http/artist"
	httpaudioVariant "github.com/aakashloyar/beats/track/internal/adapters/in/http/audio_variant"
	httptrack "github.com/aakashloyar/beats/track/internal/adapters/in/http/track"
	postgres "github.com/aakashloyar/beats/track/internal/adapters/out/postgres"
	"github.com/aakashloyar/beats/track/internal/application/ports/out/system"
	albumsvc "github.com/aakashloyar/beats/track/internal/application/service/album"
	artistsvc "github.com/aakashloyar/beats/track/internal/application/service/artist"
	audioVariantsvc "github.com/aakashloyar/beats/track/internal/application/service/audio_variant"
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
	//all repositories
	trackRepo := postgres.NewTrackRepository(db)
	artistRepo := postgres.NewArtistRepository(db)
	albumRepo := postgres.NewAlbumRepository(db)
	audioVariantRepo := postgres.NewAudioVariantRepository(db)

	//track services
	createtrackService := tracksvc.NewCreateTrackService(trackRepo, idGen, clock)
	gettrackService := tracksvc.NewGetTrackService(trackRepo)
	listtracksService := tracksvc.NewListTracksService(trackRepo)
	listaudioVariantsByTrackService := tracksvc.NewListAudioVariantsByTrackService(trackRepo)

	//track handler
	trackHandler := httptrack.NewHandler(createtrackService, gettrackService, listtracksService, listaudioVariantsByTrackService)

	//artist services
	createartistService := artistsvc.NewCreateTrackService(artistRepo, idGen, clock)
	getartistService := artistsvc.NewGetArtistService(artistRepo)

	//artist handler
	artistHandler := httpartist.NewHandler(createartistService, getartistService)

	//album services
	createablumService := albumsvc.NewCreateAlbumService(albumRepo, idGen, clock)
	getalbumService := albumsvc.NewGetAlbumService(albumRepo)
	listalbumsService := albumsvc.NewListAlbumsService(albumRepo)

	//album handler
	albumHandler := httpalbum.NewHandler(createablumService, getalbumService, listalbumsService)

	//audio_variant services
	createaudioVariantService := audioVariantsvc.NewCreateAudioVariantService(audioVariantRepo, idGen, clock)

	//audio_variant handler
	audioVariantHandler := httpaudioVariant.NewHandler(createaudioVariantService)

	mux := http.NewServeMux()
	httptrack.RegisterRoutes(mux, trackHandler)
	httpartist.RegisterRoutes(mux, artistHandler)
	httpalbum.RegisterRoutes(mux, albumHandler)
	httpaudioVariant.RegisterRoutes(mux, audioVariantHandler)

	http.ListenAndServe(":8080", mux)
}
