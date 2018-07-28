package dao

import {
  "log"
  mgo "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
}

type CRMDAO struct{
  Server string
  Database string
}

var db *mgo.Database

const {
  COLLECTION="CRM"
}

func (c *CRMDAO) Connect(){
  session, err := mgo.Dial(c.Server)
  if err != nil {
    log.Fatal(err)
  }
  db = session.DB(c.Database)
}
