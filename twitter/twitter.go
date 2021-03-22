package twitter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	jadwal "github.com/dionarya23/twitter-bot-jadwal-shalat/jadwal_shalat"
)

type Credentials struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

func GetClient(creds *Credentials) (*twitter.Client, error) {
	config := oauth1.NewConfig(creds.ConsumerKey, creds.ConsumerSecret)
	token := oauth1.NewToken(creds.AccessToken, creds.AccessTokenSecret)

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}

	user, _, err := client.Accounts.VerifyCredentials(verifyParams)
	if err != nil {
		return nil, err
	}

	log.Printf("User's ACCOUNT:\n%+v\n", user)
	return client, nil
}

func UpdateStatus() {
	creds := Credentials{
		AccessToken:       os.Getenv("ACCESS_TOKEN"),
		AccessTokenSecret: os.Getenv("ACCESS_TOKEN_SECRET"),
		ConsumerKey:       os.Getenv("CONSUMER_KEY"),
		ConsumerSecret:    os.Getenv("CONSUMER_SECRET"),
	}

	fmt.Printf("Crendetial : %+v\n", creds)

	client, err__ := GetClient(&creds)
	if err__ != nil {
		log.Println("Error getting Twitter Client")
		log.Println(err__)
	}

	fmt.Printf("client : %+v\n", client)

	currentDate := time.Now()
	currentHourMinute := currentDate.Format("15:04")

	var jadwal jadwal.JadwalShalat
	url_endpoint := "https://api.banghasan.com/sholat/format/json/jadwal/kota/679/tanggal/" + currentDate.Format("2006-01-02")
	response, err_ := http.Get(url_endpoint)
	if err_ != nil {
		fmt.Print(err_.Error())
		os.Exit(1)
	}

	responseData, _ := ioutil.ReadAll(response.Body)
	var err = json.Unmarshal(responseData, &jadwal)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	data := jadwal.Jadwal.Data
	v := reflect.ValueOf(data)
	typeOfS := v.Type()

	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Interface() == currentHourMinute && (typeOfS.Field(i).Name == "Subuh" || typeOfS.Field(i).Name == "Dzuhur" || typeOfS.Field(i).Name == "Ashar" || typeOfS.Field(i).Name == "Maghrib" || typeOfS.Field(i).Name == "Isya") {
			tweet, resp, err := client.Statuses.Update("Waktu menunjukan pukul "+currentHourMinute+" WIB\nSaatnya untuk shalat "+typeOfS.Field(i).Name+" untuk kota Bandung \n"+data.Tanggal, nil)
			if err != nil {
				log.Println(err)
			}
			log.Printf("Response : %+v\n", resp)
			log.Printf("Tweet : %+v\n", tweet)
			return
		}
	}
}
