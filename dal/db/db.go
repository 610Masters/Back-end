package db

import (
	"log"

	"github.com/610masters/Backend/dal/model"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

const (
	host   = "127.0.0.1:27017"
	dbsource = "Articles"
	username   = "service_test"
	password   = "123456"

	Atccollection = "ArticleData"
	Usrcollection = "Users"
)
var globalS *mgo.Session

// Initialize the database and the function needed
func Init() {
	dialInfo := &mgo.DialInfo{
		Addrs:    []string{host},
		Source:   dbsource,
		Username: username,
		Password: password,
	}
	s, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		log.Fatalln("create session error ", err)
	}
	globalS = s
}

func connect(db, collection string) (*mgo.Session, *mgo.Collection) {
	mgoSess := globalS.Copy()
	mgoCollec := mgoSess.DB(db).C(collection)
	return mgoSess, mgoCollec
}

func Insert(db, collection string, docs ...interface{}) error {
	ms, mc := connect(db, collection)
	defer ms.Close()
	return mc.Insert(docs...)
}

func Find(db, collection string, query, selector, result interface{}) error {
	ms, mc := connect(db, collection)
	defer ms.Close()
	return mc.Find(query).Select(selector).One(result)
}

func FindAll(db, collection string, query, selector, result interface{}) error {
	ms, mc := connect(db, collection)
	defer ms.Close()
	return mc.Find(query).Select(selector).All(result)
}

// PutArticles : put articles to database
func PutArticles(articles []model.Article) error {
	for i := 0; i < len(articles); i++ {
		err := Insert(dbsource, Atccollection, articles[i])
		if err != nil {
			return err
		}
	}
}

// PutUsers : put users to database
func PutUsers(users []model.User) error {
	for i := 0; i < len(articles); i++ {
		err := Insert(dbsource, Usrcollection, users[i])
		if err != nil {
			return err
		}
	}
}

// GetArticles : get articles by id from database
// if id == -1 ,find all articles and only return 'page' articles
func GetArticles(id int64, page int64) []model.Article {
	var articles []model.Article
	if id == -1 {
		err := FindAll(dbsource, Atccollection, nil, nil, &articles)
		if err != nil {
			log.Fatal(err)
		}
		if len(articles) > 5
			articles = articles[:5]
	}
	else {
		err := Find(dbsource, Atccollection, bson.M{"id": id}, nil, &articles)
		if err != nil {
			log.Fatal(err)
		}
	}
	return articles
}

// GetUser : get users by username from database
func GetUser(username string) model.User {
	var result model.User
	err := Find(dbsource, Usrcollection, bson.M{"username": username}, nil, &result)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

