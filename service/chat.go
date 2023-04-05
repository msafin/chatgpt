package service

import (
	"bytes"
	"io"
	"net/http"
)

func ChatHandler(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("{\"retcode\":0, \"retmsg\":\"ok\"}"))

	data := `{"prompt":"清明节","options":{"parentMessageId":"chatcmpl-71mzBHp4w5gykvKAQtpcYin66SOhB"}}`
	buf := bytes.NewBuffer([]byte(data))

	rsp, err := http.Post("http://43.136.68.168:1002/", "application/json", buf)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	body := rsp.Body
	defer body.Close()
	rspData, err := io.ReadAll(body)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(rspData)
}
