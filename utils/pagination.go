package utils

import "github.com/gin-gonic/gin"

const (
	DEFAULT_SIZE = 10
)

type Pagination struct {
	Size  int    `json:"size"`
	Page  int    `json:"page"`
	Count int    `json:"count"`
	Sort  string `json:"sort"`
}

func GeneratePaginationFromRequest(c *gin.Context, default_size int) *Pagination {
	size := Atoi(c.Query("size"), default_size)
	page := Atoi(c.Query("page"), 1)
	count := Atoi(c.Query("count"), 0)
	sort := c.Query("sort")
	return &Pagination{size, page, count, sort}
}
