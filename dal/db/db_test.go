package db_test

import (
	"testing"

	"github.com/610masters/Backend/dal/db"
	"github.com/610masters/Backend/dal/model"
)

func TestPutGetUsers(t *testing.T) {
	db.Init()
	u0 := model.User{"testUser1", "testpwd"}
	u1 := model.User{"testUser2", "testpwd"}

	users := []model.User{u0, u1}

	err := db.PutUsers(users)
	if err != nil {
		t.Error(err)
	}

	if db.GetUser("testUser1").Password != u0.Password {
		t.Error("GetUsers error")
	}

	if db.GetUser("testUser2").Password != u1.Password {
		t.Error("GetUsers error")
	}
}

func TestPutGetArticles(t *testing.T) {
	db.Init()
	atc0 := model.Article{0, "title0", "", nil, "2020-12-21", "content0", nil}
	atc1 := model.Article{1, "title0", "", nil, "2020-06-03", "content1", nil}
	articles := []model.Article{atc0, atc1}
	err := db.PutArticles(articles)
	if err != nil {
		t.Error(err)
	}

	testArticles := db.GetArticles(1, 0)
	if len(testArticles) != 1||testArticles[0].Id != 1{
		t.Error("Get one aritical wrong!")
	}

	testArticles = db.GetArticles(-1, 2)
	if len(testArticles) != 2||testArticles[0].Id != 0||testArticles[1].Id != 1 {
		t.Error("Get all ariticals wrong!")
	}

	testArticles := db.GetArticles(5, 0)
	if len(articles) != 0 {
		t.Error("get ariticle not exist!")
	}
}
