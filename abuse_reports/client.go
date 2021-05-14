// Code generated by apic. DO NOT EDIT.

package abuse_reports

import (
	"bytes"
	"context"
	"net/url"
	"text/template"

	"github.com/ngrok/ngrok-api-go"
)

type Client struct {
	apiClient *ngrok.Client
}

func NewClient(apiClient *ngrok.Client) *Client {
	return &Client{apiClient: apiClient}
}

// Creates a new abuse report which will be reviewed by our system and abuse
// response team. This API is only available to authorized accounts. Contact
// abuse@ngrok.com to request access
func (c *Client) Create(
	ctx context.Context,
	arg *ngrok.AbuseReportCreate,
) (*ngrok.AbuseReport, error) {
	var res ngrok.AbuseReport
	var path bytes.Buffer
	if err := template.Must(template.New("create_path").Parse("/abuse_reports")).Execute(&path, arg); err != nil {
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

// Get the detailed status of abuse report by ID.
func (c *Client) Get(
	ctx context.Context,
	id string,

) (*ngrok.AbuseReport, error) {
	arg := &ngrok.Item{ID: id}
	var res ngrok.AbuseReport
	var path bytes.Buffer
	if err := template.Must(template.New("get_path").Parse("/abuse_reports/{{ .ID }}")).Execute(&path, arg); err != nil {
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
