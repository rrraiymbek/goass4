package repository

import (
	"cleanArch/pkg/store/postgres"
	"cleanArch/services/entity/internal/context"
	"cleanArch/services/entity/internal/domain/entity"
	"fmt"
	"github.com/opentracing/opentracing-go"
)

type ContactRepositoryImpl struct {
	conn *postgres.Conn
}

func (c *ContactRepositoryImpl) CreateContact(ctx context.Context, contact *entity.Contact) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CreateContact")
	defer span.Finish()

	sqlStatement := `
        INSERT INTO contacts (id, firstName, lastName, middleName, phone)
        VALUES ($1, $2, $3, $4, $5)
    `

	_, err := c.conn.Pool.Exec(ctx, sqlStatement, contact.ID, contact.FirstName, contact.LastName, contact.MiddleName, contact.Phone)
	if err != nil {
		return err
	}

	fmt.Println("Contact created successfully")
	return nil
}

func (c *ContactRepositoryImpl) ReadContact(ctx context.Context, id int) (*entity.Contact, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ReadContact")
	defer span.Finish()

	sqlStatement := `
        SELECT id, firstName, lastName, middleName, phone
        FROM contacts
        WHERE id = $1
    `

	var contact entity.Contact
	err := c.conn.Pool.QueryRow(ctx, sqlStatement, id).Scan(&contact.ID, &contact.FirstName, &contact.LastName, &contact.MiddleName, &contact.Phone)
	if err != nil {
		return nil, err
	}

	fmt.Println("Contact read successfully")
	return &contact, nil
}

func (c *ContactRepositoryImpl) UpdateContact(ctx context.Context, contact *entity.Contact) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UpdateContact")
	defer span.Finish()

	sqlStatement := `
        UPDATE contacts
        SET firstName = $2, lastName = $3, middleName = $4, phone = $5
        WHERE id = $1
    `

	_, err := c.conn.Pool.Exec(ctx, sqlStatement, contact.ID, contact.FirstName, contact.LastName, contact.MiddleName, contact.Phone)
	if err != nil {
		return err
	}

	fmt.Println("Contact updated successfully")
	return nil
}

func (c *ContactRepositoryImpl) DeleteContact(ctx context.Context, id int) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "DeleteContact")
	defer span.Finish()

	sqlStatement := `
        DELETE FROM contacts
        WHERE id = $1
    `

	_, err := c.conn.Pool.Exec(ctx, sqlStatement, id)
	if err != nil {
		return err
	}

	fmt.Println("Contact deleted successfully")
	return nil
}

type ContactRepository interface {
	CreateContact(ctx context.Context, contact *entity.Contact) error
	ReadContact(ctx context.Context, id int) (*entity.Contact, error)
	UpdateContact(ctx context.Context, contact *entity.Contact) error
	DeleteContact(ctx context.Context, id int) error
}

func NewContactRepository(conn *postgres.Conn) *ContactRepositoryImpl {
	return &ContactRepositoryImpl{
		conn: conn,
	}
}
