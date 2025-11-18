package controller

import (
	"app/src/model"
	"app/src/service"
	"app/src/validation"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type NoteController struct {
	NoteService service.NoteService
}

func NewNoteController(noteService service.NoteService) *NoteController {
	return &NoteController{
		NoteService: noteService,
	}
}

func getUserFromCtx(c *fiber.Ctx) (*model.User, error) {
	u, ok := c.Locals("user").(*model.User)
	if !ok || u == nil {
		return nil, fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}
	return u, nil
}

// POST /v1/notes
func (ctl *NoteController) CreateNote(c *fiber.Ctx) error {
    user, err := getUserFromCtx(c)
    if err != nil {
        return err
    }

    var req validation.CreateNote
    if err := c.BodyParser(&req); err != nil {
        return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
    }

    note, err := ctl.NoteService.CreateNote(c, user.ID.String(), req)
    if err != nil {
        return err
    }

    return c.Status(fiber.StatusCreated).JSON(note)
}

// GET /v1/notes
func (ctl *NoteController) GetNotes(c *fiber.Ctx) error {
    user, err := getUserFromCtx(c)
    if err != nil {
        return err
    }

    notes, err := ctl.NoteService.GetNotes(c, user.ID.String())
    if err != nil {
        return err
    }

    return c.JSON(notes)
}

// GET /v1/notes/:noteId
func (ctl *NoteController) GetNoteByID(c *fiber.Ctx) error {
    user, err := getUserFromCtx(c)
    if err != nil {
        return err
    }

    idParam := c.Params("noteId")
    noteID64, err := strconv.ParseUint(idParam, 10, 32)
    if err != nil {
        return fiber.NewError(fiber.StatusBadRequest, "Invalid note id")
    }
    noteID := uint(noteID64)

    note, err := ctl.NoteService.GetNoteByID(c, user.ID.String(), noteID)
    if err != nil {
        return err
    }

    return c.JSON(note)
}

// DELETE /v1/notes/:noteId
func (ctl *NoteController) DeleteNote(c *fiber.Ctx) error {
    user, err := getUserFromCtx(c)
    if err != nil {
        return err
    }

    idParam := c.Params("noteId")
    noteID64, err := strconv.ParseUint(idParam, 10, 32)
    if err != nil {
        return fiber.NewError(fiber.StatusBadRequest, "Invalid note id")
    }
    noteID := uint(noteID64)

    if err := ctl.NoteService.DeleteNote(c, user.ID.String(), noteID); err != nil {
        return err
    }

    return c.SendStatus(fiber.StatusNoContent)
}
