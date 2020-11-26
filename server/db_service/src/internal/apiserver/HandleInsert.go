package apiserver

import (
	"encoding/json"
	"github.com/kletskovg/typecode/server/db_service/src/internal/db"
	"github.com/kletskovg/typecode/server/db_service/src/internal/utils"
	"net/http"
	// "github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	// "context"
	// "go.mongodb.org/mongo-driver/bson"
)


func (server *APIServer) HandleInsert () http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		utils.EnableCors(&w)

		var file db.File
		err := json.NewDecoder(r.Body).Decode(&file)

		if err != nil {
			defer log.WithFields(log.Fields{
				"file": r.Body,
				"err": err,
			}).Error("During parse request body")
		}

		go db.InsertFile(server.DbClient, file)

		log.Info(file)
		fileMarshal, _ := json.Marshal(file)
		w.Write(fileMarshal)
	}
}