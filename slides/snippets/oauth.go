package snippets

import (
	"log"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"google.golang.org/api/slides/v1"
)

// A group of Google services.
type Services struct {
	Drive  *drive.Service
	Slides *slides.Service
	Sheets *sheets.Service
}

// Gets Google services authenticated with Drive, Slides, and Sheets.
func getServices() *Services {
	ctx := context.Background()
	// Uses env GOOGLE_APPLICATION_CREDENTIALS
	client, err := google.DefaultClient(ctx,
		drive.DriveScope,
		slides.PresentationsScope,
		sheets.SpreadsheetsScope)
	if err != nil {
		log.Fatalf("Error creating Google client: %v", err)
	}
	driveService, err := drive.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Error creating Drive client: %v", err)
	}
	slidesService, err := slides.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Error creating Slides client: %v", err)
	}
	sheetsService, err := sheets.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Error creating Sheets client: %v", err)
	}
	return &Services{
		Drive:  driveService,
		Slides: slidesService,
		Sheets: sheetsService,
	}
}
