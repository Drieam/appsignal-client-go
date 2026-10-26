package appsignal

import (
	"context"
	"time"
)

type LogExport struct {
	ID              string         `json:"id"`
	S3URL           string         `json:"s3Url"`
	S3AccessKey     string         `json:"s3AccessKey"`
	State           LogExportState `json:"state"`
	Compressed      bool           `json:"compressed"`
	RetryCount      int32          `json:"retryCount"`
	LastRequestedAt time.Time      `json:"lastRequestedAt"`
}

const getLogExportQuery = `
query GetLogSourceExport($appId: String!, $logSourceId: String!) {
	app(id: $appId) {
		logs {
			source(id: $logSourceId) {
				export {
					id
					s3Url
					s3AccessKey
					state
					compressed
					retryCount
					lastRequestedAt
				}
			}
		}
	}
}
`

// GetLogSourceExport gets the export of a single log source by app ID and log source ID. It returns a
// nil log export when the log source has no export.
func (c *Client) GetLogSourceExport(ctx context.Context, appID string, logSourceID string) (*LogExport, error) {
	var out struct {
		App struct {
			Logs struct {
				Source struct {
					LogExport *LogExport `json:"export"`
				} `json:"source"`
			} `json:"logs"`
		} `json:"app"`
	}

	if err := c.Query(ctx, getLogExportQuery, map[string]any{"appId": appID, "logSourceId": logSourceID}, &out); err != nil {
		return nil, err
	}

	return out.App.Logs.Source.LogExport, nil
}

const createLogExportMutation = `
mutation CreateLogSourceExport(
	$appId: String!,
	$logSourceId: String!,
	$s3Url: String!,
	$s3AccessKey: String!,
	$s3SecretKey: String!,
	$compressed: Boolean!
) {
  createLogExport(
		appId: $appId
		sourceId: $logSourceId
		s3Url: $s3Url
		s3AccessKey: $s3AccessKey
		s3SecretKey: $s3SecretKey
		compressed: $compressed
	) {
		id
		s3Url
		s3AccessKey
		state
		compressed
		retryCount
		lastRequestedAt
	}
}
`

type CreateLogSourceExportInput struct {
	AppID       string `json:"appId"`
	LogSourceID string `json:"logSourceId"`
	S3URL       string `json:"s3Url"`
	S3AccessKey string `json:"s3AccessKey"`
	S3SecretKey string `json:"s3SecretKey"`
	Compressed  bool   `json:"compressed"`
}

// CreateLogSourceExport creates an export for a log source and returns the created log export.
func (c *Client) CreateLogSourceExport(ctx context.Context, input CreateLogSourceExportInput) (*LogExport, error) {
	var out struct {
		LogExport *LogExport `json:"createLogExport"`
	}

	variables := map[string]any{
		"appId":       input.AppID,
		"logSourceId": input.LogSourceID,
		"s3Url":       input.S3URL,
		"s3AccessKey": input.S3AccessKey,
		"s3SecretKey": input.S3SecretKey,
		"compressed":  input.Compressed,
	}

	if err := c.Mutate(ctx, createLogExportMutation, variables, &out); err != nil {
		return nil, err
	}

	return out.LogExport, nil
}

const updateLogExportMutation = `
mutation UpdateLogSourceExport(
	$appId: String!,
	$logExportId: String!
	$s3Url: String,
	$s3AccessKey: String,
	$s3SecretKey: String,
	$compressed: Boolean
) {
  updateLogExport(
		appId: $appId
		id: $logExportId
		s3Url: $s3Url
		s3AccessKey: $s3AccessKey
		s3SecretKey: $s3SecretKey
		compressed: $compressed
	) {
		id
		s3Url
		s3AccessKey
		state
		compressed
		retryCount
		lastRequestedAt
	}
}
`

type UpdateLogSourceExportInput struct {
	AppID       string  `json:"appId"`
	LogExportID string  `json:"logExportId"`
	S3URL       *string `json:"s3Url,omitempty"`
	S3AccessKey *string `json:"s3AccessKey,omitempty"`
	S3SecretKey *string `json:"s3SecretKey,omitempty"`
	Compressed  *bool   `json:"compressed,omitempty"`
}

// UpdateLogSourceExport changes an existing log export and returns the updated log export.
// Only the fields you set on the input are sent; the rest stay as they are.
func (c *Client) UpdateLogSourceExport(ctx context.Context, input UpdateLogSourceExportInput) (*LogExport, error) {
	var out struct {
		LogExport *LogExport `json:"updateLogExport"`
	}

	variables := map[string]any{
		"appId":       input.AppID,
		"logExportId": input.LogExportID,
	}

	if input.S3URL != nil {
		variables["s3Url"] = input.S3URL
	}
	if input.S3AccessKey != nil {
		variables["s3AccessKey"] = input.S3AccessKey
	}
	if input.S3SecretKey != nil {
		variables["s3SecretKey"] = input.S3SecretKey
	}
	if input.Compressed != nil {
		variables["compressed"] = input.Compressed
	}

	if err := c.Mutate(ctx, updateLogExportMutation, variables, &out); err != nil {
		return nil, err
	}

	return out.LogExport, nil
}

const deleteLogExportMutation = `
mutation DeleteLogSourceExport(
	$appId: String!,
	$logExportId: String!
) {
	deleteLogExport(
		appId: $appId
		id: $logExportId
	) {
		id
		s3Url
		s3AccessKey
		state
		compressed
		retryCount
		lastRequestedAt
	}
}
`

type DeleteLogSourceExportInput struct {
	AppID       string `json:"appId"`
	LogExportID string `json:"logExportId"`
}

// DeleteLogSourceExport removes an export from a log source and returns
// the log export as it was before deletion.
func (c *Client) DeleteLogSourceExport(ctx context.Context, input DeleteLogSourceExportInput) (*LogExport, error) {
	var out struct {
		LogExport *LogExport `json:"deleteLogExport"`
	}

	variables := map[string]any{
		"appId":       input.AppID,
		"logExportId": input.LogExportID,
	}

	if err := c.Mutate(ctx, deleteLogExportMutation, variables, &out); err != nil {
		return nil, err
	}

	return out.LogExport, nil
}
