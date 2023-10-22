package usecase

import (
	"errors"

	variables "github.com/igsr5/github-project-automation"
	"github.com/igsr5/github-project-automation/domain"
)

type reviewPrAutomation struct {
	reviewPrFetcher ReviewPrFetcher
	projectV2Setter ProjectV2Setter
}

// NewReviewPrAutomation is a factory method to create a new instance of ReviewPrAutomation.
func NewReviewPrAutomation(reviewPrFetcher ReviewPrFetcher, projectV2Setter ProjectV2Setter) domain.ReviewPrAutomation {
	return &reviewPrAutomation{
		reviewPrFetcher: reviewPrFetcher,
		projectV2Setter: projectV2Setter,
	}
}

func (a *reviewPrAutomation) SetInProgress() error {
	prs, err := a.reviewPrFetcher.UnReviewedPrs()
	if err != nil {
		return errors.Join(err)
	}

	categoryID := variables.REVIEW_PR_CATEGORY_ID
	statusID := variables.IN_PROGRESS_STATUS_ID

	projectItems := make([]ProjectItem, 0, len(prs))
	for _, i := range prs {
		projectItems = append(projectItems, ProjectItem{URL: i.URL})
	}

	a.projectV2Setter.Set(categoryID, statusID, projectItems)
	return nil
}

func (a *reviewPrAutomation) SetInPending() error {
	commentedPrs, err := a.reviewPrFetcher.CommentedPrs()
	if err != nil {
		return errors.Join(err)
	}

	changesRequestedPrs, err := a.reviewPrFetcher.ChangesRequestedPrs()
	if err != nil {
		return errors.Join(err)
	}

	prs := append(commentedPrs, changesRequestedPrs...)

	categoryID := variables.REVIEW_PR_CATEGORY_ID
	statusID := variables.IN_PENDING_STATUS_ID

	projectItems := make([]ProjectItem, 0, len(prs))
	for _, i := range prs {
		projectItems = append(projectItems, ProjectItem{URL: i.URL})
	}

	a.projectV2Setter.Set(categoryID, statusID, projectItems)
	return nil
}

func (a *reviewPrAutomation) SetComplete() error {
	prs, err := a.reviewPrFetcher.ApprovedPrs()
	if err != nil {
		return errors.Join(err)
	}

	categoryID := variables.REVIEW_PR_CATEGORY_ID
	statusID := variables.COMPLETE_STATUS_ID

	projectItems := make([]ProjectItem, 0, len(prs))
	for _, i := range prs {
		projectItems = append(projectItems, ProjectItem{URL: i.URL})
	}

	a.projectV2Setter.Set(categoryID, statusID, projectItems)

	return nil
}