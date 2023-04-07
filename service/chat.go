package service

import (
	"bufio"
	"fmt"
	"net/http"
)

type ChatReq struct {
	Prompt  string `json:"prompt"`
	Options struct {
		ParentMessageID string `json:"parentMessageId"`
	} `json:"options"`
}

func ChatHandler(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	rsp, err := http.Post("http://43.157.28.85:3004/chat-process", "application/json", r.Body)
	if err != nil {
		fmt.Printf("%v\n", err)
		w.Write([]byte(err.Error()))
		return
	}

	defer rsp.Body.Close()
	reader := bufio.NewReader(rsp.Body)

	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			fmt.Printf("%v\n", err)
			return
		}

		fmt.Printf("%v", string(line))
		w.Write(line)
		flusher.Flush()
	}
}
