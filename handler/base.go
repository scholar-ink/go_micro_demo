package handler

import (
	"encoding/json"
)

func DeepCopy(model interface{}, rpc interface{}) {

	b, _ := json.Marshal(model)

	_ = json.Unmarshal(b, rpc)

}
