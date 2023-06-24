package model

import (
	"math"
)

type PaginationResponse struct {
	Data       interface{} `json:"data"`
	Page       int64       `json:"page"`
	Size       int64       `json:"size"`
	TotalPages int64       `json:"total_pages"`
}

func NewPaginationResponse(data interface{}, page, size, dataCount int64) PaginationResponse {
	totalPages := math.Ceil(float64(dataCount) / float64(size))
	return PaginationResponse{
		Data:       data,
		Page:       page,
		Size:       size,
		TotalPages: int64(totalPages),
	}
}

func Offset(page, size int64) int64 {
	offset := (page - 1) * size
	if offset < 0 {
		return 0
	}
	return offset
}
