package main

import (
	"os"

	"github.com/headzoo/surf"
	"github.com/headzoo/surf/browser"
	log "github.com/sirupsen/logrus"
)

func main() {
	bow := surf.NewBrowser()
	err := bow.Open("http://www.reddit.com")
	if err != nil {
		log.Fatalf("browser open failed, err: %v", err)
	}

	// Download the images on the page and write them to files.
	for _, image := range bow.Images() {
		filename := "/tmp" + image.URL.Path
		fout, err := os.Create(filename)
		if err != nil {
			log.Errorf("Error creating file '%s'.", filename)
			continue
		}
		defer fout.Close()

		_, err = image.Download(fout)
		if err != nil {
			log.Printf(
				"Error downloading file '%s'.", filename)
		}
	}

	// Downloading assets asynchronously takes a little more work, but isn't difficult.
	// The DownloadAsync() method takes an io.Writer just like the Download() method,
	// plus an instance of AsyncDownloadChannel. The DownloadAsync() method will send
	// an instance of browser.AsyncDownloadResult to the channel when the download is
	// complete.
	ch := make(browser.AsyncDownloadChannel, 1)
	queue := 0
	for _, image := range bow.Images() {
		filename := "/tmp" + image.URL.Path
		fout, err := os.Create(filename)
		if err != nil {
			log.Errorf("Error creating file '%s'.", filename)
			continue
		}

		image.DownloadAsync(fout, ch)
		queue++
	}

	// Now we wait for each download to complete.
	for {
		select {
		case result := <-ch:
			// result is the instance of browser.AsyncDownloadResult sent by the
			// DownloadAsync() method. It contains the writer which you need to
			// close. It also contains the asset itself, and an error instance if
			// there was an error.
			//result.Writer.Close()
			if result.Error != nil {
				log.Errorf("Error download '%s'. %s\n", result.Asset.Url(), result.Error)
			} else {
				log.Infof("Downloaded '%s'.\n", result.Asset.Url())
			}

			queue--
			if queue == 0 {
				goto FINISHED
			}
		}
	}

FINISHED:
	close(ch)
	log.Infoln("Downloads complete!")
}
