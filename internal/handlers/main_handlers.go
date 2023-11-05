package handlers

// type ShortUrlAnswer struct {
// 	Url string `json:"url"`
// }

// func IndexHandler(rw http.ResponseWriter, req *http.Request) {
// 	id := req.URL.Query().Get("id")
// 	fmt.Println(id)

// 	rw.Header().Set("Content-Type", "application/json")

// 	link := "new link"
// 	err := json.NewEncoder(rw).Encode(&ShortUrlAnswer{link})

// 	if err != nil {
// 		fmt.Println("error")
// 	}

// 	rw.WriteHeader(http.StatusNotFound)
// }
