package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

type Photo struct {
	AlbumID     	uint      `json:"albumId"`
	Id      		uint      `json:"id"`
	Title 			string    `json:"title"`
	Url 			string    `json:"url"`
	ThumbnailUrl	string    `json:"thumbnailUrl"`
}


func getPhoto(albumID uint) (Photo, error){


	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered -> ", r)
		}
	}()

	resp, err := http.Get("https://jsonplaceholder.typicode.com/photos?id="+fmt.Sprint(albumID))

	if err != nil {
		panic(err.Error())
	}


	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}

	var photos []Photo

	err = json.Unmarshal(body, &photos)
	if err != nil{
		panic(err.Error())
	}

	var photo Photo

	if len(photos) > 0{
		photo = photos[0]
	}

	return photo, nil


}


func main(){




	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {


		var albumIds = [50]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50}
		var wg = sync.WaitGroup{}
		var photos [len(albumIds)]Photo

		for i := 0; i<len(albumIds); i++{

			wg.Add(1)
			go func(albumId uint,i int) {
				photo, _ := getPhoto(albumId)
				photos[i] = photo
				wg.Done()

			}(uint(albumIds[i]),i)

		}


		wg.Wait()


		t, _ := template.ParseFiles("aggregatorfinish.html")
		t.Execute(writer, struct {
			Photos [len(albumIds)]Photo
			Len int
		}{
			photos,
			len(photos),
		})


	})


	log.Fatal(http.ListenAndServe(":8080", nil))






}
