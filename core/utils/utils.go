package utils

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/kkdai/youtube/v2"
	"github.com/kkdai/youtube/v2/downloader"
)

const outputPath = "./downloads"

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	videoURL := r.URL.Query().Get("url")

	if videoURL == "" {
		http.Error(w, "URL parameter is required", http.StatusBadRequest)
		return
	}

	parsedURL, err := url.Parse(videoURL)
	if err != nil || parsedURL.Host != "www.youtube.com" {
		http.Error(w, "Invalid YouTube video URL", http.StatusBadRequest)
		return
	}

	getVideoId := func(url string) string {
		return parsedURL.Query().Get("v")
	}

	videoID := getVideoId(videoURL)
	fmt.Println("Video ID:", videoID)

	if isDownloaded(videoID) {
		io.WriteString(w, "Video already downloaded\n")
		return
	}

	go func() {
		err := downloadVideo(videoURL, outputPath)
		if err != nil {
			log.Printf("Error downloading video: %v\n", err)
		} else {
			log.Printf("Video downloaded successfully to %s\n", outputPath)
		}
	}()

	io.WriteString(w, "Download started successfully\n")
}

func downloadVideo(videoURL, outputPath string) error {

	client := &youtube.Client{}

	video, err := client.GetVideo(videoURL)
	if err != nil {
		return fmt.Errorf("error fetching video information: %w", err)
	}
	if video == nil {
		return fmt.Errorf("video is nil")
	}
	dl := downloader.Downloader{
		OutputDir: outputPath,
	}
	format := video.Formats.FindByItag(18)
	if format == nil {
		return fmt.Errorf("unable to find requested format")
	}

	outputFilename := fmt.Sprintf("%s.mp4", video.ID)
	err = dl.Download(context.Background(), video, format, outputFilename)

	if err != nil {
		return fmt.Errorf("error downloading the video: %w", err)
	}

	fmt.Printf("Video downloaded to %s\n", outputPath)
	return nil
}

func isDownloaded(videoID string) bool {
	_, err := os.Stat(fmt.Sprintf("%s/%s.mp4", outputPath, videoID))
	return err == nil
}
