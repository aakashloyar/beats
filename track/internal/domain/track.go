package domain

import "time"

type Track struct {
	ID            string
	Title         string
	ArtistID      string
	AlbumID       *string
	CoverImageURL *string
	DurationMS    int64
	Language      Language
	ReleaseDate   *time.Time
	CreatedAt     time.Time
}

type TrackFilter struct {
	Title    *string
	ArtistID *string
	AlbumID  *string
	Limit    *string
	Offset   *string
}
