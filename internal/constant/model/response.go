package model

type Response struct {
	// OK is only true if the request was successful.
	OK bool `json:"ok"`
	// MetaData contains additional data like filtering, pagination, etc.
	MetaData interface{} `json:"meta_data,omitempty"`
	// Data contains the actual data of the response.
	Data interface{} `json:"data,omitempty"`
	// Error contains the error detail if the request was not successful.
	Error *ErrorResponse `json:"error,omitempty"`
}

type MetaData struct {
	// PageNo is the page number of the response data.
	Page int `json:"page,omitempty"`
	// PerPage is the page of the page.
	PerPage int `json:"page_page,omitempty"`
	// TotalPages is the total count of data for the response.
	TotalPages int `json:"total_pages,omitempty"`
	// Extra contains other response specific data
	TotalCount int         `json:"total_count,omitempty"`
	Extra      interface{} `json:"extra,omitempty"`
	// Sort is the sort order of the response data.
	Sort string `form:"sort" json:"sort,omitempty"`
}

type ErrorResponse struct {
	// Code is the error code. It is not status code
	Code int `json:"code"`
	// Message is the error message.
	Message string `json:"message,omitempty"`
	// Description is the error description.
	Description string `json:"description,omitempty"`
	// StackTrace is the stack trace of the error.
	// It is only returned for debugging
	StackTrace string `json:"stack_trace,omitempty"`
	// FieldError is the error detail for each field, if available that is.
	FieldError []FieldError `json:"field_error,omitempty"`
}

type FieldError struct {
	// Name is the name of the field that caused the error.
	Name string `json:"name"`
	// Description is the error description for this field.
	Description string `json:"description"`
}

func PrepareMeta(page, perpage, totalCount int) *MetaData {
	return &MetaData{
		Page:       page,
		PerPage:    perpage,
		TotalPages: totalCount,
	}
}
