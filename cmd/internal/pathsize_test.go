package internal

import "testing"

type testSize struct {
	name         string
	path         string
	expectedSize int64
	all          bool
	recursive    bool
	hasError     bool
}

func TestGetPathSize(t *testing.T) {
	tests := []testSize{
		{
			name:         "With data file",
			path:         "testdata/data.json",
			expectedSize: 347,
			hasError:     false,
		},
		{
			name:         "Empty file",
			path:         "testdata/empty.json",
			expectedSize: 0,
			hasError:     false,
		},
		{
			name:         "Undefined file",
			path:         "testdata/missing.json",
			expectedSize: 0,
			hasError:     true,
		},
		{
			name:         "Dir with files",
			path:         "testdata/dir1",
			expectedSize: 660,
			hasError:     false,
		},
		{
			name:         "Dir include hidden files",
			path:         "testdata/dir1",
			expectedSize: 872,
			all:          true,
			hasError:     false,
		},
		{
			name:         "Scan dir recursive",
			path:         "testdata",
			expectedSize: 1007,
			recursive:    true,
			hasError:     false,
		},
		{
			name:         "Scan dir recursive include hidden files",
			path:         "testdata",
			expectedSize: 1219,
			all:          true,
			recursive:    true,
			hasError:     false,
		},
		{
			name:         "Missing dir",
			path:         "testdata/missingdir",
			expectedSize: 0,
			hasError:     true,
		},
	}
	for _, tf := range tests {
		t.Run(tf.name, func(t *testing.T) {
			size, err := GetPathSize(tf.path, tf.all, tf.recursive)

			if (err != nil) != tf.hasError {
				t.Errorf("Failed: %v", err)
			}
			if size != tf.expectedSize {
				t.Errorf("Failed: expected size %d, but got %d", tf.expectedSize, size)
			}
		})
	}
}
