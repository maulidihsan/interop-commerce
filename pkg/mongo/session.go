package mongo

import (
	"time"

	"gopkg.in/mgo.v2"
)

type Session struct {
	session *mgo.Session
}

func NewSession(url string, dbAdmin string, username string, password string) (*Session,error) {
	dialInfo := &mgo.DialInfo{
		Addrs: []string{url},
		Timeout:  60 * time.Second,
		Database: dbAdmin,
		Username: username,
		Password: password,
	}

	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		return nil,err
	}
	return &Session{session}, err
}

func(s *Session) Copy() *Session {
	return &Session{s.session.Copy()}
}

func(s *Session) GetCollection(db string, col string) *mgo.Collection {
	return s.session.DB(db).C(col)
}

func(s *Session) Close() {
	if(s.session != nil) {
		s.session.Close()
	}
}