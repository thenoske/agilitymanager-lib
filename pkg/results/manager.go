package results

import "C"
import (
	"github.com/thenoske/agilitymanager-lib/domain"
)

// Manager interface for results manager
type Manager interface {
	// CalculatePenaltyPoints calculate penalty points for run record
	CalculatePenaltyPoints(*domain.RunRecord) (float64, float64, float64, error)
}

// NewManager creates a new results manager
func NewManager() Manager {
	return &manager{}
}

// manager implements Manager interface
type manager struct {
}

// CalculatePenaltyPoints calculate penalty points for run record
func (m *manager) CalculatePenaltyPoints(record *domain.RunRecord) (float64, float64, float64, error) {

	var (
		runPenaltyPoints   float64
		timePenaltyPoints  float64
		totalPenaltyPoints float64
	)

	// penalty points for refusals and faults
	runPenaltyPoints += float64(record.Faults) * record.DefaultPenaltyPoints
	runPenaltyPoints += float64(record.Refusals) * record.DefaultPenaltyPoints

	// penalty points for time
	if record.Time > record.StandardTime {
		timePenaltyPoints += float64(record.Time-record.StandardTime) * record.DefaultTimePenaltyPoints
	}

	totalPenaltyPoints = runPenaltyPoints + timePenaltyPoints

	// three refusals or max time overtake is disqualification
	if record.Refusals >= 3 || record.Time > record.MaxTime {
		record.Dis = true
	}

	// disqualification
	if record.Dis {
		record.Time = 0
		timePenaltyPoints = 0
		runPenaltyPoints = float64(record.DefaultDisPenaltyPoints)
		totalPenaltyPoints = float64(record.DefaultDisPenaltyPoints)
	}

	// not running
	if record.NotRunning {
		record.Time = 0
		runPenaltyPoints = float64(record.DefaultDisPenaltyPoints)
		totalPenaltyPoints = float64(record.DefaultDisPenaltyPoints)
	}

	return runPenaltyPoints, timePenaltyPoints, totalPenaltyPoints, nil
}
