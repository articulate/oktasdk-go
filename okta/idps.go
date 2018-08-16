package okta

import (
	"fmt"
	"time"
)

type IdentityProvidersService service

func (p *IdentityProvidersService) IdentityProvider() IdentityProvider {
	return IdentityProvider{}
}

type IdentityProvider struct {
	ID          string    `json:"id,omitempty"`
	Type        string    `json:"type,omitempty"`
	Status      string    `json:"status,omitempty"`
	Name        string    `json:"name,omitempty"`
	Created     time.Time `json:"created,omitempty"`
	LastUpdated time.Time `json:"lastUpdated,omitempty"`
	Protocol    struct {
		Type      string `json:"type,omitempty"`
		Endpoints struct {
			Authorization struct {
				Url     string `json:"url,omitempty"`
				Binding string `json:"binding,omitempty"`
			} `json:"authorization,omitempty"`
			Token struct {
				Url     string `json:"url,omitempty"`
				Binding string `json:"binding,omitempty"`
			}
		} `json:"endpoints,omitempty"`
		Scopes      []string `json:"scopes,omitempty"`
		Credentials struct {
			Client struct {
				ClientID     string `json:"client_id,omitempty"`
				ClientSecret string `json:"client_secret,omitempty"`
			} `json:"client,omitempty"`
		} `json:"credentials,omitempty"`
	} `json:"protocol,omitempty"`
	Policy struct {
		Provisioning struct {
			Action        string `json:"action,omitempty"`
			ProfileMaster bool   `json:"profileMaster,omitempty"`
			Groups        struct {
				Action string `json:"action,omitempty"`
			} `json:"groups,omitempty"`
			Conditions struct {
				Deprovisioned struct {
					Action string `json:"action,omitempty"`
				} `json:"deprovisioned,omitempty"`
				Suspended struct {
					Action string `json:"action,omitempty"`
				} `json:"suspended,omitempty"`
			} `json:"conditions,omitempty"`
		} `json:"provisioning,omitempty"`
		AccountLink struct {
			Filter string `json:"filter,omitempty"`
			Action string `json:"action,omitempty"`
		} `json:"accountLink,omitempty"`
		Subject struct {
			UserNameTemplate struct {
				Template string `json:"template,omitempty"`
			} `json:"userNameTemplate,omitempty"`
			Filter    string `json:"filter,omitempty"`
			MatchType string `json:"matchType,omitempty"`
		} `json:"subject,omitempty"`
		MaxClockSkew int `json:"maxClockSkew,omitempty"`
	} `json:"policy,omitempty"`
	Links struct {
		Authorize struct {
			Href      string `json:"href,omitempty"`
			Templated bool   `json:"templated,omitempty"`
			Hints     struct {
				Allow []string `json:"allow,omitempty"`
			} `json:"hints,omitempty"`
		} `json:"authorize,omitempty"`
		ClientRedirectUri struct {
			Href  string `json:"href,omitempty"`
			Hints struct {
				Allow []string `json:"allow,omitempty"`
			} `json:"hints,omitempty"`
		} `json:"clientRedirectUri,omitempty"`
	} `json:"_links,omitempty"`
}

// GetIdentityProvider: Get an IdP
// Requires IdentityProvider ID from IdentityProvider object
func (p *IdentityProvidersService) GetIdentityProvider(id string) (*IdentityProvider, *Response, error) {
	u := fmt.Sprintf("idps/%v", id)
	req, err := p.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	idp := new(IdentityProvider)
	resp, err := p.client.Do(req, idp)
	if err != nil {
		return nil, resp, err
	}

	return idp, resp, err
}

// CreateIdentityprovider: Create a identityprovider
// You must pass in the Identityprovider object created from the desired input identityprovider
func (p *IdentityProvidersService) CreateIdentityProvider(idp interface{}) (*IdentityProvider, *Response, error) {
	u := fmt.Sprintf("idps")
	req, err := p.client.NewRequest("POST", u, idp)
	if err != nil {
		return nil, nil, err
	}

	newIdp := new(IdentityProvider)
	resp, err := p.client.Do(req, newIdp)
	if err != nil {
		return nil, resp, err
	}

	return newIdp, resp, err
}

// UpdateIdentityProvider: Update a policy
// Requires IdentityProvider ID from IdentityProvider object & IdentityProvider object from the desired input policy
func (p *IdentityProvidersService) UpdateIdentityProvider(id string, idp interface{}) (*IdentityProvider, *Response, error) {
	u := fmt.Sprintf("idps/%v", id)
	req, err := p.client.NewRequest("PUT", u, idp)
	if err != nil {
		return nil, nil, err
	}

	updateIdentityProvider := new(IdentityProvider)
	resp, err := p.client.Do(req, updateIdentityProvider)
	if err != nil {
		return nil, resp, err
	}

	return updateIdentityProvider, resp, err
}

// DeleteIdentityprovider: Delete a identityprovider
// Requires Identityprovider ID from Identityprovider object
func (p *IdentityProvidersService) DeleteIdentityProvider(id string) (*Response, error) {
	u := fmt.Sprintf("idps/%v", id)
	req, err := p.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}
	resp, err := p.client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, err
}
