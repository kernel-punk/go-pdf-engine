package examples

import (
	"math/rand"
	"sync/atomic"
	"time"
)

var randomSeed uint64 = uint64(time.Now().UnixNano())

var webServerStates = []string{
	"WORK", "ERROR",
}

var updateStatuses = []string{
	"Required", "Not Required",
}

var operatingSystems = []string{
	"Debian", "Ubuntu", "Red Hat", "Cent OS",
}

func MultipleRandomTests(iteration int) []*ServerTestData {
	rng := newRandomSource()
	counter := make([]*ServerTestData, 0, iteration)

	for i := 0; i < iteration; i++ {
		test := singleRandomTest(rng)
		test.Server = defaultServerName(i)
		counter = append(counter, test)
	}

	return counter
}

func singleRandomTest(rng *rand.Rand) *ServerTestData {
	testData := &ServerTestData{
		LinkUp: rng.Intn(2) == 0,
	}

	if testData.LinkUp {
		ping := rng.Intn(1001)
		ssdUsed := rng.Intn(101)
		ramUsed := rng.Intn(101)
		uptime := randomUptime(rng)

		testData.PingMS = &ping
		testData.SSDUsedPercent = &ssdUsed
		testData.RAMUsedPercent = &ramUsed
		testData.NeedUpdate = updateStatuses[rng.Intn(len(updateStatuses))]
		testData.WebServerState = webServerStates[rng.Intn(len(webServerStates))]
		testData.OperatingSystem = operatingSystems[rng.Intn(len(operatingSystems))]
		testData.Uptime = &uptime
	}

	return testData
}

func newRandomSource() *rand.Rand {
	seed := atomic.AddUint64(&randomSeed, 1)

	return rand.New(rand.NewSource(int64(seed)))
}
