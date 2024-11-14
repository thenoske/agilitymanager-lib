package main

import (
	"C"
	"bytes"
	"encoding/json"
	"fmt"
	agilitymanager_lib "github.com/thenoske/agilitymanager-lib"
	"github.com/thenoske/agilitymanager-lib/domain"

	"unsafe"
)

//export validation
func validation(token *C.char, out *byte, outN int64) *byte {

	outBytes := unsafe.Slice(out, outN)[:0]
	buf := bytes.NewBuffer(outBytes)

	var t string = C.GoString(token)
	claims, err := agilitymanager_lib.VerifyToken(t)
	if err != nil {
		buf.WriteString(fmt.Sprintf("%d", 0))
		buf.WriteByte(0)
		return out
	}

	buf.WriteString(string(claims.ToJson()))
	buf.WriteByte(0) // Null terminator, important!

	return out
}

//export calculatePenaltyPoints
func calculatePenaltyPoints(token *C.char, record *C.char, out *byte, outN int64) *byte {

	var (
		err                error
		runRecord          domain.RunRecord
		runPenaltyPoints   float64
		timePenaltyPoints  float64
		totalPenaltyPoints float64
	)

	// prepare output buffer
	outBytes := unsafe.Slice(out, outN)[:0]
	buf := bytes.NewBuffer(outBytes)

	var t string = C.GoString(token)
	_, err = agilitymanager_lib.VerifyToken(t)
	if err != nil {
		buf.WriteString(err.Error())
		buf.WriteByte(0)
		return out
	}

	// unmarshal run record
	var r string = C.GoString(record)

	if err = json.Unmarshal([]byte(r), &runRecord); err != nil {
		buf.WriteString(err.Error())
		buf.WriteByte(0)
		return out
	}

	// calculate penalty points

	// penalty points for refusals and faults
	runPenaltyPoints += float64(runRecord.Faults) * runRecord.DefaultPenaltyPoints
	runPenaltyPoints += float64(runRecord.Refusals) * runRecord.DefaultPenaltyPoints

	// penalty points for time
	if runRecord.Time > runRecord.StandardTime {
		timePenaltyPoints += float64(runRecord.Time-runRecord.StandardTime) * runRecord.DefaultTimePenaltyPoints
	}

	totalPenaltyPoints = runPenaltyPoints + timePenaltyPoints

	// max time overtake
	if runRecord.Time > runRecord.MaxTime {
		runRecord.Dis = true
		runRecord.Time = 0
		totalPenaltyPoints = float64(runRecord.DefaultDisPenaltyPoints)
	}

	// disqualification
	if runRecord.Dis {
		runPenaltyPoints = float64(runRecord.DefaultDisPenaltyPoints)
		totalPenaltyPoints = float64(runRecord.DefaultDisPenaltyPoints)
	}

	// not running
	if runRecord.NotRunning {
		runPenaltyPoints = float64(runRecord.DefaultDisPenaltyPoints)
		totalPenaltyPoints = float64(runRecord.DefaultDisPenaltyPoints)
	}

	buf.WriteString(fmt.Sprintf(
		`{"run_penalty_points": %f, "time_penalty_points": %f, "total_penalty_points": %f, "dis": %t, "not_running": %t, "time": %d}`,
		runPenaltyPoints,
		timePenaltyPoints,
		totalPenaltyPoints,
		runRecord.Dis,
		runRecord.NotRunning,
		runRecord.Time))
	buf.WriteByte(0) // Null terminator, important!
	return out
}

func main() {}
