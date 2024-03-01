package grpc

import (
	contact2 "cleanArch/services/entity/internal/useCase/contact"
	"cleanArch/services/entity/internal/useCase/group"
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
