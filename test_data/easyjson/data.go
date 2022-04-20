package easyjson

//easyjson:json
type Person struct {
	Name     string `json:"name,omitempty" idx:"normal"`
	Phone    string `json:"phone,omitempty" idx:"unique"`
	Age      int32  `json:"age,omitempty" idx:"normal"`
	BirthDay int32  `json:"birthDay,omitempty"`
	Gender   uint8  `json:"gender,omitempty"`
}

func (person *Person) Marshal() (dAtA []byte, err error) {
	return person.MarshalJSON()
}

func (person *Person) Unmarshal(dAtA []byte) error {
	return person.UnmarshalJSON(dAtA)
}

func (person *Person) DeepCp() *Person {
	return &Person{}
}
