package controller

import (
	_ "fmt"
	"log"
	"net/http"
	"net/http/cookiejar"
	"strconv"
	"strings"

	"github.com/carloct/profile/model"
	"github.com/carloct/profile/shared/session"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/asaskevich/govalidator"
)

func ItemCreate(w http.ResponseWriter, r *http.Request) {
	sess := session.Instance(r)

	userId := sess.Values["id"].(uint32)
	closetId, err := strconv.Atoi(r.PostFormValue("closet_id"))
	if err != nil {
		log.Fatal("Cannot parse closet id")
	}
	url := r.PostFormValue("url")

	/*if !govalidator.IsURL(url) {
		sess.AddFlash("The url is not valid", nil)
		http.Redirect(w, r, "/closet/", http.StatusBadRequest)
	}*/

	cookieJar, _ := cookiejar.New(nil)
	client := &http.Client{
		Jar: cookieJar,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/48.0.2564.116 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	image, _ := doc.Find("head link[rel=\"image_src\"]").First().Attr("href")
	imageLg := strings.Replace(image, "/thumbnail/", "/view_large/", 1)

	err = model.ItemCreate(userId, closetId, url, imageLg)
	if err != nil {
		log.Fatal(err)
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
