package network

import (
	"cleanArch/services/entity/internal/context"
	"cleanArch/services/entity/internal/domain/entity"
	contact2 "cleanArch/services/entity/internal/useCase/contact"
	"cleanArch/services/entity/internal/useCase/group"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"log"
	"net/http"
	"strconv"
)

type ContactDelivery struct {
	useCase contact2.ContactUseCase
}

type GroupDelivery struct {
	useCase group.GroupUseCase
}

func NewContactDelivery(useCase contact2.ContactUseCase) *ContactDelivery {
	return &ContactDelivery{
		useCase: useCase,
	}
}

func NewGroupDelivery(useCase group.GroupUseCase) *GroupDelivery {
	return &GroupDelivery{
		useCase: useCase,
	}
}

func (d *ContactDelivery) HandleRequests(w http.ResponseWriter, r *http.Request) {

	newUUID := uuid.New()
	ctx := context.WithValue(context.Background(), "ID", newUUID)

	switch r.Method {
	case http.MethodGet:
		d.handleGetContact(ctx, w, r)

	case http.MethodPost:
		d.handlePostContact(ctx, w, r)

	case http.MethodPut:
		d.handlePutContact(ctx, w, r)

	case http.MethodDelete:
		d.handleDeleteContact(ctx, w, r)

	default:
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}
func (d *GroupDelivery) HandleRequests(w http.ResponseWriter, r *http.Request) {

	newUUID := uuid.New()
	ctx := context.WithValue(context.Background(), "ID", newUUID)

	switch r.Method {
	case http.MethodGet:
		d.handleGetGroup(ctx, w, r)

	case http.MethodPost:
		d.handlePostGroup(ctx, w, r)

	case http.MethodPut:
		d.handlePutGroup(ctx, w, r)

	default:
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

func (d *ContactDelivery) handleGetContact(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "handleGetContact")
	defer span.Finish()

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Printf("Error parsing ID: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	contact, err := d.useCase.ReadContact(ctx, id)
	if err != nil {
		log.Printf("Error reading contact: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(contact)
}
func (d *ContactDelivery) handlePostContact(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "handlePostContact")
	defer span.Finish()

	var contact entity.Contact
	err := json.NewDecoder(r.Body).Decode(&contact)
	if err != nil {
		log.Printf("Error decoding contact: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = d.useCase.CreateContact(ctx, &contact)
	if err != nil {
		log.Printf("Error creating contact: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
func (d *ContactDelivery) handlePutContact(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "handlePutContact")
	defer span.Finish()

	var contact entity.Contact
	err := json.NewDecoder(r.Body).Decode(&contact)
	if err != nil {
		log.Printf("Error decoding contact: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = d.useCase.UpdateContact(ctx, &contact)
	if err != nil {
		log.Printf("Error updating contact: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
func (d *ContactDelivery) handleDeleteContact(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "handleDeleteContact")
	defer span.Finish()

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Printf("Error parsing ID: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = d.useCase.DeleteContact(ctx, id)
	if err != nil {
		log.Printf("Error deleting contact: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (d *GroupDelivery) handleGetGroup(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "handleGetGroup")
	defer span.Finish()

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Printf("Error parsing ID: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	group, err := d.useCase.ReadGroup(ctx, id)
	if err != nil {
		log.Printf("Error reading group: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(group)

}
func (d *GroupDelivery) handlePostGroup(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "handlePostGroup")
	defer span.Finish()

	var group entity.Group
	err := json.NewDecoder(r.Body).Decode(&group)
	if err != nil {
		log.Printf("Error decoding group: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = d.useCase.CreateGroup(ctx, &group)
	if err != nil {
		log.Printf("Error creating group: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
func (d *GroupDelivery) handlePutGroup(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "handlePutGroup")
	defer span.Finish()

	contactID, err1 := strconv.Atoi(r.URL.Query().Get("contactID"))
	groupID, err2 := strconv.Atoi(r.URL.Query().Get("groupID"))
	if err1 != nil || err2 != nil {
		log.Printf("Invalid contactID or groupID: %v, %v", err1, err2)
		http.Error(w, "Invalid contactID or groupID", http.StatusBadRequest)
		return
	}
	err := d.useCase.AddContactToGroup(ctx, contactID, groupID)
	if err != nil {
		log.Printf("Error adding contact to group: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Contact added to group successfully")
}
