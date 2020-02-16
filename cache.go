package main

type Jvalue struct {
	Key     string `json:"Key"`
	Value   []byte `json:"Value"`
	Deltime string `json:"Deltyme"` //*time.Time
}

type Svalue struct {
	Value   []byte
	Deltime string //*time.Time
}
