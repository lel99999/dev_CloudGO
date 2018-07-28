type contact struct{
     ID bson.ObjectId 'bson:"_id" jsone:"id"'
     LastName string 'bson:"lastname" json:"lastname"'
     FirstName string 'bson:"firstname" json:"firstname"'
     Address string 'bson:"address" json:"address"'
     City string 'bson:"city" json:"city"'
     State string 'bson:"state" json:"state"'
     Zip string 'bson:"zip" json:"zip"'
}
