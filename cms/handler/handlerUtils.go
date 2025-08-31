package handler

import (
	"log"
	"time"
)

const limitPerPage = 10

type Filter struct {
	SearchTerm  string
	PageNumber  int32
	CurrentPage int32
	Offset      int32
	Limit       int32
	SortBy      string
	Order       string
}

func (h *Handler) stringToDateTime(date string) time.Time {
	layout := "2006-01-02T15:04"
	fdate, err := time.Parse(layout, date)
	if err != nil {
		log.Println(err)
	}
	return fdate
}

// func GetFilterData(r *http.Request) *Filter {
// 	var data Filter
// 	queryParams := r.URL.Query()
// 	fmt.Println("queryParams", queryParams)
// 	var err error
// 	data.SearchTerm, err = url.PathUnescape(queryParams.Get("SearchTerm"))
// 	if err != nil {
// 		data.SearchTerm = ""
// 	}
// 	data.SortBy = "created_at"
// 	data.SortBy, err = url.PathUnescape(queryParams.Get("SortBy"))
// 	if err != nil {
// 		data.SortBy = "created_at"
// 	}
// 	data.Order = "ASC"
// 	data.Order, err = url.PathUnescape(queryParams.Get("Order"))
// 	if err != nil {
// 		data.Order = "ASC"
// 	}
// 	page, err := url.PathUnescape(queryParams.Get("page"))
// 	if err != nil {
// 		page = "1"
// 	}
// 	pageNumber, err := strconv.Atoi(page)
// 	if err != nil {
// 		pageNumber = 1
// 	}
// 	data.PageNumber = int32(pageNumber)
// 	var offset int32 = 0
// 	currentPage := pageNumber
// 	if currentPage <= 0 {
// 		currentPage = 1
// 	} else {
// 		offset = limitPerPage*int32(currentPage) - limitPerPage
// 	}
// 	data.CurrentPage = int32(currentPage)
// 	data.Offset = offset
// 	return &data
// }
