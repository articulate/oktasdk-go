package okta

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

var testTrustedOrigin *TrustedOrigin

func setupTestTrustedOrigin() {
	scopeTypes := []string{"CORS", "Redirect"}

	var scopes []map[string]string

	for _, scopeType := range scopeTypes {
		scopes = append(scopes, map[string]string{"Type": scopeType})
	}

	testTrustedOrigin = &TrustedOrigin{
		Origin: "http://testing.com",
		Name:   "Testing",
		Scopes: scopes,
	}
}

func TestGetTrustedOrigin(t *testing.T) {
	setup()
	defer teardown()
	setupTestTrustedOrigin()
	testTrustedOrigin.ID = "ow1y4s2cpS59f1xs2p7"

	temp, err := json.Marshal(testTrustedOrigin)
	if err != nil {
		t.Errorf("TrustedOrigins.GetTrustedOrigin json Marshall returned error: %v", err)
	}

	trustedOriginTestJSONString := string(temp)

	mux.HandleFunc("/trustedOrigins/ow1y4s2cpS59f1xs2p7", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testAuthHeader(t, r)
		fmt.Fprint(w, trustedOriginTestJSONString)
	})

	outputTrustedOrigin, _, err := client.TrustedOrigins.GetTrustedOrigin("ow1y4s2cpS59f1xs2p7")
	if err != nil {
		t.Errorf("TrustedOrigins.GetTrustedOrigin returned error: %v", err)
	}
	if !reflect.DeepEqual(outputTrustedOrigin, testTrustedOrigin) {
		t.Errorf("client.TrustedOrigins.GetTrustedOrigin returned \n\t%+v, want \n\t%+v\n", outputTrustedOrigin, testTrustedOrigin)
	}
}
