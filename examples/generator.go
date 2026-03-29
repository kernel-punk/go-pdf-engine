package examples

import (
	"math/rand"
	"strconv"
	"time"
)

func MultipleRandomTests(iteration int) []*ServerTestData {

	var counter []*ServerTestData

	for i := 0; i < iteration; i++ {
		counter = append(counter, SingleRandomTest())
	}

	return counter
}

func SingleRandomTest() *ServerTestData {

	rand.New(rand.NewSource(time.Now().UnixNano()))

	var testData ServerTestData

	var a = []string{
		"WORK", "ERROR",
	}

	var e = []string{
		"YES", "NO",
	}

	var b = []string{
		"Required", "Not Required",
	}

	var h = []string{
		"Debian", "Ubuntu", "Red Hat", "Cent OS",
	}

	testData.E = e[rand.Intn(len(e))]

	if testData.E == "YES" {
		testData.F = strconv.Itoa(rand.Intn(1001))
		testData.G = strconv.Itoa(rand.Intn(101))
		testData.H = strconv.Itoa(rand.Intn(101))
		testData.A = b[rand.Intn(len(b))]
		testData.B = a[rand.Intn(len(a))]
		testData.C = h[rand.Intn(len(h))]
		testData.D = randomUptime()
	} else {
		testData.F = "N/A"
		testData.G = "N/A"
		testData.H = "N/A"
		testData.A = "N/A"
		testData.B = "N/A"
		testData.C = "N/A"
		testData.D = "N/A"
	}

	return &testData
}
