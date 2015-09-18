package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
)

type bonus struct {
	Amount        uint   `json:"amount"`
	Reason        string `json:"reason"`
	ReceiverEmail string `json:"receiver_email"`
}

const (
	server   = "https://bonus.ly/api/v1"
	resource = "bonuses"
)

var (
	token  = flag.String("token", "", "Bonus.ly access token")
	reason = flag.String("reason", "", "reason. e.g. for nyancats #mateship")
	email  = flag.String("email", "", "email address of recipient")
	points = flag.Int("points", 0, "number of points to give")
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Give bonus.ly points, one at a time ;)\n")
		flag.PrintDefaults()

		fmt.Fprintln(os.Stderr, "\nUsage:")
		fmt.Fprintln(os.Stderr, "    bonusly -points 10 -email a@b.com -reason \"for the lulz #mateship\"")
	}
}

func validate() bool {
	return *points > 0 && len(*email) > 0 && len(*reason) > 0 && len(*token) > 0
}

func give(url *string, points int, payload *[]byte) error {
	buf := bytes.NewBuffer(*payload)
	resp, err := http.Post(*url, "application/json", buf)

	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		return fmt.Errorf("Post failed : %+v", err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	var response map[string]interface{}

	if err = json.Unmarshal(body, &response); err != nil {
		return fmt.Errorf("Unable to unmarshal response")
	}

	if response["success"] != "true" {
		return fmt.Errorf("%v", response["message"])
	}

	return nil
}

func main() {
	flag.Parse()

	if !validate() {
		flag.Usage()
		os.Exit(1)
	}

	// create the http parameters
	params := fmt.Sprintf("access_token=%v", *token)

	// create the Bonus message, which will be the payload
	bonus := &bonus{
		Amount:        1,
		Reason:        *reason,
		ReceiverEmail: *email,
	}

	// create the url
	url := fmt.Sprintf("%v/%v?%v", server, resource, params)

	var payload []byte
	var err error

	if payload, err = json.Marshal(&bonus); err != nil {
		fmt.Printf("%v", err)
		os.Exit(2)
	}

	// we need to wait for the goroutines to finish
	var wg sync.WaitGroup

	wg.Add(*points)

	for i := 0; i < *points; i++ {
		go func() {
			defer wg.Done()
			if err := give(&url, 1, &payload); err != nil {
				fmt.Printf("%v\n", err)
			}
		}()
	}

	// don't exit until the goroutines have finished
	wg.Wait()
}
