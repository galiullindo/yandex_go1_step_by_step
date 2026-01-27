package traffic

import "testing"

func TestGetTrafficLevel(t *testing.T) {
	lat, lng := 1.0, 2.0 // :)
	expected := 3        // :)

	trafficClient := RealTrafficClient{}
	if got := trafficClient.GetTrafficLevel(lat, lng); got != expected {
		t.Errorf("RealTrafficClient.GetTrafficLevel(%0.1f, %0.1f) got %d, expected %d\n", lat, lng, got, expected)
	}

}

func TestGetTrafficMultiplier(t *testing.T) {
	var tests = []struct {
		name     string
		level    int
		expected float64
	}{
		{"Standart", 1, 1.0},
		{"Standart", 2, 1.1},
		{"Standart", 3, 1.2},
		{"Standart", 4, 1.3},
		{"Standart", 5, 1.4},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := GetTrafficMultiplier(test.level); got != test.expected {
				t.Errorf("GetTrafficMultiplier(%d) got %0.1f, expected %0.1f\n", test.level, got, test.expected)
			}
		})
	}
}
