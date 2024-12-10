package api

import (
	"errors"

	v1 "github.com/Alina9496/documents/pkg/api/v1"
	"github.com/gin-gonic/gin"
)

var (
	errAdminUnauthorized = errors.New("admin unauthorized")
	errInvalidMetaData   = errors.New("invalid meta")
	errInvalidLimit      = errors.New("invalid limit")
)

func (s *Server) errorResponse(c *gin.Context, code int, err error) {
	c.JSON(code, v1.RespError{
		Code: code,
		Text: err.Error(),
	})
}
