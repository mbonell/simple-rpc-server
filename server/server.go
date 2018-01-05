package server

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/mbonell/rpc-server/models"
)

type Service struct {
	Provider string
}

func (s *Service) Execute(rq *models.Request, rp *models.Response) error {

	log.Printf("Exectuting request with data: %+v", rq)
	if rq.Function == "FunctionWithExpectedError" {
		rp.Code = http.StatusInternalServerError
		rp.Error = errors.New("A standard Go error")
	} else {
		rp.Code = http.StatusOK
		result, err := json.Marshal("A valid result")
		if err != nil {
			return err
		}
		rp.Msg = result
	}

	return nil
}
