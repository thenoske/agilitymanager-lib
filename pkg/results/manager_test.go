package results

import (
	"github.com/thenoske/agilitymanager-lib/domain"
	"testing"
)

func Test_manager_CalculatePenaltyPoints(t *testing.T) {
	type args struct {
		record domain.RunRecord
	}
	tests := []struct {
		name        string
		args        args
		runPoints   float64
		timePoints  float64
		totalPoints float64
		dis         bool
		notRunning  bool
		wantErr     bool
	}{
		{
			name: "Clear run",
			args: args{
				record: domain.RunRecord{
					DefaultPenaltyPoints:     5,
					DefaultTimePenaltyPoints: .01,
					DefaultDisPenaltyPoints:  100,
					StandardTime:             4000,
					MaxTime:                  5000,
					Faults:                   0,
					Refusals:                 0,
					Time:                     3000,
					Dis:                      false,
					NotRunning:               false,
				},
			},
			runPoints:   0,
			timePoints:  0,
			totalPoints: 0,
			dis:         false,
			notRunning:  false,
			wantErr:     false,
		},
		{
			name: "Run with faults",
			args: args{
				record: domain.RunRecord{
					DefaultPenaltyPoints:     5,
					DefaultTimePenaltyPoints: .01,
					DefaultDisPenaltyPoints:  100,
					StandardTime:             4000,
					MaxTime:                  5000,
					Faults:                   1,
					Refusals:                 0,
					Time:                     3000,
					Dis:                      false,
					NotRunning:               false,
				},
			},
			runPoints:   5,
			timePoints:  0,
			totalPoints: 5,
			dis:         false,
			notRunning:  false,
			wantErr:     false,
		},
		{
			name: "Run with refusals",
			args: args{
				record: domain.RunRecord{
					DefaultPenaltyPoints:     5,
					DefaultTimePenaltyPoints: .01,
					DefaultDisPenaltyPoints:  100,
					StandardTime:             4000,
					MaxTime:                  5000,
					Faults:                   0,
					Refusals:                 1,
					Time:                     3000,
					Dis:                      false,
					NotRunning:               false,
				},
			},
			runPoints:   5,
			timePoints:  0,
			totalPoints: 5,
			dis:         false,
			notRunning:  false,
			wantErr:     false,
		},
		{
			name: "Run with faults and refusals",
			args: args{
				record: domain.RunRecord{
					DefaultPenaltyPoints:     5,
					DefaultTimePenaltyPoints: .01,
					DefaultDisPenaltyPoints:  100,
					StandardTime:             4000,
					MaxTime:                  5000,
					Faults:                   1,
					Refusals:                 1,
					Time:                     3000,
					Dis:                      false,
					NotRunning:               false,
				},
			},
			runPoints:   10,
			timePoints:  0,
			totalPoints: 10,
			dis:         false,
			notRunning:  false,
			wantErr:     false,
		},
		{
			name: "Run with three refusals",
			args: args{
				record: domain.RunRecord{
					DefaultPenaltyPoints:     5,
					DefaultTimePenaltyPoints: .01,
					DefaultDisPenaltyPoints:  100,
					StandardTime:             4000,
					MaxTime:                  5000,
					Faults:                   0,
					Refusals:                 3,
					Time:                     3000,
					Dis:                      false,
					NotRunning:               false,
				},
			},
			runPoints:   100,
			timePoints:  0,
			totalPoints: 100,
			dis:         true,
			notRunning:  false,
			wantErr:     false,
		},
		{
			name: "Run with four refusals",
			args: args{
				record: domain.RunRecord{
					DefaultPenaltyPoints:     5,
					DefaultTimePenaltyPoints: .01,
					DefaultDisPenaltyPoints:  100,
					StandardTime:             4000,
					MaxTime:                  5000,
					Faults:                   0,
					Refusals:                 4,
					Time:                     3000,
					Dis:                      false,
					NotRunning:               false,
				},
			},
			runPoints:   100,
			timePoints:  0,
			totalPoints: 100,
			dis:         true,
			notRunning:  false,
			wantErr:     false,
		},
		{
			name: "Run with disqualification",
			args: args{
				record: domain.RunRecord{
					DefaultPenaltyPoints:     5,
					DefaultTimePenaltyPoints: .01,
					DefaultDisPenaltyPoints:  100,
					StandardTime:             4000,
					MaxTime:                  5000,
					Faults:                   0,
					Refusals:                 0,
					Time:                     0,
					Dis:                      true,
					NotRunning:               false,
				},
			},
			runPoints:   100,
			timePoints:  0,
			totalPoints: 100,
			dis:         true,
			notRunning:  false,
			wantErr:     false,
		},
		{
			name: "Run with not running",
			args: args{
				record: domain.RunRecord{
					DefaultPenaltyPoints:     5,
					DefaultTimePenaltyPoints: .01,
					DefaultDisPenaltyPoints:  100,
					StandardTime:             4000,
					MaxTime:                  5000,
					Faults:                   0,
					Refusals:                 0,
					Time:                     0,
					Dis:                      false,
					NotRunning:               true,
				},
			},
			runPoints:   100,
			timePoints:  0,
			totalPoints: 100,
			dis:         false,
			notRunning:  true,
			wantErr:     false,
		},
		{
			name: "Run with time over standard time",
			args: args{
				record: domain.RunRecord{
					DefaultPenaltyPoints:     5,
					DefaultTimePenaltyPoints: .01,
					DefaultDisPenaltyPoints:  100,
					StandardTime:             4000,
					MaxTime:                  5000,
					Faults:                   0,
					Refusals:                 0,
					Time:                     4200,
					Dis:                      false,
					NotRunning:               false,
				},
			},
			runPoints:   0,
			timePoints:  2,
			totalPoints: 2,
			dis:         false,
			notRunning:  false,
			wantErr:     false,
		},
		{
			name: "Run with fault and time over standard time",
			args: args{
				record: domain.RunRecord{
					DefaultPenaltyPoints:     5,
					DefaultTimePenaltyPoints: .01,
					DefaultDisPenaltyPoints:  100,
					StandardTime:             4000,
					MaxTime:                  5000,
					Faults:                   1,
					Refusals:                 0,
					Time:                     4200,
					Dis:                      false,
					NotRunning:               false,
				},
			},
			runPoints:   5,
			timePoints:  2,
			totalPoints: 7,
			dis:         false,
			notRunning:  false,
			wantErr:     false,
		},
		{
			name: "Run with time over max time",
			args: args{
				record: domain.RunRecord{
					DefaultPenaltyPoints:     5,
					DefaultTimePenaltyPoints: .01,
					DefaultDisPenaltyPoints:  100,
					StandardTime:             4000,
					MaxTime:                  5000,
					Faults:                   0,
					Refusals:                 0,
					Time:                     5200,
					Dis:                      false,
					NotRunning:               false,
				},
			},
			runPoints:   100,
			timePoints:  0,
			totalPoints: 100,
			dis:         true,
			notRunning:  false,
			wantErr:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &manager{}
			got, got1, got2, err := m.CalculatePenaltyPoints(&tt.args.record)

			if (err != nil) != tt.wantErr {
				t.Errorf("CalculatePenaltyPoints() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.runPoints {
				t.Errorf("CalculatePenaltyPoints() runPoints = %v, want %v", got, tt.runPoints)
			}
			if got1 != tt.timePoints {
				t.Errorf("CalculatePenaltyPoints() timePoints = %v, want %v", got1, tt.timePoints)
			}
			if got2 != tt.totalPoints {
				t.Errorf("CalculatePenaltyPoints() totalPoints = %v, want %v", got2, tt.totalPoints)
			}
			if tt.args.record.Dis != tt.dis {
				t.Errorf("CalculatePenaltyPoints() dis = %v, want %v", tt.args.record.Dis, tt.dis)
			}
			if tt.args.record.NotRunning != tt.notRunning {
				t.Errorf("CalculatePenaltyPoints() notRunning = %v, want %v", tt.args.record.NotRunning, tt.notRunning)
			}
		})
	}
}
