package ardidas_insta

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/ahmdrz/goinsta"
	"gocv.io/x/gocv"
)

var insta *goinsta.Instagram

func init() {
	insta = goinsta.New("hack2018acc", "AdidasHack2018")

	if err := insta.Login(); err != nil {
		log.Fatal("Failed logging in to Instagram: ", err)
	}

	// insta.Logout()
}

func GetValidImages(hashtag string) []string {
	// defer insta.Logout()

	var validURLs []string
	var nextMaxID string
	neededHashtags := []string{"#outfit", "#ootd", "#outfitsociety", "#outfitoftheday"}

	xmlFile := "haarcascade_frontalface_default.xml"

	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()

	if !classifier.Load(xmlFile) {
		fmt.Printf("Error reading cascade file: %v\n", xmlFile)
		return
	}

	for len(validURLs) < 10 {
		tags, err := insta.TagFeed(hashtag, nextMaxID)
		if err != nil {
			fmt.Println("Failed getting Tag media from Instagram", err)
		}

		nextMaxID = tags.NextMaxID

		for _, res := range tags.Items {
			if len(res.ImageVersions.Candidates) > 0 {
				// neededHTFound := false

				for _, ht := range neededHashtags {
					if strings.Contains(res.Caption.Text, ht) {
						// neededHTFound = true
						break
					}
				}

				if res.CommentCount > 3 {
					response, e := http.Get(res.ImageVersions.Candidates[0].URL)
					if e != nil {
						fmt.Println("Failed downloading insta image")
						continue
					}

					defer response.Body.Close()

					file, err := os.Create("temp.jpg")
					if err != nil {
						fmt.Println("Failed creating temp file")
						continue
					}

					_, err = io.Copy(file, response.Body)
					if err != nil {
						fmt.Println("Failed copying file contents")
						continue
					}

					file.Close()

					// here we have an image in temp.jpg
					// check for faces!
					img := gocv.IMRead("temp.jpg", gocv.IMReadColor)
					defer img.Close()

					// detect faces
					rects := classifier.DetectMultiScale(img)
					fmt.Printf("found %d faces\n", len(rects))

					// ALTERNATIVE: default cv2 people detector
					// CHECK: params: winStride=(4, 4),
					// 		          padding=(8, 8), scale=1.05
					// hog := gocv.NewHOGDescriptor()
					// hog.SetSVMDetector(gocv.HOGDefaultPeopleDetector())
					// rects := hog.DetectMultiScale(img)

					if len(rects) > 0 {
						fmt.Println("Got image!")
						validURLs = append(validURLs, res.ImageVersions.Candidates[0].URL)
					}
				}
			}
		}
	}

	// for i := range validURLs {
	// 	// fmt.Println("<img src='" + validURLs[i] + "' style='width: 300px'><br>")
	// 	// fmt.Println(validURLs[i])
	// }

	return validURLs
}
