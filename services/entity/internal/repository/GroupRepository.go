package repository

import (
	"cleanArch/pkg/store/postgres"
	"cleanArch/services/entity/internal/context"
	"cleanArch/services/entity/internal/domain/entity"
	"fmt"
	"github.com/opentracing/opentracing-go"
)

type GroupRepositoryImpl struct {
	conn *postgres.Conn
}

func (g *GroupRepositoryImpl) CreateGroup(ctx context.Context, group *entity.Group) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CreateGroup")
	defer span.Finish()

	sqlStatement := `
        INSERT INTO groups (id, name)
        VALUES ($1, $2)
    `

	_, err := g.conn.Pool.Exec(ctx, sqlStatement, group.ID, group.Name)
	if err != nil {
		return err
	}

	fmt.Println("Group created successfully")
	return nil
}

func (g *GroupRepositoryImpl) ReadGroup(ctx context.Context, id int) (*entity.Group, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ReadGroup")
	defer span.Finish()

	sqlStatement := `
        SELECT id, name
        FROM groups
        WHERE id = $1
    `

	var group entity.Group
	err := g.conn.Pool.QueryRow(ctx, sqlStatement, id).Scan(&group.ID, &group.Name)
	if err != nil {
		return nil, err
	}

	fmt.Println("Group read successfully")
	return &group, nil
}

func (g *GroupRepositoryImpl) AddContactToGroup(ctx context.Context, contactID, groupID int) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "AddContactToGroup")
	defer span.Finish()

	sqlStatement := `
        INSERT INTO group_contacts (group_id, contact_id)
        VALUES ($1, $2)
    `

	_, err := g.conn.Pool.Exec(ctx, sqlStatement, groupID, contactID)
	if err != nil {
		return err
	}

	fmt.Println("Contact added to group successfully")
	return nil
}

type GroupRepository interface {
	CreateGroup(ctx context.Context, group *entity.Group) error
	ReadGroup(ctx context.Context, id int) (*entity.Group, error)
	AddContactToGroup(ctx context.Context, contactID, groupID int) error
}

func NewGroupRepository(conn *postgres.Conn) *GroupRepositoryImpl {
	return &GroupRepositoryImpl{
		conn: conn,
	}
}
