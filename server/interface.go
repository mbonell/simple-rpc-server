package server

import "github.com/mbonell/simple-rpc-server/models"

type RemoteMethods interface {
	Execute(*models.Request, *models.Response) error
}
