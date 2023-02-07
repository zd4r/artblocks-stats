package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/zd4r/artblocks-stats/internal/api/entity"
	"github.com/zd4r/artblocks-stats/internal/api/usecase"
)

type test struct {
	name string
	mock func()
	res  interface{}
	err  error
}

var someErr = errors.New("some error")

func collection(t *testing.T) (*usecase.CollectionUseCase, *MockHoldersRepo, *MockCollectionWebAPI) {
	t.Helper()

	mockCtl := gomock.NewController(t)

	repo := NewMockHoldersRepo(mockCtl)
	webAPI := NewMockCollectionWebAPI(mockCtl)

	collectionUseCase := usecase.NewCollection(repo, webAPI)

	return collectionUseCase, repo, webAPI
}

func TestCollectionUseCase_СalculateStats(t *testing.T) {
	collectionUseCase, _, webAPI := collection(t)

	tests := []test{
		{
			name: "empty result",
			mock: func() {
				c := entity.Collection{}

				webAPI.EXPECT().GetHoldersCount(c).Return(c, nil)
				c.Holders = make([]entity.Holder, c.HoldersCount)

				webAPI.EXPECT().GetHolders(c).Return(c, nil)
			},
			res: entity.Collection{
				Holders: make([]entity.Holder, 0),
			},
			err: nil,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mock()

			res, err := collectionUseCase.СalculateStats(context.Background(), entity.Collection{})

			require.Equal(t, res, testCase.res)
			require.ErrorIs(t, err, testCase.err)
		})
	}
}

func TestCollectionUseCase_GetHolders(t *testing.T) {
	collectionUseCase, _, webAPI := collection(t)

	tests := []test{
		{
			name: "empty result",
			mock: func() {
				c := entity.Collection{}

				webAPI.EXPECT().GetHoldersCount(c).Return(c, nil)
				c.Holders = make([]entity.Holder, c.HoldersCount)

				webAPI.EXPECT().GetHolders(c).Return(c, nil)
			},
			res: entity.Collection{
				Holders: make([]entity.Holder, 0),
			},
			err: nil,
		},
		{
			name: "result with webAPI.GetHoldersCount error",
			mock: func() {
				c := entity.Collection{}

				webAPI.EXPECT().GetHoldersCount(c).Return(c, someErr)
			},
			res: entity.Collection{},
			err: someErr,
		},
		{
			name: "result with webAPI.GetHolders error",
			mock: func() {
				c := entity.Collection{}

				webAPI.EXPECT().GetHoldersCount(c).Return(c, nil)
				c.Holders = make([]entity.Holder, c.HoldersCount)

				webAPI.EXPECT().GetHolders(c).Return(c, someErr)
			},
			res: entity.Collection{},
			err: someErr,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mock()

			res, err := collectionUseCase.GetHolders(context.Background(), entity.Collection{})

			require.Equal(t, res, testCase.res)
			require.ErrorIs(t, err, testCase.err)
		},
		)
	}
}

func TestCollectionUseCase_GatherHoldersScores(t *testing.T) {
	collectionUseCase, repo, webAPI := collection(t)

	tests := []test{
		{
			name: "empty result",
			mock: func() {
				h := entity.Holder{}

				repo.EXPECT().Get(h).Return(h, nil)

				webAPI.EXPECT().GetHolderScores(h).Return(h, nil)

				repo.EXPECT().UpdateScores(h).Return(h, nil)
			},
			res: entity.Collection{
				Holders: make([]entity.Holder, 1),
			},
			err: nil,
		},
		{
			name: "result with repo.Get error",
			mock: func() {
				h := entity.Holder{}

				repo.EXPECT().Get(h).Return(h, someErr)
			},
			res: entity.Collection{},
			err: someErr,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mock()

			c := entity.Collection{}
			c.Holders = make([]entity.Holder, 1)
			res, err := collectionUseCase.GatherHoldersScores(context.Background(), c)

			require.Equal(t, res, testCase.res)
			require.ErrorIs(t, err, testCase.err)
		},
		)
	}
}
