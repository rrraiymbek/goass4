package contact

import (
	"cleanArch/services/entity/internal/context"
	"cleanArch/services/entity/internal/domain/entity"
	"cleanArch/services/entity/internal/repository"
	"fmt"
	"github.com/opentracing/opentracing-go"
)

type ContactUseCase interface {
	CreateContact(ctx context.Context, contact *entity.Contact) error
	ReadContact(ctx context.Context, id int) (*entity.Contact, error)
	UpdateContact(ctx context.Context, contact *entity.Contact) error
	DeleteContact(ctx context.Context, id int) error
}

type ContactUseCaseImpl struct {
	repo repository.ContactRepository
}

func (c ContactUseCaseImpl) CreateContact(ctx context.Context, contact *entity.Contact) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CreateContact")
	defer span.Finish()

	err := c.repo.CreateContact(ctx, contact)
	if err != nil {
		fmt.Println("Error creating contact:", err)
	}
	return err
}

func (c ContactUseCaseImpl) ReadContact(ctx context.Context, id int) (*entity.Contact, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ReadContact")
	defer span.Finish()

	contact, err := c.repo.ReadContact(ctx, id)
	if err != nil {
		fmt.Println("Error reading contact:", err)
	}
	return contact, err
}

func (c ContactUseCaseImpl) UpdateContact(ctx context.Context, contact *entity.Contact) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UpdateContact")
	defer span.Finish()

	err := c.repo.UpdateContact(ctx, contact)
	if err != nil {
		fmt.Println("Error updating contact:", err)
	}
	return err
}

func (c ContactUseCaseImpl) DeleteContact(ctx context.Context, id int) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "DeleteContact")
	defer span.Finish()

	err := c.repo.DeleteContact(ctx, id)
	if err != nil {
		fmt.Println("Error deleting contact:", err)
	}
	return err
}

func NewContactUseCase(repo repository.ContactRepository) *ContactUseCaseImpl {
	return &ContactUseCaseImpl{
		repo: repo,
	}
}
