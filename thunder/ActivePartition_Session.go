package thunder

import (
	"bytes"
	"errors"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"util"
	go_thunder "github.com/go_thunder/thunder"
)

type GetActivePartitionSessionStructInstance struct {
	PartitionName string `json:"partition-name"`
}

type GetActivePartitionSessionStruct struct {
	PartitionName GetActivePartitionSessionStructInstance `json:"active-partition"`
}

type ActivePartitionSessionInstance struct {
	CurrPartName string `json:"curr_part_name,omitempty"`
	Shared       int    `json:"shared,omitempty"`
}

type ActivePartitionSession struct {
	Shared ActivePartitionSessionInstance `json:"active-partition,omitempty"`
}


func GetActivePartitionSession(id string, host string) (*GetActivePartitionSessionStruct, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetActivePartition")

	resp, err := go_thunder.DoHttp("GET", "https://"+host+"/axapi/v3/active-partition", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m GetActivePartitionSessionStruct
		erro := json.Unmarshal(data, &m)
		logger.Println(string(data))
		logger.Printf("mm--> %v",m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			return &m, nil
		}
	}

}


func PostActivePartitionSession(id string, inst ActivePartitionSession, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostActivePartition")
	partition_name := inst.Shared.CurrPartName
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	_, err = go_thunder.DoHttp("POST", "https://"+host+"/axapi/v3/active-partition", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request For Partition failed with error ", err)
		return err
	}

	datapr, err := GetActivePartitionSession(id, host)
	logger.Println("provider active partition status --> ", (*datapr).PartitionName.PartitionName )
	if err != nil {
		return err
	}
	
	if string(partition_name) == string((*datapr).PartitionName.PartitionName) {
		logger.Println("partition switch to -----> ", string(partition_name))
		return nil
	} else {
		logger.Println("partition switch fail enter valid partition")
		partition_string := fmt.Sprintf("Fail to switch partition %s, Please enter valid partition",partition_name )
		return errors.New(partition_string)
	}

}

func ActivePartitionEnable(p ActivePartitionSession, client Thunder) error {
	logger := util.GetLoggerInstance()
	//var diags diag.Diagnostics
	if client.Host != "" {
		err := PostActivePartitionSession(client.Token, p, client.Host)
		 if err != nil {
			logger.Println("Activating partition fail, please check whether partiton is presesnt or not")
		 	return err
		}

		//return resourceActivePartitionRead(ctx, d, meta)

	}
	return nil
	//return diags
}
