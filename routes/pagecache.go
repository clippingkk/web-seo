package routes

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

var pageCache []byte

func refreshPageCache(pagePath string) error {
	c, err := os.ReadFile(pagePath)

	if err != nil {
		return err
	}

	pageCache = c
	return nil
}

func loopRefreshPageCache(pagePath string) {
	for tick := range time.Tick(time.Minute * 5) {
		logrus.Infoln("refresh page cache", tick.Format(time.RFC3339))
		if err := refreshPageCache(pagePath); err != nil {
			logrus.Errorln(err)
		}
	}
}
