package services

import "github.com/sid-shaha/go-library-management-system/internals/models"

type MembershipService interface {
	InitiateMembership(userID int, planID int) (*models.MembershipDetails, error)
	ActivateMembership(userID int, planID int, paymentID int) (*models.MembershipDetails, error)
	GetMembershipDetailsByUserID(userID int) (*[]models.MembershipDetails, error)
	GetCurrentMembership(userID int) (*models.MembershipDetails, error)
	GetMembershipDetailsByID(membershipDetailsID int) (*models.MembershipDetails, error)
	CancelMembership(membershipDetailsID int, cancellationReason string) error
}
