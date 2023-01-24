package collector

import (
	"bufio"
	"encoding/json"
	"time"
)

type TestEvent struct {
	TestName string     `json:"Test"`
	Time     time.Time  `json:"Time"` // encodes as an RFC3339-format string
	Action   TestAction `json:"Action"`
	Package  string     `json:"Package"`
	Elapsed  float64    `json:"Elapsed"` // seconds
	Output   string     `json:"Output"`
}

type TestEventForView struct {
	TestName string
	Package  string
	Elapsed  float64
	Outputs  []string
	Done     bool
	State    TestAction
}

type TestAction string

const (
	RUN    TestAction = "run"
	PAUSE  TestAction = "pause"
	CONT   TestAction = "cout"
	PASS   TestAction = "pass"
	BENCH  TestAction = "bench"
	FAIL   TestAction = "fail"
	OUTPUT TestAction = "output"
	SKIP   TestAction = "skip"
)

func (ta TestAction) IsFinished() bool {
	return ta == PASS || ta == FAIL || ta == SKIP
}

type Results struct {
	Pass int64
	Fail int64
	Skip int64
}

func NewTestEventForView(te *TestEvent) *TestEventForView {
	return &TestEventForView{
		TestName: te.TestName,
		Package:  te.Package,
		Elapsed:  te.Elapsed,
		Done:     false,
	}
}

func (r *Results) CountTestResults(testAction TestAction) {
	switch testAction {
	case PASS:
		r.Pass++
	case FAIL:
		r.Fail++
	case SKIP:
		r.Skip++
	default:
	}
}

func (r *Results) Total() int64 {
	return r.Pass + r.Fail + r.Skip
}

func UnmarshalTestEvent(b []byte) (TestEvent, error) {
	var te TestEvent
	err := json.Unmarshal(b, &te)
	if err != nil {
		return TestEvent{}, err
	}
	return te, nil
}

func ReadLogFile() {

}

func ReadLogStdout(stdinScanner *bufio.Scanner) ([]*TestEventForView, *Results, error) {
	results := Results{}
	testMap := map[string]*TestEventForView{}
	tests := []*TestEventForView{}

	for stdinScanner.Scan() {
		line := stdinScanner.Bytes()
		te, err := UnmarshalTestEvent(line)
		if err != nil {
			return nil, nil, err
		}
		if te.TestName != "" {
			tev, ok := testMap[te.TestName]
			if ok {
				if te.Action == OUTPUT {
					tev.Outputs = append(tev.Outputs, te.Output)
				}
				if !tev.Done && te.Action.IsFinished() {
					tev.Elapsed = te.Elapsed
					tev.State = te.Action
					tev.Done = true
					tests = append(tests, tev)
				}
			} else {
				testMap[te.TestName] = NewTestEventForView(&te)
			}
			results.CountTestResults(te.Action)
		}
	}
	return tests, &results, nil
}
