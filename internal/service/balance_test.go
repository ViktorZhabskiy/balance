package service

import (
	"balance/internal/repository"
	"testing"
)

func TestBalance_PostTransaction(t *testing.T) {

	mockRepo := &repository.BalanceMock{}
	balanceSrv := NewBalance(mockRepo)

	tt := []struct {
		name    string
		userId  string
		wantErr bool
	}{
		{
			name:    "should create balance transaction",
			userId:  "1",
			wantErr: false,
		},
		{
			name:    "should failed create balance transaction",
			userId:  "200",
			wantErr: true,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			err := balanceSrv.PostTransaction(Transaction{
				UserId:     tc.userId,
				Currency:   "USD",
				Amount:     10.110,
				TimePlaced: "24-JAN-20 10:27:44",
				Type:       "deposit",
			})
			if err != nil != tc.wantErr {
				t.Errorf("PostTransaction() expected error %s, actual error %v", err, tc.wantErr)
			}
		})
	}
}
