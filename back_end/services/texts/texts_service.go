package texts_service

import (
	"context"
	"log/slog"
	"type_writer_api/providers/texts"
	"type_writer_api/structures"
)

type TextsServiceInterface interface {
	GetTexts(ctx context.Context) ([]*structures.Text, error)
	GetTextByIdOrTitle(ctx context.Context, textId *int, title *string) (*structures.Text, error)
	CreateText(ctx context.Context, textInfo structures.TextReq) (*structures.Text, error)
	UpdateText(ctx context.Context, textInfo structures.TextReq, textId int) (*structures.Text, error)
	DeleteText(ctx context.Context, textId int) (bool, error)
}

type TextsService struct {
	TextsProvider texts_provider.TextsProviderInterface
}

func (t *TextsService) GetTexts(ctx context.Context) ([]*structures.Text, error) {
	var result []*structures.Text

	texts, err := t.TextsProvider.GetTexts(ctx)
	if err != nil {
		return nil, err
	}

	for _, text := range texts {
		result = append(result, text)
	}

	return result, nil
}

func (t *TextsService) GetTextByIdOrTitle(ctx context.Context, textId *int, title *string) (*structures.Text, error) {
	text, err := t.TextsProvider.GetTextByIdOrTitle(ctx, textId, title)
	if err != nil {
		return nil, err
	}

	result := text
	return result, nil
}

func (t *TextsService) CreateText(ctx context.Context, textInfo structures.TextReq) (*structures.Text, error) {
	textToCreate := structures.ConvertRequestToText(&textInfo)

	createdText, err := t.TextsProvider.CreateText(ctx, *textToCreate)
	if err != nil {
		slog.ErrorContext(ctx, "failed to create text", "error", err)
		return nil, err
	}

	result := createdText
	return result, nil
}

func (t *TextsService) UpdateText(ctx context.Context, textInfo structures.TextReq, textId int) (*structures.Text, error) {
	existingText, err := t.TextsProvider.GetTextByIdOrTitle(ctx, &textId, nil)

	if err != nil {
		slog.ErrorContext(ctx, "failed to update text", "error", err)
		return nil, err
	}

	if textInfo.TextType != "" {
		existingText.TextType = textInfo.TextType
	}
	if textInfo.Title != "" {
		existingText.Title = textInfo.Title
	}
	if textInfo.Difficulty != "" {
		existingText.Difficulty = textInfo.Difficulty
	}
	if textInfo.TextBody != "" {
		existingText.TextBody = textInfo.TextBody
		existingText.TextLength = len(textInfo.TextBody)
	}

	updatedText, err := t.TextsProvider.UpdateText(ctx, *existingText)
	if err != nil {
		slog.ErrorContext(ctx, "failed to update text", "error", err)
		return nil, err
	}

	result := updatedText
	return result, nil
}

func (t *TextsService) DeleteText(ctx context.Context, textId int) (bool, error) {
	deleted, err := t.TextsProvider.DeleteText(ctx, textId)
	if err != nil {
		return false, err
	}

	return deleted, nil
}

func NewTextsService(textsProvider texts_provider.TextsProviderInterface) *TextsService {
	return &TextsService{
		TextsProvider: textsProvider,
	}
}
