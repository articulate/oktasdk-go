package okta

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

var testGroup *Group
var testGroupProfile *GroupProfile
var testGroupLinks *GroupLinks

func setupTestGroup() {
	testGroupLinks = &GroupLinks{}

	testGroup = &Group{
		GroupLinks:        testGroupLinks,
		GroupProfile: &GroupProfile{
			Name:         "testGroup",
			Description:  "testGroupDescription",
		},
	}
}

func TestGroupCreate(t *testing.T) {

	setup()
	defer teardown()
	setupTestGroup()

	temp, err := json.Marshal(testGroup)

	if err != nil {
		t.Errorf("Groups json Marshall returned error: %v", err)
	}

	GroupTestJSONString := string(temp)

	mux.HandleFunc("/groups", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testAuthHeader(t, r)
		fmt.Fprint(w, GroupTestJSONString)
	})

	outputGroup, _, err := client.Groups.Add(testGroup.GroupProfile.Name, testGroup.GroupProfile.Description)
	if err != nil {
		t.Errorf("GroupsService.Add returned error: %v", err)
	}
	if !reflect.DeepEqual(outputGroup, testGroup) {
		t.Errorf("client.GroupsService.Add returned \n\t%+v, want \n\t%+v\n", outputGroup, testGroup)
	}
}

func TestGetGroup(t *testing.T) {
	setup()
	defer teardown()
	setupTestGroup()
	testGroup.ID = "0oa62bfdiumsUndnZ0h7"

	temp, err := json.Marshal(testGroup)
	if err != nil {
		t.Errorf("Groups.GetByID json Marshall returned error: %v", err)
	}
	GroupTestJSONString := string(temp)

	mux.HandleFunc("/groups/0oa62bfdiumsUndnZ0h7", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testAuthHeader(t, r)
		fmt.Fprint(w, GroupTestJSONString)
	})

	outputGroup, _, err := client.Groups.GetByID("0oa62bfdiumsUndnZ0h7")
	if err != nil {
		t.Errorf("Groups.GetGroup returned error: %v", err)
	}
	if !reflect.DeepEqual(outputGroup, testGroup) {
		t.Errorf("client.Groups.GetGroup returned \n\t%+v, want \n\t%+v\n", outputGroup, testGroup)
	}
}

func TestGroupUpdate(t *testing.T) {
	setup()
	defer teardown()
	setupTestGroup()
	testGroup.ID = "0oa62bfdiumsUndnZ0h7"

	testGroup.GroupProfile.Name = "Something Completely Different"
	temp, err := json.Marshal(testGroup)
	if err != nil {
		t.Errorf("Groups.Update json Marshall returned error: %v", err)
	}
	updateTestJSONString := string(temp)

	mux.HandleFunc("/groups/0oa62bfdiumsUndnZ0h7", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		testAuthHeader(t, r)
		fmt.Fprint(w, updateTestJSONString)
	})

	outputGroup, _, err := client.Groups.Update("0oa62bfdiumsUndnZ0h7", testGroup)
	if err != nil {
		t.Errorf("Groups.UpdateGroup returned error: %v", err)
	}
	if !reflect.DeepEqual(outputGroup.GroupProfile.Name, testGroup.GroupProfile.Name) {
		t.Errorf("client.Groups.UpdateGroup returned \n\t%+v, want \n\t%+v\n", outputGroup.GroupProfile.Name, testGroup.GroupProfile.Name)
	}
}

func TestGroupDelete(t *testing.T) {
	setup()
	defer teardown()
	setupTestGroup()

	testGroup.ID = "0oa62bfdiumsUndnZ0h7"

	mux.HandleFunc("/groups/0oa62bfdiumsUndnZ0h7", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		testAuthHeader(t, r)
		fmt.Fprint(w, "")
	})

	_, err := client.Groups.Delete("0oa62bfdiumsUndnZ0h7")
	if err != nil {
		t.Errorf("Groups.Delete returned error: %v", err)
	}

}
