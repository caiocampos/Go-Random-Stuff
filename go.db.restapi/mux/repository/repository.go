package repository

import (
	"crypto/tls"
	"log"
	"net"

	mgo "gopkg.in/mgo.v2"

	"go.db.restapi/config"
)

var db *mgo.Database

// connect function establish a connection to database
func connect() {
	if db == nil {
		config.ReadTOML()
		dialInfo, err := mgo.ParseURL(config.TOMLConfig.DB.Server)
		tlsConfig := &tls.Config{}
		dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
			conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
			return conn, err
		}
		session, err := mgo.DialWithInfo(dialInfo)
		if err != nil {
			log.Fatal(err)
		}
		db = session.DB(config.TOMLConfig.DB.Database)
	}
}
