package notes

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterRoute(r *gin.Engine, db *mongo.Database) {
	//create repo and hanlder onece
	repo := NewRepo(db)
	h := NewHandler(repo)

	notesGrp := r.Group("/notes")
	{
		notesGrp.POST("", h.CreateNotes)
		notesGrp.GET("", h.ListNotes)
		notesGrp.GET("/:id", h.ListSingleNote)
		notesGrp.PUT("/:id", h.UpdateById)
	}

}
