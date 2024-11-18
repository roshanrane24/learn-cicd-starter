package auth

import (
	"testing"
	"net/http"
	//"errors"
)

func TestGetAPIKey(t *testing.T) {
	h := http.Header{}

	h1 := http.Header{}
	h1.Add("Authorization", "")

	h2 := http.Header{}
	h2.Add("Authorization", "ApiKey RandomChars")

	// h3 := http.Header{}
	// h3.Add("Authorization", "ApiKey")

	tests := []struct{
		input 		http.Header
		expKey	string
		expErr 		error
	}{
		{h, "", ErrNoAuthHeaderIncluded},
		{h1, "", ErrNoAuthHeaderIncluded},
		{h2, "RandomChars", nil},
		//{h3, "", errors.New("malformed authorization header")},
	}

	for _, test := range tests {
		key, err := GetAPIKey(test.input)
		if key != test.expKey || err != test.expErr {
			t.Fatalf("Inputs -- %v\nExpected:\nKey=%v\nError=%v\n\nGot:\nKey=%v\nErr=%v\n", test.input,  test.expKey, test.expErr, key, err)
		}
	}


}
