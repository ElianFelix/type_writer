package texts_provider

import (
	"context"
	"type_writer_api/structures"

	"gorm.io/gorm"
)

type TextsProviderInterface interface {
	GetTexts(ctx context.Context) ([]*structures.Text, error)
	GetTextByIdOrTitle(ctx context.Context, textId *int, title *string) (*structures.Text, error)
	CreateText(ctx context.Context, textInfo structures.Text) (*structures.Text, error)
	UpdateText(ctx context.Context, updatedtextInfo structures.Text) (*structures.Text, error)
	DeleteText(ctx context.Context, textId int) (bool, error)
}

type TextsProvider struct {
	Db *gorm.DB
}

func (t *TextsProvider) GetTexts(ctx context.Context) ([]*structures.Text, error) {
	var texts []*structures.Text
	err := t.Db.WithContext(ctx).Table(structures.TEXT_TABLE_NAME).Find(&texts).Error
	if err != nil {
		return nil, err
	}
	return texts, nil
}

func (t *TextsProvider) GetTextByIdOrTitle(ctx context.Context, textId *int, title *string) (*structures.Text, error) {
	var text *structures.Text
	err := t.Db.WithContext(ctx).Table(structures.TEXT_TABLE_NAME).
		First(&text, "id = ? OR title = ?", textId, title).Error
	if err != nil {
		return nil, err
	}
	return text, nil
}

func (t *TextsProvider) CreateText(ctx context.Context, textInfo structures.Text) (*structures.Text, error) {
	var text *structures.Text
	err := t.Db.WithContext(ctx).Table(structures.TEXT_TABLE_NAME).FirstOrCreate(&text, &textInfo).Error
	if err != nil {
		return nil, err
	}
	return text, nil
}

func (t *TextsProvider) UpdateText(ctx context.Context, updatedTextInfo structures.Text) (*structures.Text, error) {
	err := t.Db.WithContext(ctx).Table(structures.TEXT_TABLE_NAME).Updates(&updatedTextInfo).Error
	if err != nil {
		return nil, err
	}
	return &updatedTextInfo, nil
}

func (t *TextsProvider) DeleteText(ctx context.Context, textId int) (bool, error) {
	var deleteText = structures.Text{Id: textId}
	err := t.Db.WithContext(ctx).Table(structures.TEXT_TABLE_NAME).Delete(&deleteText).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func NewTextsProvider(db *gorm.DB) *TextsProvider {
	return &TextsProvider{
		Db: db,
	}
}
