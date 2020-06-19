package inmem

import (
	"context"
	"testing"

	"github.com/trewanek-org/ddd-practice/domain/entity"
	"github.com/trewanek-org/ddd-practice/domain/value"
)

func Test_user_Find(t *testing.T) {
	ctx := context.Background()
	name := value.UserName("test_name")
	testUser, err := entity.NewUser("test_id", name)
	if err != nil {
		t.Errorf("entity.NewUser error = %v", err)
		t.Fail()
	}

	UserStore[name] = testUser

	repo := NewUser()
	result, err := repo.Find(ctx, name)
	if err != nil {
		t.Errorf("repo.Find error = %v", err)
		t.Fail()
	}

	if result == nil {
		t.Errorf("user store target user not found in store, name = %v", name)
		t.Fail()
	}

	nfName := value.UserName("not_found_name")
	if _, err := repo.Find(ctx, nfName); err == nil {
		t.Errorf("expected user not found, but user found")
		t.Fail()
	}
}

func Test_user_Save(t *testing.T) {
	ctx := context.Background()
	name := value.UserName("user_name")
	user, err := entity.NewUser("user_id", name)
	if err != nil {
		t.Errorf("entity.NewUser error = %v", err)
		t.Fail()
	}

	repo := NewUser()
	if err := repo.Save(ctx, user); err != nil {
		t.Errorf("repo.Save error = %v", err)
		t.Fail()
	}

	if UserStore[name] == nil {
		t.Errorf("user store target user not found in store, name = %v", name)
		t.Fail()
	}
}
