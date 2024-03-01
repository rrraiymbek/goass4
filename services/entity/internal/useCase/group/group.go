package group

import (
	"cleanArch/services/entity/internal/context"
	"cleanArch/services/entity/internal/domain/entity"
	"cleanArch/services/entity/internal/repository"
	"fmt"
	"github.com/opentracing/opentracing-go"
)

type GroupUseCase interface {
	CreateGroup(ctx context.Context, group *entity.Group) error
	ReadGroup(ctx context.Context, id int) (*entity.Group, error)
	AddContactToGroup(ctx context.Context, contactID, groupID int) error
}

type GroupUseCaseImpl struct {
	repo repository.GroupRepository
}

func (g GroupUseCaseImpl) CreateGroup(ctx context.Context, group *entity.Group) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CreateGroup")
	defer span.Finish()

	err := g.repo.CreateGroup(ctx, group)
	if err != nil {
		fmt.Println("Error creating group:", err)
	}
	return err
}

func (g GroupUseCaseImpl) ReadGroup(ctx context.Context, id int) (*entity.Group, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ReadGroup")
	defer span.Finish()

	group, err := g.repo.ReadGroup(ctx, id)
	if err != nil {
		fmt.Println("Error reading group:", err)
	}
	return group, err
}

func (g GroupUseCaseImpl) AddContactToGroup(ctx context.Context, contactID, groupID int) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "AddContactToGroup")
	defer span.Finish()

	err := g.repo.AddContactToGroup(ctx, contactID, groupID)
	if err != nil {
		fmt.Println("Error adding contact to group:", err)
	}
	return err
}

func NewGroupUseCase(repo repository.GroupRepository) *GroupUseCaseImpl {
	return &GroupUseCaseImpl{
		repo: repo,
	}
}
