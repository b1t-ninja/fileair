package main

import (
	"FileAir/aws3"
	"FileAir/prompt"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/log"
	"os"
)

func main() {
	var confirm bool
	prompt.StartPrompt()

	huh.NewConfirm().
		Title("Are you sure you want to upload " + prompt.FilePath + " ?").
		Affirmative("Yes!").
		Negative("No.").
		Value(&confirm).Run()

	log.Info("Upload confirmed.")

	if confirm {
		url, err := aws3.UploadToS3AndGeneratePresignedURL(prompt.FilePath, prompt.Duration)
		if err != nil {
			log.Error("Error uploading to S3", err)
			os.Exit(1)
		}
		log.Info("Upload successful.\nLink valid for " + prompt.Duration.String())
		log.Info("⬇️ Download Link ⬇️")
		log.Info(url)
	} else {
		log.Info("Upload cancelled.")
	}
}
