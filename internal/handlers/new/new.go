package new

import (
	"encoding/json"
	"net/http"
	"notepad/internal/storage/postgres"
	"strconv"

	"go.uber.org/zap"
)

type Request struct {
	Author  string `json:"author" validate:"required,url"`
	Topic   string `json:"topic"`
	Content string `json:"content"`
}

type Response struct {
	Code int `json:"code"`
}

func New(logger *zap.Logger, db *postgres.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		size, _ := strconv.Atoi(r.Header.Get("Content-Length"))
		var b = make([]byte, size)
		r.Body.Read(b)

		req := Request{}
		err := json.Unmarshal(b, &req)
		if err != nil {
			logger.Error("can't unmarshal json --->", zap.Error(err))

			ans, _ := json.Marshal(Response{Code: 401})
			w.Write(ans)
			return
		}

		err = db.SaveNote(req.Author, req.Topic, req.Content)
		if err != nil {
			logger.Error("can't save data into db --->", zap.Error(err))

			ans, _ := json.Marshal(Response{Code: 401})
			w.Write(ans)
			return
		}

		ans, _ := json.Marshal(Response{Code: 200})
		w.Write(ans)
		logger.Info("note was saved", zap.String("author: ", req.Author), zap.String("topic: ", req.Topic))
	}
}
