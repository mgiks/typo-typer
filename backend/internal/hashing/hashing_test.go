package hashing

import (
	"testing"
)

func TestGenerateSalt(t *testing.T) {
	salt, err := GenerateSalt()

	assertNoError(t, err)
	assertNotEmpty(t, salt)
}

func TestHash(t *testing.T) {
	t.Run("should generate a hash", func(t *testing.T) {
		password := "somePassword123"
		salt := "someSalt"

		hash := Hash(password, salt)

		assertNotEmpty(t, hash)
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

		assertError(t, err)
	})

	t.Run("should return 'true' for correct password", func(t *testing.T) {
		password := "somePassword123"
		salt := "someSalt"

		hashedPass := Hash(password, salt)

		isEqual, err := IsEqualToHash("somePassword123", hashedPass)

		assertNoError(t, err)
		assertTruthiness(t, isEqual)
	})

	t.Run("should return 'false' for incorrect password", func(t *testing.T) {
		password := "somePassword123"
		salt := "someSalt"

		hashedPass := Hash(password, salt)

		isEqual, err := IsEqualToHash("wrongPassword", hashedPass)

		assertNoError(t, err)
		assertFalsiness(t, isEqual)
	})
}

func assertTruthiness(t testing.TB, boolean bool) {
	if !boolean {
		t.Errorf("should return true but doesn't")
	}
}

func assertFalsiness(t testing.TB, boolean bool) {
	if boolean {
		t.Errorf("should return false but doesn't")
	}
}

func assertError(t testing.TB, err error) {
	t.Helper()
	if err == nil {
		t.Errorf("should return an error but doesn't")
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("should run without any errors but doesn't")
	}
}

func assertNotEmpty(t testing.TB, str string) {
	t.Helper()
	if len(str) == 0 {
		t.Errorf("should return non-empty string but doesn't")
	}
}
