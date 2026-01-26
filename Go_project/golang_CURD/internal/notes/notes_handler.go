package notes

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Handler struct {
	repo *Repo
}

func NewHandler(repo *Repo) *Handler {
	return &Handler{repo: repo}
}
func (h *Handler) CreateNotes(c *gin.Context) {
	var req CreateNoteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid json",
		})
		return
	}
	now := time.Now().UTC()
	note := Note{
		ID:        primitive.NewObjectID(),
		Title:     req.Title,
		Content:   req.Content,
		Pinned:    req.Pinned,
		CreatedAt: now,
		UpdatedAt: now,
	}
	created, err := h.repo.Create(c.Request.Context(), note)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed",
		})
	}
	c.JSON(http.StatusCreated, created)

}

func (h *Handler) ListNotes(c *gin.Context) {
	notes, err := h.repo.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch notes",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"notes": notes,
	})

}

func (h *Handler) ListSingleNote(c *gin.Context) {
	idstr := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(idstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID",
		})
		return
	}
	note, err := h.repo.GetById(c.Request.Context(), objID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "note not found for given id",
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"notes": note,
	})

}

func (h *Handler) UpdateById(c *gin.Context) {
	idstr := c.Param("id")
	objId, err := primitive.ObjectIDFromHex(idstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID",
		})
		return
	}
	var req UpdateNoteReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON format",
		})
		return
	}
	updated, err := h.repo.UpdateByID(c.Request.Context(), objId, req)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "note not found for given id",
			})
			return
		}
	}
	c.JSON(http.StatusOK, updated)
}
