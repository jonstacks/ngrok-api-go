// Code generated for API Clients. DO NOT EDIT.

package reserved_domains

import (
	"bytes"
	"context"
	"net/url"
	"text/template"

	"github.com/ngrok/ngrok-api-go/v5"
	"github.com/ngrok/ngrok-api-go/v5/internal/api"
)

type Iter = api.Iter[ngrok.ReservedDomain, ngrok.ReservedDomainList]

// Reserved Domains are hostnames that you can listen for traffic on. Domains
//  can be used to listen for http, https or tls traffic. You may use a domain
//  that you own by creating a CNAME record specified in the returned resource.
//  This CNAME record points traffic for that domain to ngrok's edge servers.

type Client struct {
	apiClient *api.Client
}

func NewClient(cfg *ngrok.ClientConfig) *Client {
	return &Client{apiClient: api.NewClient(cfg)}
}

// Create a new reserved domain.
//
// https://ngrok.com/docs/api#api-reserved-domains-create
func (c *Client) Create(ctx context.Context, arg *ngrok.ReservedDomainCreate) (*ngrok.ReservedDomain, error) {
	if arg == nil {
		arg = new(ngrok.ReservedDomainCreate)
	}
	var res ngrok.ReservedDomain
	var path bytes.Buffer
	if err := template.Must(template.New("create_path").Parse("/reserved_domains")).Execute(&path, arg); err != nil {
		panic(err)
	}
	var (
		apiURL  = &url.URL{Path: path.String()}
		bodyArg interface{}
	)
	apiURL.Path = path.String()
	bodyArg = arg

	if err := c.apiClient.Do(ctx, "POST", apiURL, bodyArg, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// Delete a reserved domain.
//
// https://ngrok.com/docs/api#api-reserved-domains-delete
func (c *Client) Delete(ctx context.Context, id string) error {
	arg := &ngrok.Item{ID: id}

	var path bytes.Buffer
	if err := template.Must(template.New("delete_path").Parse("/reserved_domains/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	arg.ID = ""
	var (
		apiURL  = &url.URL{Path: path.String()}
		bodyArg interface{}
	)
	apiURL.Path = path.String()

	if err := c.apiClient.Do(ctx, "DELETE", apiURL, bodyArg, nil); err != nil {
		return err
	}
	return nil
}

// Get the details of a reserved domain.
//
// https://ngrok.com/docs/api#api-reserved-domains-get
func (c *Client) Get(ctx context.Context, id string) (*ngrok.ReservedDomain, error) {
	arg := &ngrok.Item{ID: id}

	var res ngrok.ReservedDomain
	var path bytes.Buffer
	if err := template.Must(template.New("get_path").Parse("/reserved_domains/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	arg.ID = ""
	var (
		apiURL  = &url.URL{Path: path.String()}
		bodyArg interface{}
	)
	apiURL.Path = path.String()

	if err := c.apiClient.Do(ctx, "GET", apiURL, bodyArg, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// List all reserved domains on this account.
//
// https://ngrok.com/docs/api#api-reserved-domains-list
func (c *Client) List(paging *ngrok.Paging) *Iter {
	if paging == nil {
		paging = new(ngrok.Paging)
	}
	var path bytes.Buffer
	if err := template.Must(template.New("list_path").Parse("/reserved_domains")).Execute(&path, paging); err != nil {
		panic(err)
	}
	var apiURL = &url.URL{Path: path.String()}
	apiURL.RawQuery = paging.URLValues().Encode()
	return api.NewIter[ngrok.ReservedDomain, ngrok.ReservedDomainList](c.apiClient, apiURL)
}

// Update the attributes of a reserved domain.
//
// https://ngrok.com/docs/api#api-reserved-domains-update
func (c *Client) Update(ctx context.Context, arg *ngrok.ReservedDomainUpdate) (*ngrok.ReservedDomain, error) {
	if arg == nil {
		arg = new(ngrok.ReservedDomainUpdate)
	}
	var res ngrok.ReservedDomain
	var path bytes.Buffer
	if err := template.Must(template.New("update_path").Parse("/reserved_domains/{{ .ID }}")).Execute(&path, arg); err != nil {
		panic(err)
	}
	arg.ID = ""
	var (
		apiURL  = &url.URL{Path: path.String()}
		bodyArg interface{}
	)
	apiURL.Path = path.String()
	bodyArg = arg

	if err := c.apiClient.Do(ctx, "PATCH", apiURL, bodyArg, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// Detach the certificate management policy attached to a reserved domain.
//
// https://ngrok.com/docs/api#api-reserved-domains-delete-certificate-management-policy
func (c *Client) DeleteCertificateManagementPolicy(ctx context.Context, id string) error {
	arg := &ngrok.Item{ID: id}

	var path bytes.Buffer
	if err := template.Must(template.New("delete_certificate_management_policy_path").Parse("/reserved_domains/{{ .ID }}/certificate_management_policy")).Execute(&path, arg); err != nil {
		panic(err)
	}
	arg.ID = ""
	var (
		apiURL  = &url.URL{Path: path.String()}
		bodyArg interface{}
	)
	apiURL.Path = path.String()

	if err := c.apiClient.Do(ctx, "DELETE", apiURL, bodyArg, nil); err != nil {
		return err
	}
	return nil
}

// Detach the certificate attached to a reserved domain.
//
// https://ngrok.com/docs/api#api-reserved-domains-delete-certificate
func (c *Client) DeleteCertificate(ctx context.Context, id string) error {
	arg := &ngrok.Item{ID: id}

	var path bytes.Buffer
	if err := template.Must(template.New("delete_certificate_path").Parse("/reserved_domains/{{ .ID }}/certificate")).Execute(&path, arg); err != nil {
		panic(err)
	}
	arg.ID = ""
	var (
		apiURL  = &url.URL{Path: path.String()}
		bodyArg interface{}
	)
	apiURL.Path = path.String()

	if err := c.apiClient.Do(ctx, "DELETE", apiURL, bodyArg, nil); err != nil {
		return err
	}
	return nil
}
