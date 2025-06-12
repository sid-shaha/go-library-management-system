package services

import "github.com/sid-shaha/go-library-management-system/internals/models"

type PlanService interface {
	addPlans(plan *models.Plan) (*models.Plan, error)
	updatePlan(planID int, newPlanData *models.Plan) (*models.Plan, error)
	deletePlan(planID int) error
	getAllPans() ([]*models.Plan, error)
	getPlanDetails(planID int) (*models.Plan, error)
}
