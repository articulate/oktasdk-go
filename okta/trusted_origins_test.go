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

func TestTrustedOriginCreate(t *testing.T) {
	setup()
	defer teardown()
	setupTestTrustedOrigin()

	temp, err := json.Marshal(testTrustedOrigin)

	if err != nil {
		t.Errorf("TrustedOrigins.CreateTrustedOrigin json Marshall returned error: %v", err)
	}

	TrustedOriginTestJSONString := string(temp)

	mux.HandleFunc("/trustedOrigins", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testAuthHeader(t, r)
		fmt.Fprint(w, TrustedOriginTestJSONString)
	})

	outputTrustedOrigin, _, err := client.TrustedOrigins.CreateTrustedOrigin(testTrustedOrigin)
	if err != nil {
		t.Errorf("client.TrustedOrigins.CreateTrustedOrigin returned error: %v", err)
	}
	if !reflect.DeepEqual(outputTrustedOrigin, testTrustedOrigin) {
		t.Errorf("client.TrustedOrigins.CreateTrustedOrigin returned \n\t%+v, want \n\t%+v\n", outputTrustedOrigin, testTrustedOrigin)
	}
}

func TestTrustedOriginUpdate(t *testing.T) {
	setup()
	defer teardown()
	setupTestTrustedOrigin()
	testTrustedOrigin.ID = "ow1y4s2cpS59f1xs2p7"

	testTrustedOrigin.Name = "Testing Update"
	temp, err := json.Marshal(testTrustedOrigin)
	if err != nil {
		t.Errorf("TrustedOrigins.UpdateTrustedOrigin json Marshall returned error: %v", err)
	}
	updateTestJSONString := string(temp)

	mux.HandleFunc("/trustedOrigins/ow1y4s2cpS59f1xs2p7", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		testAuthHeader(t, r)
		fmt.Fprint(w, updateTestJSONString)
	})

	outputTrustedOrigin, _, err := client.TrustedOrigins.UpdateTrustedOrigin("ow1y4s2cpS59f1xs2p7", testTrustedOrigin)
	if err != nil {
		t.Errorf("client.TrustedOrigins.UpdateTrustedOrigin returned error: %v", err)
	}
	if !reflect.DeepEqual(outputTrustedOrigin.Name, testTrustedOrigin.Name) {
		t.Errorf("client.TrustedOrigins.UpdateTrustedOrigin returned \n\t%+v, want \n\t%+v\n", outputTrustedOrigin.Name, testTrustedOrigin.Name)
	}
}

func TestTrustedOriginDelete(t *testing.T) {
	setup()
	defer teardown()
	setupTestTrustedOrigin()

	testTrustedOrigin.ID = "ow1y4s2cpS59f1xs2p7"

	mux.HandleFunc("/trustedOrigins/ow1y4s2cpS59f1xs2p7", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		testAuthHeader(t, r)
		fmt.Fprint(w, "")
	})

	_, err := client.TrustedOrigins.DeleteTrustedOrigin("ow1y4s2cpS59f1xs2p7")
	if err != nil {
		t.Errorf("TrustedOrigins.DeleteTrustedOrigin returned error: %v", err)
	}
}

func TestTrustedOriginActivate(t *testing.T) {
	setup()
	defer teardown()
	setupTestTrustedOrigin()

	testTrustedOrigin.ID = "ow1y4s2cpS59f1xs2p7"

	mux.HandleFunc("/trustedOrigins/ow1y4s2cpS59f1xs2p7/lifecycle/activate", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testAuthHeader(t, r)
		fmt.Fprint(w, "")
	})

	_, err := client.TrustedOrigins.ActivateTrustedOrigin("ow1y4s2cpS59f1xs2p7", true)
	if err != nil {
		t.Errorf("TrustedOrigins.ActivateTrustedOrigin returned error: %v", err)
	}
}
