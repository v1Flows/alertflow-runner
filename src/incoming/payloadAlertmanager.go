package incoming

import (
	"alertflow-runner/src/config"
	"alertflow-runner/src/models"
	"alertflow-runner/src/outgoing/payload"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Receiver struct {
	Receiver string `json:"receiver"`
}

func AlertmanagerPayloadHandler(context *gin.Context) {
	log.Info("Received Alertmanager Payload")
	incPayload, err := io.ReadAll(context.Request.Body)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read request body",
		})
		return
	}

	// get receiver from payload
	receiver := Receiver{}
	json.Unmarshal(incPayload, &receiver)

	payloadData := models.Payload{
		Payload:  incPayload,
		FlowID:   receiver.Receiver,
		RunnerID: config.Config.RunnerID,
	}

	payload.SendPayload(payloadData)
}
