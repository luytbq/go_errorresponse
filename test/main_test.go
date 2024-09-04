package go_errorresponse_test

import (
	"log"
	"strings"
	"testing"

	"github.com/luytbq/go_errorresponse"
)

func Test(t *testing.T) {
	err := go_errorresponse.NewErrorWithName("error-name")

	if strings.Compare(err.Error(), `map[details: name:error-name]`) != 0 {
		t.Error("NewErrorWithName failed")
	}

	err = err.WithDetails("error-details")

	if strings.Compare(err.Error(), `map[details:error-details name:error-name]`) != 0 {
		log.Println(err.Error())
		t.Error("WithDetails failed")
	}
}
