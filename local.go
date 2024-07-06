package telebbb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// LocalListen Listen for bot updates local with for loop every second
func (b *TbBot) LocalListen() {
	for {
		time.Sleep(time.Minute * 5)
		r, err := http.Get(fmt.Sprintf(URL, b.token, "getUpdates"))
		if err != nil {
			b.Errors <- err
			time.Sleep(time.Minute)
		}
		if r.StatusCode == http.StatusBadRequest {
			continue
		}
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			b.Errors <- err
			time.Sleep(time.Minute)
		}
		var i interface{}
		if err = json.Unmarshal(d, &i); err != nil {
			b.Errors <- err
			time.Sleep(time.Minute)
		}
		b.Incoming <- i
	}
}
