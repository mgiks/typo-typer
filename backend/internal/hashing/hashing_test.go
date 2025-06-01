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
	password := "somePass"
	salt := "someSalt"

	t.Run("should generate a hash", func(t *testing.T) {
		hash := Hash(password, salt)

		assertNotEmpty(t, hash)
	})

	t.Run("should be idempotent", func(t *testing.T) {
		got1 := Hash(password, salt)
		got2 := Hash(password, salt)

		assertStringEquality(t, got1, got2)
	})
}

func TestIsEqualToHash(t *testing.T) {
	t.Run("should return error on invalid hash", func(t *testing.T) {
		_, err := IsEqualToHash("", "invalidHash")

		assertError(t, err)
	})

	hashedPass := Hash("correctPassword", "someSalt")

	t.Run("should return 'true' for correct password", func(t *testing.T) {
		isEqual, err := IsEqualToHash("correctPassword", hashedPass)

		assertNoError(t, err)
		assertTruthiness(t, isEqual)
	})

	t.Run("should return 'false' for incorrect password", func(t *testing.T) {
		isEqual, err := IsEqualToHash("wrongPassword", hashedPass)

		assertNoError(t, err)
		assertFalsiness(t, isEqual)
	})
}

func assertStringEquality(t testing.TB, str1, str2 string) {
	if str1 != str2 {
		t.Errorf("should have produced equal strings")
	}
}

func assertTruthiness(t testing.TB, boolean bool) {
	if !boolean {
		t.Errorf("should have returned true")
	}
}

func assertFalsiness(t testing.TB, boolean bool) {
	if boolean {
		t.Errorf("should have returned false")
	}
}

func assertError(t testing.TB, err error) {
	t.Helper()
	if err == nil {
		t.Errorf("should have returned an error")
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("should have returned no errors")
	}
}

func assertNotEmpty(t testing.TB, str string) {
	t.Helper()
	if len(str) == 0 {
		t.Fatalf("should have returned non-empty string")
	}
}
