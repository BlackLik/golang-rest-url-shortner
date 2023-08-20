package utils_test

import (
	"testing"

	"urlshort.ru/m/utils"
)

// TestGenerateShortHash is a unit test function that tests the GenerateShortHashMD5 function in the utils package.
//
// It tests the function with different inputs and verifies that the generated hash matches the expected hash.
// The function takes a string input and returns a string hash using the MD5 algorithm.
func TestGenerateShortHash(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "Empty string",
			input: "",
			want:  "d41d8cd98f00b204e9800998ecf8427e",
		},
		{
			name:  "Short string",
			input: "hello",
			want:  "5d41402abc4b2a76b9719d911017c592",
		},
		{
			name:  "Long string",
			input: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
			want:  "818c6e601a24f72750da0f6c9b8ebe28",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := utils.GenerateShortHashMD5(tt.input)
			if got != tt.want {
				t.Errorf("GenerateShortHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestGenerateShortHashSHA256 tests the GenerateShortHashSHA256 function.
//
// It tests the function with various input cases including empty input,
// non-empty input, and input with special characters. It compares the
// output of the function with the expected output and reports any
// failures using the testing.T.Errorf function.
func TestGenerateShortHashSHA256(t *testing.T) {
	// Test case 1: Empty input
	input1 := ""
	expectedOutput1 := "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"
	output1 := utils.GenerateShortHashSHA256(input1)
	if output1 != expectedOutput1 {
		t.Errorf("Test case 1 failed: expected %s, got %s", expectedOutput1, output1)
	}

	// Test case 2: Non-empty input
	input2 := "Hello, world!"
	expectedOutput2 := "315f5bdb76d078c43b8ac0064e4a0164612b1fce77c869345bfc94c75894edd3"
	output2 := utils.GenerateShortHashSHA256(input2)
	if output2 != expectedOutput2 {
		t.Errorf("Test case 2 failed: expected %s, got %s", expectedOutput2, output2)
	}

	// Test case 3: Input with special characters
	input3 := "!@#$%^&*()_+"
	expectedOutput3 := "36d3e1bc65f8b67935ae60f542abef3e55c5bbbd547854966400cc4f022566cb"
	output3 := utils.GenerateShortHashSHA256(input3)
	if output3 != expectedOutput3 {
		t.Errorf("Test case 3 failed: expected %s, got %s", expectedOutput3, output3)
	}
}

// TestBase64Decode is a test function for the Base64Decode function.
//
// It tests the decoding of a valid base64 encoded byte array, an invalid base64 encoded byte array,
// and an empty byte array.
// It checks the decoded result against the expected value and ensures that the proper errors are returned.
func TestBase64Decode(t *testing.T) {
	// Test case 1: Decoding a valid base64 encoded byte array
	input1 := []byte("SGVsbG8gd29ybGQ=")
	expected1 := "Hello world"
	result1, err := utils.Base64Decode(input1)
	if err != nil {
		t.Errorf("Error decoding base64: %s", err.Error())
	}
	if result1 != expected1 {
		t.Errorf("Expected: %s, but got: %s", expected1, result1)
	}

	// Test case 2: Decoding an invalid base64 encoded byte array
	input2 := []byte("SGVsbG8gd29ybGQ==")
	_, err2 := utils.Base64Decode(input2)
	if err2 == nil {
		t.Error("Expected an error while decoding invalid base64, but got nil")
	}

	// Test case 3: Decoding an empty byte array
	input3 := []byte("")
	expected3 := ""
	result3, err3 := utils.Base64Decode(input3)
	if err3 != nil {
		t.Errorf("Error decoding base64: %s", err3.Error())
	}
	if result3 != expected3 {
		t.Errorf("Expected: %s, but got: %s", expected3, result3)
	}
}

// TestBase64Encode is a unit test for the function that encodes a given string to Base64.
//
// It tests different input strings and checks if the actual result matches the expected result.
// The test cases include a non-empty string, a string with numbers, and an empty string.
// The expected results are precomputed Base64 encodings of the input strings.
// If the actual result does not match the expected result for any test case, an error is reported.
func TestBase64Encode(t *testing.T) {
	testCases := []struct {
		input    string
		expected []byte
	}{
		{
			input:    "Hello, World!",
			expected: []byte("SGVsbG8sIFdvcmxkIQ=="),
		},
		{
			input:    "1234567890",
			expected: []byte("MTIzNDU2Nzg5MA=="),
		},
		{
			input:    "",
			expected: []byte(""),
		},
	}

	for _, testCase := range testCases {
		actual := utils.Base64Encode(testCase.input)
		if string(actual) != string(testCase.expected) {
			t.Errorf("Input: %s\nExpected: %s\nActual: %s\n", testCase.input, testCase.expected, actual)
		}
	}
}

// TestConver10IntTo32String tests the Conver10IntTo32String function.
//
// It verifies the correctness of converting a given base-10 integer to a base-32 string.
// The function takes in a list of test cases, each containing a name, input integer, and expected output string.
// For each test case, the function runs a sub-test with the name and performs the conversion.
// It compares the result with the expected output and raises an error if they don't match.
// The function is designed to be used with the testing package and is intended for unit testing.
func TestConver10IntTo32String(t *testing.T) {
	tests := []struct {
		name  string
		input int64
		want  string
	}{
		{
			name:  "Positive integer",
			input: 12345,
			want:  "c1p",
		},
		{
			name:  "Negative integer",
			input: -98765,
			want:  "-30ed",
		},
		{
			name:  "Zero",
			input: 0,
			want:  "0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := utils.Conver10IntTo32String(tt.input)
			if got != tt.want {
				t.Errorf("Conver10IntTo32String() = %v, want %v", got, tt.want)
			}
		})
	}
}
