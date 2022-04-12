package usecase

import (
	"gomockt/domain"
	"gomockt/model/db/mysql"
	"gomockt/model/db/tx"
	"gomockt/model/repo/mock_repo"
	"gomockt/utils"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestBehaviorInsertBook(t *testing.T) {
	querier := mysql.NewQuerier(BuildTx(true))
	bookInteractor := NewBookInteractor(querier)

	err := bookInteractor.InsertBook(domain.Book{Title: "ducthang"})
	require.NoError(t, err)
}

func TestInsertBook(t *testing.T) {
	testCase := []struct {
		name       string
		title      string
		buildStubs func(mockQuerier *mock_repo.MockBookRepository)
		checkResp  func(t *testing.T, err error)
	}{
		{
			name:  "ok",
			title: "golang",
			buildStubs: func(mockQuerier *mock_repo.MockBookRepository) {
				mockQuerier.EXPECT().
					QueryBook(gomock.Eq("golang")).
					Times(1).
					Return(nil, nil)

				mockQuerier.EXPECT().
					InsertBook(gomock.Any()).
					Times(1).
					Return(nil)
			},
			checkResp: func(t *testing.T, err error) {
				require.NoError(t, err)
			},
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockQuerier := mock_repo.NewMockBookRepository(ctrl)
			bookInteractor := NewBookInteractor(mockQuerier)

			tc.buildStubs(mockQuerier)

			err := bookInteractor.InsertBook(domain.Book{Title: tc.title})
			tc.checkResp(t, err)
		})
	}
}

func BuildTx(flag bool) tx.DBTX {
	db, err := utils.OpenConnection()
	if err != nil {
		panic(err)
	}

	if flag {
		txs, err := db.Begin()
		if err != nil {
			panic(err)
		}

		return &tx.TxConn{
			DB: txs,
		}
	}

	return &tx.DbConn{
		DB: db,
	}
}
