package service

import "net/http"

func ChatHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("{\"retcode\":0, \"retmsg\":\"ok\"}"))
}
