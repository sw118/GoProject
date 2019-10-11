package common

import (
	"encoding/json"
	"net/http"
)

type Profile struct {
	Code   int
	Counts []int
	Data   map[int]string
}

func Request() {
	http.HandleFunc("/adam", fool)
	_ = http.ListenAndServe(":8989", nil)
}

func fool(w http.ResponseWriter, r *http.Request) {
	profile := Profile{0,
		[]int{2, 3},
		map[int]string{1: "sunwei", 2: "caozhiwei"}}
	js, err := json.Marshal(profile)
	if err != nil {
		panic(err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
