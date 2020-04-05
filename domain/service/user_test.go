package service

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"testing"

	"github.com/google/uuid"
	"github.com/trewanek-org/ddd-practice/domain/entity"
	"github.com/trewanek-org/ddd-practice/domain/value"
	"github.com/trewanek-org/ddd-practice/infrastructure/inmem"
)

func TestUser_Exists(t *testing.T) {
	ctx := context.Background()

	repo := inmem.NewUser()
	srv := NewUser(repo)

	// 1. userがnil
	if err := srv.Exists(ctx, nil); err == nil {
		t.Errorf("expected error received, but error is nil")
		t.Fail()
	}

	// 2. not found なユーザ
	if err := srv.Exists(ctx, &entity.User{}); err == nil {
		t.Errorf("expected error received, but error is nil")
		t.Fail()
	}

	// 3. 正常系
	id := uuid.New().String()
	userID := value.UserID(id)

	hash := md5.Sum([]byte(id))
	userName := value.UserName(hex.EncodeToString(hash[:]))

	user, err := entity.NewUser(userID, userName)
	if err != nil {
		t.Errorf("new user error: %v", err)
		t.Fail()
	}

	if err := repo.Save(ctx, user); err != nil {
		t.Errorf("user repo save error: %v", err)
		t.Fail()
	}

	if err := srv.Exists(ctx, user); err != nil {
		t.Errorf("expected user exist, but user not found: %v", err)
		t.Fail()
	}
}
