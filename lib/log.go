package lib

import (
	"os"
	"runtime"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func Logwrite(r *http.Request) {
	//MemStats Setup
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)

	f, err := os.OpenFile("./log/access.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		//<!>Do not Write log for not open access.log<!>
		if err != nil {
			log.WithFields(log.Fields{
				"URL" : r.URL,
				"Alloc" : ms.Alloc / 1024,
			}).Warn("Can not open ./log/access.log file")
			return
		}
		defer f.Close()

		log.SetOutput(f)

		log.WithFields(log.Fields{
			"URL" : r.URL,
			"Alloc" : ms.Alloc / 1024,
		}).Info("Access has come")
}