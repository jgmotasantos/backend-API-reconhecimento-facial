package group

import (
	"context"
	"myproject/member"
	"time"
	"errors"

	"github.com/google/uuid"
)

type GroupService struct {
	repo *GroupRepository
}

func NewGroupService(repo *GroupRepository) *GroupService {
	return &GroupService{repo: repo}
}

// GET /grupos

func (s *GroupService) GetGroups (userID string, ctx context.Context) ([]GroupByName, error) {

	groups, err := s.repo.FindAllGroupsByUserID(userID, ctx)

	if err != nil {

		return nil, err
	
	}

	var groupsByName []GroupByName

	for _, g  := range groups {

		groupsByName = append(groupsByName, GroupByName{Name: g.Name})

	}

	return groupsByName, nil

}

// POST /grupos

func (s *GroupService) CreateGroup(ctx context.Context, req *CreateGroupRequest) (*Group, error) {
	
	group := &Group{
		ID:        uuid.New().String(),
		Name:      req.Name,
		CreatedAt: time.Now().Format(time.RFC3339),
		Members:   []member.Member{},
		CreatedBy: req.CreatedBy,

	}

	err := s.repo.CreateGroup(ctx, group)
	if err != nil {
		return nil, err
	}

	return group, nil

}


// GET /grupos/{nome-do-grupo}

func (s *GroupService) GetGroupByName(groupName, userId string, ctx context.Context) (*Group, error) {
	
	group, ok := s.repo.FindOneByNameAndCreator(ctx,groupName, userId)

	if !ok {

		return nil, errors.New("not found")

	}

	return group, nil

}