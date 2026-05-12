package internal

import "testing"

type testFmt struct {
	name           string
	size           int64
	hasHumanFlag   bool
	expectedOutput string
}

var testPath = "~/some.json"
var expectedPathSfx = "\t" + testPath

func TestOutputFmt(t *testing.T) {
	tests := []testFmt{
		{
			name:           "One Bite output fmt",
			size:           1,
			hasHumanFlag:   false,
			expectedOutput: "1B" + expectedPathSfx,
		},
		{
			name:           "One Bite output fmt with human flag",
			size:           1,
			hasHumanFlag:   true,
			expectedOutput: "1.0B" + expectedPathSfx,
		},
		{
			name:           "One KBite output fmt",
			size:           int64(kb),
			hasHumanFlag:   false,
			expectedOutput: "1024B" + expectedPathSfx,
		},
		{
			name:           "One KBite output fmt with human flag",
			size:           int64(kb),
			hasHumanFlag:   true,
			expectedOutput: "1.0KB" + expectedPathSfx,
		},
		{
			name:           "Float MBite output fmt with human flag",
			size:           int64(mb) + int64(kb)*200,
			hasHumanFlag:   true,
			expectedOutput: "1.2MB" + expectedPathSfx,
		},
	}

	for _, to := range tests {
		t.Run(to.name, func(t *testing.T) {
			output := OutputFmt(to.size, testPath, to.hasHumanFlag)
			if output != to.expectedOutput {
				t.Errorf("Failed: expected output `%s`, but got `%s`", to.expectedOutput, output)
			}
		})
	}
}
