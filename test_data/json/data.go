package json

import "encoding/json"

type Person struct {
	Name     string `json:"name,omitempty" idx:"normal"`
	Phone    string `json:"phone,omitempty" idx:"unique"`
	Age      int32  `json:"age,omitempty" idx:"normal"`
	BirthDay int32  `json:"birthDay,omitempty"`
	Gender   uint8  `json:"gender,omitempty"`
}

func (p *Person) Marshal() ([]byte, error) {
	return json.Marshal(p)
}
func (p *Person) Unmarshal(data []byte) error {
	return json.Unmarshal(data, p)
}
func (p *Person) DeepCp() *Person {
	return &Person{}
}
