package app

import (
	appConfig "VanO/scores_server/config"
	"os"
)

// WebsocketD структура приложения websocketd
type WebsocketD struct {
	config       *appConfig.Config
	logErrorFile *os.File
	logInfoFile  *os.File
}
