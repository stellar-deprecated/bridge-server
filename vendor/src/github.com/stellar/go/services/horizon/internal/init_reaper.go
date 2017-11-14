package horizon

import (
	"github.com/stellar/go/services/horizon/internal/reap"
)

func initReaper(app *App) {
	app.reaper = reap.New(app.config.HistoryRetentionCount, app.HorizonSession(nil))
}

func init() {
	appInit.Add("reaper", initReaper, "app-context", "log", "horizon-db")
}
