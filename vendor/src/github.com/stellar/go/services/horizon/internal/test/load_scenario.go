package test

import (
	"github.com/stellar/go/services/horizon/internal/test/scenarios"
)

func loadScenario(scenarioName string, includeHorizon bool) {
	stellarCorePath := scenarioName + "-core.sql"
	horizonPath := scenarioName + "-horizon.sql"

	if !includeHorizon {
		horizonPath = "blank-horizon.sql"
	}

	scenarios.Load(StellarCoreDatabaseURL(), stellarCorePath)
	scenarios.Load(DatabaseURL(), horizonPath)
}
