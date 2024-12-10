package api

import (
	"encoding/base64"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"net/http"

	"github.com/gin-contrib/cors"

	"github.com/Alina9496/documents/config"
	v1 "github.com/Alina9496/documents/pkg/api/v1"
	"github.com/Alina9496/tool/pkg/logger"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Server struct {
	service Service
	l       *logger.Logger
	admin   string
}

func NewServer(handler *gin.Engine, l *logger.Logger, t Service, cfg *config.Config) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowCredentials = true
	handler.Use(cors.New(corsConfig))

	// K8s probe
	handler.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })

	// Prometheus metrics
	handler.GET("/metrics", gin.WrapH(promhttp.Handler()))
	s := &Server{t, l, cfg.AdminToken}

	h := handler.Group("/api")
	{
		h.POST("/register", s.Registration)
		h.POST("/auth", s.Authentication)
		h.DELETE("/auth/:token", s.LogOut)
		h.POST("/docs", s.Upload)
		h.GET("/docs", s.GetDocuments)
		h.GET("/docs/:id", s.GetDocument)
		h.DELETE("/docs/:id", s.DeleteDocument)
	}
}

func getAdminTokenFromContext(c *gin.Context) string {
	return c.Request.Header.Get("admin_token")
}

func getUserTokenFromContext(c *gin.Context) string {
	return c.Request.Header.Get("token")
}

func (s *Server) Registration(c *gin.Context) {
	if getAdminTokenFromContext(c) != s.admin {
		s.errorResponse(c, errToHttpStatus(errAdminUnauthorized), errAdminUnauthorized)
		return
	}

	login, err := s.service.Registration(c.Request.Context(),
		toDomainUser(v1.User{
			Login:    c.Request.FormValue("login"),
			Password: c.Request.FormValue("pswd"),
		}))
	if err != nil {
		s.errorResponse(c, errToHttpStatus(err), err)
		return
	}

	c.JSON(http.StatusOK, map[string]any{"response": toLoginResp(login)})
}

func (s *Server) Authentication(c *gin.Context) {
	token, err := s.service.Authentication(c.Request.Context(),
		toDomainUser(v1.User{
			Login:    c.Request.FormValue("login"),
			Password: c.Request.FormValue("pswd"),
		}))
	if err != nil {
		s.errorResponse(c, errToHttpStatus(err), err)
		return
	}

	c.JSON(http.StatusOK, map[string]any{"response": toTokenResp(token)})
}

func (s *Server) LogOut(c *gin.Context) {
	token := c.Param("token")
	err := s.service.LogOut(c, token)
	if err != nil {
		s.errorResponse(c, errToHttpStatus(err), err)
		return
	}

	c.JSON(http.StatusOK, map[string]any{"response": toLogOutTokenResp(token)})
}

func (s *Server) Upload(c *gin.Context) {
	documet, err := toFile(c)
	if err != nil {
		s.errorResponse(c, errToHttpStatus(err), err)
		return
	}

	name, err := s.service.Upload(c, documet)
	if err != nil {
		s.errorResponse(c, errToHttpStatus(err), err)
		return
	}

	c.JSON(http.StatusOK, toUploadResponse(name))
}

func (s *Server) GetDocument(c *gin.Context) {
	documentID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		s.errorResponse(c, errToHttpStatus(err), err)
		return
	}
	document, err := s.service.GetDocument(c, documentID, getUserTokenFromContext(c))
	if err != nil {
		s.errorResponse(c, errToHttpStatus(err), err)
		return
	}

	decodedBytes, err := base64.StdEncoding.DecodeString(document.Content)
	if err != nil {
		s.errorResponse(c, errToHttpStatus(err), err)
		return
	}
	c.Data(http.StatusOK, document.Mime, decodedBytes)
}

func (s *Server) GetDocuments(c *gin.Context) {
	filter, err := toGetDocumentsRequest(c)
	if err != nil {
		s.errorResponse(c, errToHttpStatus(err), err)
		return
	}

	documents, err := s.service.GetDocuments(c, filter)
	if err != nil {
		s.errorResponse(c, errToHttpStatus(err), err)
		return
	}

	c.JSON(http.StatusOK, toGetDocumentsResp(documents))
}
func (s *Server) DeleteDocument(c *gin.Context) {
	token := getUserTokenFromContext(c)
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		s.errorResponse(c, errToHttpStatus(err), err)
		return
	}

	id, err = s.service.DeleteDocument(c, id, token)
	if err != nil {
		s.errorResponse(c, errToHttpStatus(err), err)
		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"response": map[string]any{
			id.String(): true,
		}})
}
