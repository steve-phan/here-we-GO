package handlejson

import (
	"encoding/json"
	"log"
)

type TMyCar struct {
	Brand string //`json:"Brand"`
	Model string //`json:"Model"`
	Year  int    //`json:"Year"`
}

const myCarsJson = `
[
		{
			"Brand": "VW",
			"Model": "Passat",
			"Year": 2018
		},
		{
			"Brand": "Toyota",
			"Model": "Yarris",
			"Year": 2021
		}
	
]`

func ParseJson() {

	var unMarshalled []TMyCar

	err := json.Unmarshal([]byte(myCarsJson), &unMarshalled)

	if err != nil {
		log.Println("Error unmarshalling json", err)
	}

	log.Printf("unmarshalled, %v", unMarshalled)

}
