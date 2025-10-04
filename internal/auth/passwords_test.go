package auth_test

import (
	"github.com/roxensox/dailychapter/internal/auth"
	"testing"
)

func TestPWHashing(t *testing.T) {
	// Creates hash of "test" ahead of time
	testHash, err := auth.HashPassword("test")
	if err != nil {
		t.Errorf("Failed to hash test password during startup")
	}

	test_cases := []struct {
		Pass1    string
		Pass2    string
		Expected bool
	}{
		// Matching passwords
		{
			Pass1:    "test",
			Pass2:    "test",
			Expected: true,
		},
		// Slight password mismatch
		{
			Pass1:    "test",
			Pass2:    "Test",
			Expected: false,
		},
		// Second password is a hash of the same string as pass1
		{
			Pass1:    "test",
			Pass2:    testHash,
			Expected: false,
		},
	}

	for _, c := range test_cases {
		// Hashes pass1
		hash1, err := auth.HashPassword(c.Pass1)
		if err != nil {
			t.Errorf("Failed to hash password: %s", c.Pass1)
		}

		// Checks passwords using hash checker
		match, err := auth.CheckPasswordHash(c.Pass2, hash1)
		if err != nil {
			t.Errorf("Failed to check password: %s", c.Pass1)
		}

		// Checks if result is as expected
		if match != c.Expected {
			t.Errorf("%s == %s should be %v", c.Pass1, c.Pass2, c.Expected)
		}
	}
}
