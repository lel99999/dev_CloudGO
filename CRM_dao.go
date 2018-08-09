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

func (c *CRMDAO) FindAll() ([]Contact, error){
  var contacts []Contact
  err := db.C(COLLECTION).Find(bson.M{}).All(&contacts)
  return contacts, err
}

func (c *CRMDAO) FindById(id string) (Contact, err){
  var contact Contact
  err := db.C(COLLECTION).FindId(bson.ObjectIndex(id)).One(&contact)
  return contact, err
}
func (c *CRMDAO) Insert(contact Contact) error{
  err := db.C(COLLECTION).Insert(&contact)
  return err

}
func (c *CRMDAO) Delete(contact Contact) error{
  err := db.C(COLLECTION).Remove(&contact)
  return err

}
func (c *CRMDAO) Update(contact Contact) error{
  err := db.C(COLLECTION).UpdateId(contact.ID, &contact)
  return err
}

func CreateCRMEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var contact Contact
	if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	contact.ID = bson.NewObjectId()
	if err := dao.Insert(contact); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, contact)
}
