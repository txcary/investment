package portfolio

import (
	"fmt"
)

var testJsonBytes = []byte(`{"UserName":"MyUser","Signature":"MEUCIDJmafX+XGJV+Ws2jz0lF2YdJLcrEXAw1ZBPB0/+KjJyAiEA1CR3f/pbngSl0P0mqb7McKSbveSsQ1ir5L4ulpKamuw=","EncryptedData":"F4Zw1vYy","Timestamp":"W5D07g==","PublicKey":"BCNhwc+1nmUYLSDJnacQaKQB1YyT26gdwHCZZd1iwsB14rfGvwv9fuAHjyln9Alap2Voxp/rrdiU2QvE8HuMt5s="}`)

func ExampleInstance() {
	obj := Instance()
	err := obj.DelJson(testJsonBytes)
	if err != nil {
		panic(err)
	}
	err = obj.PutJson(testJsonBytes)
	if err != nil {
		panic(err)
	}
	outputJsonBytes, err := obj.GetJson(testJsonBytes)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(outputJsonBytes))
	//output:
	//TODO
}
