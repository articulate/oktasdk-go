package okta

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
)

var testAccountLink *AccountLink
var testAuthorization *Authorization
var testAuthorize *Authorize
var testClient *Client
var testClientRedirectUri *ClientRedirectUri
var testCredentials *Credentials
var testDeprovisioned *Deprovisioned
var testHints *Hints
var testIdentityProvider *IdentityProvider
var testLinks *links
var testPolicy *Policy
var testProtocol *Protocol
var testProvisioning *Provisioning
var testSuspended *Suspended
var testSubject *Subject
var testToken *Token
var testUserNameTemplate *UserNameTemplate

func setupTestIdentityProvider() {
	hmm, _ := time.Parse("2006-01-02T15:04:05.000Z", "2018-02-16T19:59:05.000Z")

	testAccountLink = &AccountLink{
		Action: "NONE",
		Filter: "NONE",
	}

	testPolicy = &Policy{
		MaxClockSkew: 0,
	}

	testProvisioning = &Provisioning{
		Action:        "NONE",
		ProfileMaster: false,
	}

	testSubject = &Subject{
		Filter:    "NONE",
		MatchType: "USERNAME",
	}

	testProtocol = &Protocol{
		Type: "OIDC",
	}

	testProtocol.Scopes = []string{"profile email openid"}
	testProtocol.Credentials.Client.ClientID = "your-client-id"
	testProtocol.Credentials.Client.ClientSecret = "your-client-secret"
	testProvisioning.Groups.Action = "NONE"
	testProvisioning.Conditions.Deprovisioned.Action = "NONE"
	testProvisioning.Conditions.Suspended.Action = "NONE"
	testSubject.UserNameTemplate.Template = "idpuser.userPrincipalName"

	testIdentityProvider = &IdentityProvider{
		Type: "GOOGLE",
		Name: "Google",
	}

	testIdentityProvider.Protocol = testProtocol
	testIdentityProvider.Policy = testPolicy
	testIdentityProvider.Policy.Provisioning = testProvisioning
	testIdentityProvider.Policy.AccountLink = testAccountLink
	testIdentityProvider.Policy.Subject = testSubject
}

func TestGetIdentityProvider(t *testing.T) {
	setup()
	defer teardown()
	setupTestIdentityProvider()
	testIdentityProvider.ID = "0oa62bfdiumsUndnZ0h7"

	temp, err := json.Marshal(testIdentityProvider)
	if err != nil {
		t.Errorf("IdentityProviders.GetIdentityProvider json Marshall returned error: %v", err)
	}
	IdentityProviderTestJSONString := string(temp)

	mux.HandleFunc("/idps/0oa62bfdiumsUndnZ0h7", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testAuthHeader(t, r)
		fmt.Fprint(w, IdentityProviderTestJSONString)
	})

	outputIdentityProvider, _, err := client.IdentityProviders.GetIdentityProvider("0oa62bfdiumsUndnZ0h7")
	if err != nil {
		t.Errorf("IdentityProviders.GetIdentityProvider returned error: %v", err)
	}
	if !reflect.DeepEqual(outputIdentityProvider, testIdentityProvider) {
		t.Errorf("client.IdentityProviders.GetIdentityProvider returned \n\t%+v, want \n\t%+v\n", outputIdentityProvider, testIdentityProvider)
	}
}

func TestIdentityProviderCreate(t *testing.T) {

	setup()
	defer teardown()
	setupTestIdentityProvider()

	temp, err := json.Marshal(testIdentityProvider)

	if err != nil {
		t.Errorf("IdentityProviders.CreateIdentityProvider json Marshall returned error: %v", err)
	}

	IdentityProviderTestJSONString := string(temp)

	mux.HandleFunc("/idps", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testAuthHeader(t, r)
		fmt.Fprint(w, IdentityProviderTestJSONString)
	})

	outputIdentityProvider, _, err := client.IdentityProviders.CreateIdentityProvider(testIdentityProvider)
	if err != nil {
		t.Errorf("IdentityProvider.CreateIdentityProvider returned error: %v", err)
	}
	if !reflect.DeepEqual(outputIdentityProvider, testIdentityProvider) {
		t.Errorf("client.IdentityProviders.CreateIdentityProvider returned \n\t%+v, want \n\t%+v\n", outputIdentityProvider, testIdentityProvider)
	}
}

func TestIdentityProviderUpdate(t *testing.T) {

	setup()
	defer teardown()
	setupTestIdentityProvider()
	testIdentityProvider.ID = "0oa62bfdiumsUndnZ0h7"

	testIdentityProvider.Name = "Something Completely Different"
	temp, err := json.Marshal(testIdentityProvider)
	if err != nil {
		t.Errorf("IdentityProviders.UpdateIdentityProvider json Marshall returned error: %v", err)
	}
	updateTestJSONString := string(temp)

	mux.HandleFunc("/idps/0oa62bfdiumsUndnZ0h7", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		testAuthHeader(t, r)
		fmt.Fprint(w, updateTestJSONString)
	})

	outputIdentityProvider, _, err := client.IdentityProviders.UpdateIdentityProvider("0oa62bfdiumsUndnZ0h7", testIdentityProvider)
	if err != nil {
		t.Errorf("IdentityProviders.UpdateIdentityProvider returned error: %v", err)
	}
	if !reflect.DeepEqual(outputIdentityProvider.Name, testIdentityProvider.Name) {
		t.Errorf("client.IdentityProviders.UpdateIdentityProvider returned \n\t%+v, want \n\t%+v\n", outputIdentityProvider.Name, testIdentityProvider.Name)
	}
}

func TestIdentityProviderDelete(t *testing.T) {

	setup()
	defer teardown()
	setupTestIdentityProvider()

	testIdentityProvider.ID = "0oa62bfdiumsUndnZ0h7"

	mux.HandleFunc("/idps/0oa62bfdiumsUndnZ0h7", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		testAuthHeader(t, r)
		fmt.Fprint(w, "")
	})

	_, err := client.IdentityProviders.DeleteIdentityProvider("0oa62bfdiumsUndnZ0h7")
	if err != nil {
		t.Errorf("IdentityProviders.DeleteIdentityProvider returned error: %v", err)
	}
}
