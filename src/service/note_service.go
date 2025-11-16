package service

import (
	"app/src/model"
	"app/src/validation"
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type NoteService interface {
	CreateNote(c *fiber.Ctx, userID string, req validation.CreateNote) (*model.Note, error)
	GetNotes(c *fiber.Ctx, userID string) ([]model.Note, error)
	GetNoteByID(c *fiber.Ctx, userID string, noteID uint) (*model.Note, error)
	DeleteNote(c *fiber.Ctx, userID string, noteID uint) error
}

type noteService struct {
	DB       *gorm.DB
	Validate *validator.Validate
}

func NewNoteService(db *gorm.DB, validate *validator.Validate) NoteService {
	return &noteService{
		DB:       db,
		Validate: validate,
	}
}

func (s *noteService) CreateNote(c *fiber.Ctx, userID string, req validation.CreateNote) (*model.Note, error) {
	// validasi payload
	if err := s.Validate.Struct(&req); err != nil {
		return nil, err
	}

	note := &model.Note{
		UserID:  userID,
		Title:   req.Title,
		Content: req.Content,
	}

	if err := s.DB.WithContext(c.Context()).Create(note).Error; err != nil {
		return nil, err
	}

	return note, nil
}

func (s *noteService) GetNotes(c *fiber.Ctx, userID string) ([]model.Note, error) {
	var notes []model.Note

	if err := s.DB.WithContext(c.Context()).
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&notes).Error; err != nil {
		return nil, err
	}

	return notes, nil
}

func (s *noteService) GetNoteByID(c *fiber.Ctx, userID string, noteID uint) (*model.Note, error) {
	note := new(model.Note)

	err := s.DB.WithContext(c.Context()).
		Where("id = ? AND user_id = ?", noteID, userID).
		First(note).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fiber.NewError(fiber.StatusNotFound, "Note not found")
	}

	if err != nil {
		return nil, err
	}

	return note, nil
}

func (s *noteService) DeleteNote(c *fiber.Ctx, userID string, noteID uint) error {
	// pastikan note milik user ini
	if err := s.DB.WithContext(c.Context()).
		Where("id = ? AND user_id = ?", noteID, userID).
		Delete(&model.Note{}).Error; err != nil {

		return err
	}

	return nil
}
