package telebbb

import "net/http"

// ServeHook Starts http listener for telegram server
func (s *TbBot) ServeHook(port string) {
	http.HandleFunc("/", s.hook)
	if port != "" {
		if err := http.ListenAndServe(port, nil); err != nil {
			panic(err)
		}
	} else {
		if err := http.ListenAndServe(":8000", nil); err != nil {
			panic(err)
		}
	}
}

func (s *TbBot) hook(w http.ResponseWriter, r *http.Request) {
	// TODO
	// Recieve message from telegram server
}
