package jwt_test

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"urlshort.ru/m/api/jwt"
)

// TestExtractToken is a unit test for the jwt.ExtractToken function.
//
// It tests the function's behavior for valid and invalid authorization headers.
// It checks if the extracted token and error match the expected values.
func TestExtractToken(t *testing.T) {
	// Test case: valid authorization header
	authorization := "Bearer abcdef"
	expectedToken := "abcdef"
	expectedError := error(nil)

	token, err := jwt.ExtractToken(authorization)
	if token != expectedToken {
		t.Errorf("Expected token: %s, got: %s", expectedToken, token)
	}
	if err != expectedError {
		t.Errorf("Expected error: %v, got: %v", expectedError, err)
	}

	// Test case: invalid authorization header
	authorization = "Bearer"
	expectedToken = ""
	expectedError = errors.New("invalid token")

	token, err = jwt.ExtractToken(authorization)
	if token != expectedToken {
		t.Errorf("Expected token: %s, got: %s", expectedToken, token)
	}
	if err.Error() != expectedError.Error() {
		t.Errorf("Expected error: %v, got: %v", expectedError, err)
	}
}

// TestGenerateJWT tests the GenerateJWT function.
//
// It verifies that the function generates a JSON Web Token (JWT) correctly based on the given header and payload.
// The expected token is compared with the result token, and an error is raised if they do not match.
// The function is part of the testing package and takes a testing.T object as a parameter.
// It does not return any values.
func TestGenerateJWT(t *testing.T) {
	header := "{\"alg\":\"HS256\",\"typ\":\"JWT\"}"
	payload := "{\"exp\":1692485627,\"user_id\":0}"
	expectedToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTI0ODU2MjcsInVzZXJfaWQiOjB9.3b07e9ac0d3bd88c8c3e00a597215648bcde7f653c6bac6048c50cfa80402ece"

	resultToken := jwt.GenerateJWT(header, payload)

	if resultToken != expectedToken {
		t.Errorf("Unexpected JWT generated. Expected: %s, but got: %s", expectedToken, resultToken)
	}
}

// TestCheckToken is a test function for checking the validity of a JWT token.
//
// The function performs a series of tests to validate different scenarios of token validity. It takes no parameters and does not return any values.
func TestCheckToken(t *testing.T) {
	// Test for an empty token
	t.Run("Empty Token", func(t *testing.T) {
		result := jwt.CheckToken("")
		if result {
			t.Errorf("Expected false, got true")
		}
	})

	// Test for a token with incorrect number of parts
	t.Run("Invalid Token Parts", func(t *testing.T) {
		result := jwt.CheckToken("abc.def")
		if result {
			t.Errorf("Expected false, got true")
		}
	})

	// Test for a token with empty header, payload, or signature
	t.Run("Empty Header, Payload, or Signature", func(t *testing.T) {
		result := jwt.CheckToken("...abc")
		if result {
			t.Errorf("Expected false, got true")
		}
	})

	// Test for a token with invalid signature
	t.Run("Invalid Signature", func(t *testing.T) {

		result := jwt.CheckToken("abc.def.ghi")
		if result {
			t.Errorf("Expected false, got true")
		}
	})

	// Test for a token with invalid decoded header
	t.Run("Invalid Decoded Header", func(t *testing.T) {
		result := jwt.CheckToken("abc.def.ghi")
		if result {
			t.Errorf("Expected false, got true")
		}
	})

	// Test for a token with invalid decoded payload
	t.Run("Invalid Decoded Payload", func(t *testing.T) {

		result := jwt.CheckToken("abc.def.ghi")
		if result {
			t.Errorf("Expected false, got true")
		}
	})

	// Test for a valid token
	t.Run("Valid Token", func(t *testing.T) {

		exp := time.Now().Add(time.Hour * 24).Unix()

		newToken := jwt.GenerateJWT("{\"alg\":\"HS256\",\"typ\":\"JWT\"}", fmt.Sprintf("{\"exp\":%d,\"user_id\":0}", exp))

		result := jwt.CheckToken(newToken)
		if !result {
			t.Errorf("Expected true, got false")
		}
	})
}

// TestGetPayloadRefresh tests the GetPayloadRefresh function.
//
// This function tests the GetPayloadRefresh function by providing different test cases
// and checking if the returned payload and error match the expected values.
//
// Parameters:
// - t: The testing.T object for running tests and reporting failures.
//
// Returns:
// This function does not return anything.
func TestGetPayloadRefresh(t *testing.T) {
	// Test case 1: Valid payload
	payload1 := `{"user_id": 0, "exp": 1692485627}`
	expected1 := jwt.PayloadJWTRefresh{UserID: 0, EXP: 1692485627}
	result1, err1 := jwt.GetPayloadRefresh(payload1)
	if err1 != nil {
		t.Errorf("Unexpected error: %v", err1)
	}
	if result1 != expected1 {
		t.Errorf("Expected %v, but got %v", expected1, result1)
	}

	// Test case 2: Invalid payload
	payload2 := `{"name": "John", "age": "thirty", "foo": "bar"}`
	expected2 := jwt.PayloadJWTRefresh{}
	result2, err2 := jwt.GetPayloadRefresh(payload2)
	if err2 == nil || result2 != expected2 {
		t.Errorf("Expected error, but got %v, %v", err2, result2)
	}

	// Test case 3: Empty payload
	payload3 := `{}`
	expected3 := jwt.PayloadJWTRefresh{}
	result3, err3 := jwt.GetPayloadRefresh(payload3)
	if err3 != nil {
		t.Errorf("Unexpected error: %v", err3)
	}
	if result3 != expected3 {
		t.Errorf("Expected %v, but got %v", expected3, result3)
	}
}
