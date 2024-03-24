package customTypes

// `bson:"_id"` to convert (map) 'ID' in json response to 'id'
type User struct {
	Id        string `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName string `bson:"firstName" json:"firstName"`
	LastName  string `bson:"lastName" json:"lastName"`
}
