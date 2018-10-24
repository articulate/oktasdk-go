package okta

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"reflect"
	"testing"
)

func loadRawTestData(t *testing.T) []byte {
	p, err := filepath.Abs("../test_data/factors.json")
	if err != nil {
		t.Errorf("failed to resolve path, error %v", err)
	}

	data, err := ioutil.ReadFile(p)
	if err != nil {
		t.Errorf("failed to load %s, error %v", p, err)
	}

	return data
}

func TestListFactors(t *testing.T) {
	var testResponse []*Factor
	setup()
	defer teardown()
	raw := loadRawTestData(t)
	json.Unmarshal(raw, &testResponse)

	mux.HandleFunc("/org/factors", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testAuthHeader(t, r)
		fmt.Fprint(w, string(raw))
	})

	factors, _, err := client.Org.ListFactors()
	if err != nil {
		t.Errorf("Org.ListFactors returned error: %v", err)
	}
	if !reflect.DeepEqual(factors, testResponse) {
		t.Errorf("Org.ListFactors returned \n\t%+v, want \n\t%+v\n", factors, testResponse)
	}
}

func TestActivateFactor(t *testing.T) {
	testLifecycle("activate", t)
}

func TestDeactivateFactor(t *testing.T) {
	testLifecycle("deactivate", t)
}

func testLifecycle(act string, t *testing.T) {
	var (
		factor      *Factor
		err         error
		testFactors []*Factor
	)
	setup()
	defer teardown()
	raw := loadRawTestData(t)
	json.Unmarshal(raw, &testFactors)
	testFactor := testFactors[0]
	testFactorRaw, _ := json.Marshal(testFactor)

	relUrl := fmt.Sprintf("/org/factors/%s/lifecycle/%s", testFactor.Id, act)
	mux.HandleFunc(relUrl, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testAuthHeader(t, r)
		fmt.Fprint(w, string(testFactorRaw))
	})

	switch act {
	case "activate":
		factor, _, err = client.Org.ActivateFactor(testFactor.Id)
	case "deactivate":
		factor, _, err = client.Org.DeactivateFactor(testFactor.Id)
	}

	if err != nil {
		t.Errorf("Org lifecycle test returned error: %v", err)
	}
	if !reflect.DeepEqual(factor, testFactor) {
		t.Errorf("Org lifecycle test returned \n\t%+v, want \n\t%+v\n", factor, testFactor)
	}
}
