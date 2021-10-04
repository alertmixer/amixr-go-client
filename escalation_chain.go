package amixr

import (
	"fmt"
	"net/http"
)

type EscalationChainService struct {
	client *Client
	url    string
}

func NewEscalationChainService(client *Client) *EscalationChainService {
	escalationChainService := EscalationChainService{}
	escalationChainService.client = client
	escalationChainService.url = "escalation_chains"
	return &escalationChainService
}

type PaginatedEscalationChainsResponse struct {
	PaginatedResponse
	EscalationChains []*EscalationChain `json:"results"`
}

type EscalationChain struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
    TeamId    string `json:"team_id"`
}

type ListEscalationChainOptions struct {
	ListOptions
	Name string `url:"name,omitempty" json:"name,omitempty"`
}

func (service *EscalationChainService) ListEscalationChains(opt *ListEscalationChainOptions) (*PaginatedEscalationChainsResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/", service.url)

	req, err := service.client.NewRequest("GET", u, opt)
	if err != nil {
		return nil, nil, err
	}

	var escalation_chains *PaginatedEscalationChainsResponse
	resp, err := service.client.Do(req, &escalation_chains)
	if err != nil {
		return nil, resp, err
	}

	return escalation_chains, resp, err
}

type GetEscalationChainOptions struct {
}

func (service *EscalationChainService) GetEscalationChain(id string, opt *GetEscalationChainOptions) (*EscalationChain, *http.Response, error) {
	u := fmt.Sprintf("%s/%s/", service.url, id)

	req, err := service.client.NewRequest("GET", u, opt)
	if err != nil {
		return nil, nil, err
	}

	escalation_chain := new(EscalationChain)
	resp, err := service.client.Do(req, escalation_chain)
	if err != nil {
		return nil, resp, err
	}

	return escalation_chain, resp, err
}

type CreateEscalationChainOptions struct {
	Name string `json:"name,omitempty"`
    TeamId string `json:"team_id"`
}

func (service *EscalationChainService) CreateEscalationChain(opt *CreateEscalationChainOptions) (*EscalationChain, *http.Response, error) {
	u := fmt.Sprintf("%s/", service.url)
	req, err := service.client.NewRequest("POST", u, opt)
	if err != nil {
		return nil, nil, err
	}

	escalationChain := new(EscalationChain)

	resp, err := service.client.Do(req, escalationChain)

	if err != nil {
		return nil, resp, err
	}

	return escalationChain, resp, err
}

type UpdateEscalationChainOptions struct {
	Name string `json:"name,omitempty"`
	TeamId string `json:"team_id"`
}

func (service *EscalationChainService) UpdateEscalationChain(id string, opt *UpdateEscalationChainOptions) (*EscalationChain, *http.Response, error) {
	u := fmt.Sprintf("%s/%s/", service.url, id)

	req, err := service.client.NewRequest("PUT", u, opt)
	if err != nil {
		return nil, nil, err
	}

	escalationChain := new(EscalationChain)
	resp, err := service.client.Do(req, escalationChain)
	if err != nil {
		return nil, resp, err
	}

	return escalationChain, resp, err
}

type DeleteEscalationChainOptions struct {
}

func (service *EscalationChainService) DeleteEscalationChain(id string, opt *DeleteEscalationChainOptions) (*http.Response, error) {

	u := fmt.Sprintf("%s/%s/", service.url, id)

	req, err := service.client.NewRequest("DELETE", u, opt)
	if err != nil {
		return nil, err
	}

	resp, err := service.client.Do(req, nil)
	return resp, err
}
