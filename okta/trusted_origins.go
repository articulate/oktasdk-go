package okta

import (
	"fmt"
	"time"
)

type TrustedOriginsService service

func (p *TrustedOriginsService) TrustedOrigin() TrustedOrigin {
	return TrustedOrigin{}
}

type TrustedOrigin struct {
	ID            string              `json:"id,omitempty"`
	Status        string              `json:"status,omitempty"`
	Name          string              `json:"name,omitempty"`
	Origin        string              `json:"origin,omitempty"`
	Scopes        []map[string]string `json:"scopes,omitempty"`
	Created       *time.Time          `json:"created,omitempty"`
	CreatedBy     string              `json:"createdBy,omitempty"`
	LastUpdated   *time.Time          `json:"lastUpdated,omitempty"`
	LastUpdatedBy string              `json:"lastUpdated,omitempty"`
	Links         *TrustedOriginLinks `json:"_links,omitempty"`
}

type TrustedOriginDeactive struct {
	Href  string              `json:"href,omitempty"`
	Hints *TrustedOriginHints `json:"hints,omitempty`
}

type TrustedOriginHints struct {
	Allow []string `json:"allow,omitempty"`
}

type TrustedOriginLinks struct {
	Self       *TrustedOriginSelf     `json:"self,omitempty"`
	Deactivate *TrustedOriginDeactive `json:"deactive,omitempty"`
}

type TrustedOriginSelf struct {
	Href  string              `json:"href,omitempty"`
	Hints *TrustedOriginHints `json:"hints,omitempty`
}

// GetTrustedOrigin: Get a Trusted Origin entry
// Requires TrustedOrigins ID from TrustedOrigins object
func (p *TrustedOriginsService) GetTrustedOrigin(id string) (*TrustedOrigin, *Response, error) {
	u := fmt.Sprintf("trustedOrigins/%v", id)
	req, err := p.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	trustedOrigin := new(TrustedOrigin)
	resp, err := p.client.Do(req, trustedOrigin)
	if err != nil {
		return nil, resp, err
	}

	return trustedOrigin, resp, err
}
