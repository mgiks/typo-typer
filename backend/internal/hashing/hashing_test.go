package hashing

import "testing"

func TestGenerateSalt(t *testing.T) {
	salt, err := GenerateSalt()

	if err != nil {
		t.Errorf("should run without an error but doesn't")
	}

	if len(salt) == 0 {
		t.Fatalf("should return salt but returns empty string")
	}
}

func TestHash(t *testing.T) {
	t.Run("should generate a hash", func(t *testing.T) {
		password := "somePassword123"
		salt := "someSalt"

		got := Hash(password, salt)

		if len(got) == 0 {
			t.Fatalf("should return hash but returns empty string")
		}
	})

	t.Run("should be idempotent", func(t *testing.T) {
		password := "somePassword123"
		salt := "someSalt"

		got1 := Hash(password, salt)
		got2 := Hash(password, salt)

		if got1 != got2 {
			t.Errorf("should produce equal outputs but doesn't")
		}
	})
}

func TestIsEqualToHash(t *testing.T) {
	t.Run("should return error on invalid hash", func(t *testing.T) {
		_, err := IsEqualToHash("somePassword123", "invalidHash")

		if err == nil {
			t.Errorf("should return an error but doesn't")
		}
	})

	t.Run("should return 'true' for correct password", func(t *testing.T) {
		password := "somePassword123"
		salt := "someSalt"

		hashedPass := Hash(password, salt)

		isEqual, err := IsEqualToHash("somePassword123", hashedPass)

		if err != nil {
			t.Errorf("should run without any errors but doesn't")
		}

		if !isEqual {
			t.Errorf("should return 'true' but doesn't")
		}
	})

	t.Run("should return 'false' for incorrect password", func(t *testing.T) {
		password := "somePassword123"
		salt := "someSalt"

		hashedPass := Hash(password, salt)

		isEqual, err := IsEqualToHash("wrongPassword", hashedPass)

		if err != nil {
			t.Errorf("should run without any errors but doesn't")
		}

		if isEqual {
			t.Errorf("should return 'false' but doesn't'")
		}
	})
}
