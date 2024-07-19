package cloudvalid

import (
	"net/http"
	"time"
)

type CloudValidClient struct {
	APIKey string
	client *http.Client
}

// The CreateDNSConfigurationLinkRequest struct is used to construct the request
// payload for creating a DNS configuration link.
type CreateDNSConfigurationLinkRequest struct {
	Domain        string            `json:"domain"` // Required
	Variables     map[string]string `json:"variables"`
	UseCases      []string          `json:"use_cases"`
	RawDNSRecords []struct {
		Type     string `json:"type"`
		Host     string `json:"host"`
		Content  string `json:"content"`
		Priority string `json:"priority"`
	} `json:"raw_dns_records"`
}

// The CreateDNSConfigurationLinkResponse struct is used for parsing the successful
// response from the `create-dns-configuration-link` request.
type CreateDNSConfigurationLinkResponse struct {
	Result struct {
		ID               string            `json:"id"`
		Domain           string            `json:"domain"`
		Variables        map[string]string `json:"variables"`
		UseCases         []string          `json:"use_cases"`
		CreatedAt        time.Time         `json:"created_at"`
		DNSProviderGuess struct {
			Name      string `json:"name"`
			URL       string `json:"url"`
			Logo      string `json:"logo"`
			TargetURL string `json:"target_url"`
		} `json:"dns_provider_guess"`
		DNSProviderDomain string     `json:"dns_provider_domain"`
		CancelledAt       *time.Time `json:"cancelled_at"`
		FinishedAt        *time.Time `json:"finished_at"`
		PublicURL         string     `json:"public_url"`
		PublicURLLatest   string     `json:"public_url_latest"`
	} `json:"result"`
}

// Used for parsing the Get Hosted Page (DNS Configuration Link) response.
type HostedPageResponse struct {
	Result struct {
		ID                         string     `json:"id"`
		ServiceName                string     `json:"service_name"`
		RedirectURL                string     `json:"redirect_url"`
		Domain                     string     `json:"domain"`
		FinishedAt                 *time.Time `json:"finished_at"`
		DomainConnectURLNewWindow  string     `json:"domain_connect_url_new_window"`
		DomainConnectURLSameWindow string     `json:"domain_connect_url_same_window"`
		DNSProviderDomain          string     `json:"dns_provider_domain"`
		PublicURL                  string     `json:"public_url"`
		PublicURLLatest            string     `json:"public_url_latest"`
		UseCases                   []string   `json:"use_cases"`
		DNSProviderGuess           struct {
			Name      string `json:"name"`
			URL       string `json:"url"`
			Logo      string `json:"logo"`
			TargetURL string `json:"target_url"`
		} `json:"dns_provider_guess"`
		Client struct {
			Name                string `json:"name"`
			BrandPrimaryColor   string `json:"brand_primary_color"`
			BrandSecondaryColor string `json:"brand_secondary_color"`
			BrandLogo           string `json:"brand_logo"`
		} `json:"client"`
		Records []struct {
			UseCase                    string   `json:"use_case"`
			Type                       string   `json:"type"`
			Host                       string   `json:"host"`
			Content                    string   `json:"content"`
			RemoveExistingRecords      []string `json:"remove_existing_records"`
			Propagated                 bool     `json:"propagated"`
			ConsiderExistingDMARCValid bool     `json:"consider_existing_dmarc_valid"`
		} `json:"records"`
	} `json:"result"`
}

// Used for parsing the response of the get-propagation-status API requests.
type PropagationStatusResponse struct {
	Result []struct {
		UseCase                    string   `json:"use_case"`
		Type                       string   `json:"type"`
		Host                       string   `json:"host"`
		Content                    string   `json:"content"`
		RemoveExistingRecords      []string `json:"remove_existing_records"`
		Message                    string   `json:"message"`
		Suggestion                 string   `json:"suggestion"`
		ConsiderExistingDMARCValid bool     `json:"consider_existing_dmarc_valid"`
		Propagated                 bool     `json:"propagated"`
	}
}

// Used for parsing the response of Cancel DNS Configuration Link requests.
type CancelDNSConfigurationLinkResponse struct {
	ID                         string     `json:"id"`
	UserEmail                  string     `json:"user_email"`
	Domain                     string     `json:"domain"`
	CreatedAt                  time.Time  `json:"created_at"`
	UpdatedAt                  *time.Time `json:"updated_at"`
	CancelledAt                *time.Time `json:"cancelled_at"`
	FinishedAt                 *time.Time `json:"deleted_at"`
	PublicURL                  string     `json:"public_url"`
	PublicURLLatest            string     `json:"public_url_latest"`
	LastPropagationCheckAt     *time.Time `json:"last_propagation_check_at"`
	LastPropagationCheckStatus string     `json:"last_propagation_check_status"`
	PropagationChecksCount     int        `json:"propagation_checks_count"`
	Status                     string     `json:"status"`
	DNSProviderDomain          string     `json:"dns_provider_domain"`
	DNSProviderGuess           struct {
		Name      string `json:"name"`
		URL       string `json:"url"`
		Logo      string `json:"logo"`
		TargetURL string `json:"target_url"`
	} `json:"dns_provider_guess"`
	Records []struct {
		UseCase                    string   `json:"use_case"`
		Type                       string   `json:"type"`
		Host                       string   `json:"host"`
		Content                    string   `json:"content"`
		RemoveExistingRecords      []string `json:"remove_existing_records"`
		Propagated                 bool     `json:"propagated"`
		ConsiderExistingDMARCValid bool     `json:"consider_existing_dmarc_valid"`
	} `json:"records"`
	Templates []struct {
		ID             string     `json:"id"`
		TemplateName   string     `json:"template_name"`
		CreatedAt      time.Time  `json:"created_at"`
		UpdatedAt      *time.Time `json:"updated_at"`
		UseCase        string     `json:"use_case"`
		UseCaseDisplay string     `json:"use_case_display"`
		Records        []struct {
			Type                       string   `json:"type"`
			Host                       string   `json:"host"`
			Content                    string   `json:"content"`
			Priority                   string   `json:"priority"`
			RemoveExisting             bool     `json:"removeExisting"`
			ConsiderExistingDMARCValid bool     `json:"considerExistingDmarcValid"`
			Errors                     []string `json:"errors"`
		} `json:"records"`
		Variables []string `json:"variables"`
	}
}
