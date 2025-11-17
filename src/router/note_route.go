package router

import (
	"app/src/controller"
	"app/src/middleware"
	"app/src/service"

	"github.com/gofiber/fiber/v2"
)

func NoteRoutes(r fiber.Router, noteService service.NoteService, userService service.UserService) {
	noteController := controller.NewNoteController(noteService)

	notes := r.Group("/notes", middleware.Auth(userService))

	notes.Post("/", noteController.CreateNote)
	notes.Get("/", noteController.GetNotes)
	notes.Get("/:noteId", noteController.GetNoteByID)
	notes.Delete("/:noteId", noteController.DeleteNote)
}
