package thunder

import (
	"bytes"
	"encoding/json"
	go_thunder "github.com/go_thunder/thunder"
	"util"
)

type ActivePartitionInstance struct {
	CurrPartName string `json:"curr_part_name,omitempty"`
	Shared       int    `json:"shared,omitempty"`
}

type ActivePartition struct {
	Shared ActivePartitionInstance `json:"active-partition,omitempty"`
}

func PostActivePartition(id string, inst ActivePartition, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostActivePartition")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	_, err = go_thunder.DoHttp("POST", "https://"+host+"/axapi/v3/active-partition", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request For Partition failed with error ", err)

	}

}

func ActivePartitionEnable(p ActivePartition, client Thunder) error {
	logger := util.GetLoggerInstance()

	if client.Host != "" {
		logger.Println("[INFO] Creating ActivePartition (Inside resourceActivePartitionCreate) ")

		logger.Println("[INFO] received formatted data from method data to ActivePartition --")
		PostActivePartition(client.Token, p, client.Host)

		//return resourceActivePartitionRead(d, meta)

	}
	return nil
}
