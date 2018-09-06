package gapi

import (
	"bytes"
	"encoding/json"
	"time"

	"fmt"
	"io/ioutil"

	"github.com/pkg/errors"
)

type Dashboard struct {
	Meta  DashboardMeta          `json:"meta"`
	Model map[string]interface{} `json:"dashboard"`
}

type DashboardMeta struct {
	Type        string    `json:"type"`
	CanSave     bool      `json:"canSave"`
	CanEdit     bool      `json:"canEdit"`
	CanAdmin    bool      `json:"canAdmin"`
	CanStar     bool      `json:"canStar"`
	Slug        string    `json:"slug"`
	Url         string    `json:"url"`
	Expires     time.Time `json:"expires"`
	Created     time.Time `json:"created"`
	Updated     time.Time `json:"updated"`
	UpdatedBy   string    `json:"updatedBy"`
	CreatedBy   string    `json:"createdBy"`
	Version     int       `json:"version"`
	HasAcl      bool      `json:"hasAcl"`
	IsFolder    bool      `json:"isFolder"`
	FolderId    int       `json:"folderId"`
	FolderTitle string    `json:"folderTitle"`
	FolderUrl   string    `json:"folderUrl"`
	Provisioned bool      `json:"provisioned"`
}

type DashboardSaveOpts struct {
	Model     map[string]interface{} `json:"dashboard"`
	Overwrite bool                   `json:"overwrite"`
	FolderID  int                    `json:"folderId"`
}

type DashboardSaveResponse struct {
	Id      int    `json:"id"`
	Uid     string `json:"uid"`
	Url     string `json:"url"`
	Slug    string `json:"slug"`
	Status  string `json:"status"`
	Version int64  `json:"version"`
}

func (c *Client) SaveDashboard(d *DashboardSaveOpts) (*DashboardSaveResponse, error) {
	data, err := json.Marshal(d)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to marshall dashboard JSON")
	}
	req, err := c.newRequest("POST", "/api/dashboards/db", nil, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to perform HTTP request")
	}

	if resp.StatusCode != 200 {
		var gmsg GrafanaErrorMessage
		dec := json.NewDecoder(resp.Body)
		dec.Decode(&gmsg)
		return nil, &GrafanaError{resp.StatusCode, fmt.Sprint(gmsg)}
	}

	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result := &DashboardSaveResponse{}
	err = json.Unmarshal(data, &result)
	return result, err
}

func (c *Client) GetDashboardByUID(uid string) (*Dashboard, error) {
	path := fmt.Sprintf("/api/dashboards/uid/%s", uid)
	req, err := c.newRequest("GET", path, nil, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		var gmsg GrafanaErrorMessage
		dec := json.NewDecoder(resp.Body)
		dec.Decode(&gmsg)
		return nil, &GrafanaError{resp.StatusCode, fmt.Sprint(gmsg)}
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result := &Dashboard{}
	err = json.Unmarshal(data, &result)
	return result, err
}

func (c *Client) DeleteDashboard(slug string) error {
	path := fmt.Sprintf("/api/dashboards/db/%s", slug)
	req, err := c.newRequest("DELETE", path, nil, nil)
	if err != nil {
		return err
	}

	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		var gmsg GrafanaErrorMessage
		dec := json.NewDecoder(resp.Body)
		dec.Decode(&gmsg)
		return fmt.Errorf("Request to Grafana returned %+v status code with the following message: %+v", resp.StatusCode, gmsg.Message)
	}

	return nil
}
