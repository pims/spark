package spark

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
)

type DevicesService struct {
	client *SparkClient
}

type Device struct {
	Name      string `json:"name,omitempty"`
	Id        string `json:"id,omitempty"`
	LastApp   string `json:"last_app,omitempty"`
	LastHeard string `json:"last_heard,omitempty"`
	Connected bool   `json:"connected, omitempty"`
}

func (d Device) String() string {
	return fmt.Sprintf("Device: %s [%s], connected?: %v", d.Name, d.Id, d.Connected)
}

type Info struct {
	Name      string   `json:"name,omitempty"`
	Id        string   `json:"id,omitempty"`
	Variables []string `json:"variables,omitempty"`
	Functions []string `json:"functions,omitempty"`
}

func (i Info) String() string {
	return fmt.Sprintf("Info for %s [%s]\n variables: %s\n functions: %s\n", i.Name, i.Id, i.Variables, i.Functions)
}

func (s *DevicesService) List() ([]Device, *http.Response, error) {

	req, err := s.client.NewRequest("GET", "v1/devices", nil)
	if err != nil {
		return nil, nil, err
	}

	devices := new([]Device)
	resp, err := s.client.Do(req, devices)
	if err != nil {
		return nil, resp, err
	}

	return *devices, resp, err
}

func (s *DevicesService) Rename(coreId, name string) (*http.Response, error) {
	data := url.Values{}
	data.Set("name", name)
	body := bytes.NewBufferString(data.Encode())
	req, err := s.client.NewRequest("PUT", "v1/devices/"+coreId, body)

	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req, nil)
	return resp, err
}

func (s *DevicesService) Claim(coredId string) (*http.Response, error) {
	return nil, nil
}

func (s *DevicesService) Info(coreId string) (Info, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "v1/devices/"+coreId, nil)

	if err != nil {
		return Info{}, nil, err
	}
	var info = new(Info)
	resp, err := s.client.Do(req, info)

	return *info, resp, err
}

func (s *DevicesService) Read(coreId, varName string) (string, *http.Response, error) {
	path := fmt.Sprintf("v1/devices/%s/%s", coreId, varName)
	req, err := s.client.NewRequest("GET", path, nil)

	if err != nil {
		return "", nil, err
	}

	var stuff string
	resp, err := s.client.Do(req, stuff)
	return stuff, resp, err
}

func (s *DevicesService) Exec(coreId, funcName string) (*http.Response, error) {
	data := url.Values{}
	data.Set("func", funcName)
	body := bytes.NewBufferString(data.Encode())
	path := fmt.Sprintf("v1/devices/%s/%s", coreId, funcName)
	req, err := s.client.NewRequest("POST", path, body)

	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req, nil)
	return resp, err
}
