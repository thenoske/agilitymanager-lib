package domain

type RunRecord struct {
	DefaultPenaltyPoints     float64 `json:"default_penalty_points"`
	DefaultTimePenaltyPoints float64 `json:"time_penalty_points"`
	DefaultDisPenaltyPoints  int     `json:"dis_penalty_points"`
	StandardTime             int     `json:"standard_time"`
	MaxTime                  int     `json:"max_time"`
	Faults                   int     `json:"faults"`
	Refusals                 int     `json:"refusals"`
	Time                     int     `json:"time"`
	Dis                      bool    `json:"dis"`
	NotRunning               bool    `json:"not_running"`

	// calculated fields
	RunPenaltyPoints   float64 `json:"-"`
	TimePenaltyPoints  float64 `json:"-"`
	TotalPenaltyPoints float64 `json:"-"`
}
