package examples

import "testing"

func TestPingColorUsesMetricSpecificThresholds(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		value *int
		want  RGB
	}{
		{name: "healthy ping", value: intPtr(45), want: PureGreen},
		{name: "warning ping", value: intPtr(140), want: Amber},
		{name: "critical ping", value: intPtr(650), want: Crimson},
		{name: "missing ping", value: nil, want: ColorDefault},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			if got := PingColor(tc.value); got != tc.want {
				t.Fatalf("expected %+v, got %+v", tc.want, got)
			}
		})
	}
}

func intPtr(value int) *int {
	return &value
}
