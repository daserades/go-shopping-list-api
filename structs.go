package main

var (
	userData interface{}
	settings = getSettings()
)

type ErrorResponse struct {
	Status    bool        `json:"status"`
	ErrorCode int         `json:"error_code"`
	ErrorMsg  string      `json:"error_msg"`
	Data      interface{} `json:"data"`
}

type SuccessResponse struct {
	Status bool        `json:"status"`
	Data   interface{} `json:"data"`
}

type UserInListStruct struct {
	UserStruct
	Lists []ListStruct `json:"lists"`
}

type ListStruct struct {
	ID        int     `json:"id"`
	Title     string  `json:"title"`
	UserID    int     `json:"user_id"`
	Status    int     `json:"status"`
	CreatedAt int     `json:"created_at"`
	UpdatedAt int     `json:"updated_at"`
	TotalCost float32 `json:"total_cost"`
}

type ListInItemStruct struct {
	ListStruct
	Items []ItemStruct `json:"items"`
}

type ItemStruct struct {
	ID        int     `json:"id"`
	ListID    int     `json:"list_id"`
	Title     string  `json:"title"`
	Desc      string  `json:"desc"`
	Priority  string  `json:"priority"`
	Cost      float32 `json:"cost"`
	Status    int     `json:"status"`
	CreatedAt int     `json:"created_at"`
	UpdatedAt int     `json:"updated_at"`
}

type UserStruct struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Email    string `json:"email"`
}
