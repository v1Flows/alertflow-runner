package payloadendpoints

import (
	"strconv"

	"github.com/AlertFlow/runner/pkg/models"
	"github.com/AlertFlow/runner/pkg/plugin"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func InitPayloadRouter(port int, pluginManager *plugin.Manager, plugins []models.Plugin, payloadEndpoints []models.PayloadEndpoint) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	log.Info("Open Payload Port: ", port)

	payload := router.Group("/payloads")
	for _, endpoint := range payloadEndpoints {
		log.Infof("Open %s Endpoint: %s", endpoint.Name, endpoint.Endpoint)
		payload.POST(endpoint.Endpoint, func(c *gin.Context) {
			for _, p := range plugins {
				if p.Name == endpoint.Name {
					log.Info("Received Payload: ", endpoint.Name)
				}
			}
		})
	}

	router.Run(":" + strconv.Itoa(port))
}
