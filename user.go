package amixr

import (
	"fmt"
	"net/http"
)

// Handles requests to user endpoint
// Use NewUserService instead of direct creation UserService
//
// http://api-docs.amixr.io/#users
type UserService struct {
	client *Client
	url    string
}

// NewUsersService creates UserService with defined url
func NewUserService(client *Client) *UserService {
	userService := UserService{}
	userService.client = client
	userService.url = "users"
	return &userService
}

type PaginatedUsersResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Users    []User `json:"results"`
}

type User struct {
	ID     string `json:"id"`
	TeamId string `json:"team_id"`
	Name   string `json:"name"`
	Role   string `json:"role"`
	Email  string `json:"email"`
}

type ListUserOptions struct {
	ListOptions
	Email string `url:"email,omitempty" json:"email,omitempty"`
}

// ListUsers gets all users for authorized team
//
// http://api-docs.amixr.io/#list-users
func (service *UserService) ListUsers(opt *ListUserOptions) (*PaginatedUsersResponse, *http.Response, error) {
	u := fmt.Sprintf("%s/", service.url)

	req, err := service.client.NewRequest("GET", u, opt)
	if err != nil {
		return nil, nil, err
	}

	var users *PaginatedUsersResponse
	resp, err := service.client.Do(req, &users)
	if err != nil {
		return nil, resp, err
	}

	return users, resp, err
}

type GetUserOptions struct {
}

// Get user by given id
//
// http://api-docs.amixr.io/#get-user
func (service *UserService) GetUser(id string, opt *GetUserOptions) (*User, *http.Response, error) {
	u := fmt.Sprintf("%s/%s/", service.url, id)

	req, err := service.client.NewRequest("GET", u, opt)
	if err != nil {
		return nil, nil, err
	}

	user := new(User)
	resp, err := service.client.Do(req, user)
	if err != nil {
		return nil, resp, err
	}

	return user, resp, err
}
