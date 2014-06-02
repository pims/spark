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
	Name      string            `json:"name,omitempty"`
	Id        string            `json:"id,omitempty"`
	Variables map[string]string `json:"variables,omitempty"`
	Functions []string          `json:"functions,omitempty"`
}

type ExecResponse struct {
	Device
	ReturnValue int `json:"return_value,omitempty"`
}

type ReadResponse struct {
	Device  Device `json:"coreInfo,omitempty"`
	Command string
	Name    string
	Result  interface{}
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

func (s *DevicesService) Read(coreId, varName string) (ReadResponse, *http.Response, error) {
	path := fmt.Sprintf("v1/devices/%s/%s", coreId, varName)
	req, err := s.client.NewRequest("GET", path, nil)

	if err != nil {
		return ReadResponse{}, nil, err
	}

	readResponse := new(ReadResponse)
	resp, err := s.client.Do(req, readResponse)
	return *readResponse, resp, err
}

func (s *DevicesService) Exec(coreId, funcName, params string) (ExecResponse, *http.Response, error) {
	data := url.Values{}
	data.Set("params", params)
	body := bytes.NewBufferString(data.Encode())
	path := fmt.Sprintf("v1/devices/%s/%s", coreId, funcName)
	req, err := s.client.NewRequest("POST", path, body)

	if err != nil {
		return ExecResponse{}, nil, err
	}
	exec := new(ExecResponse)
	resp, err := s.client.Do(req, exec)
	return *exec, resp, err
}
