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
	PaginatedResponse
	Users []*User `json:"results"`
}

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	Email    string `json:"email"`
}

type ListUserOptions struct {
	ListOptions
	Username string `url:"username,omitempty" json:"username,omitempty"`
}

// ListUsers gets all users for authorized organization
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
