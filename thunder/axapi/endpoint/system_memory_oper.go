package endpoint

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	"github.com/clarketm/json"
)

// based on ACOS 6_0_0-P1_10
type SystemMemoryOper struct {
	Oper SystemMemoryOperOper `json:"oper"`
}

type SystemMemoryy struct {
	DataSystemMemory SystemMemoryOper `json:"memory"`
}

type SystemMemoryOperOper struct {
	Total              int                                `json:"Total"`
	SystemMemory       []SystemMemoryOperOperSystemMemory `json:"System-memory"`
	SystemMemoryCounts int                                `json:"system-memory-counts"`
	AflexMemory        []SystemMemoryOperOperAflexMemory  `json:"aFleX-memory"`
	AflexMemoryCounts  int                                `json:"aflex-memory-counts"`
	SslMemory          []SystemMemoryOperOperSslMemory    `json:"SSL-memory"`
	SslMemoryCounts    int                                `json:"ssl-memory-counts"`
	N2Memory           []SystemMemoryOperOperN2Memory     `json:"N2-memory"`
	N2MemoryCounts     int                                `json:"n2-memory-counts"`
	TcpMemory          []SystemMemoryOperOperTcpMemory    `json:"TCP-memory"`
	TcpMemoryCounts    int                                `json:"tcp-memory-counts"`
	Usage              string                             `json:"Usage"`
	Used               int                                `json:"Used"`
	Free               int                                `json:"Free"`
	Shared             int                                `json:"Shared"`
	Buffers            int                                `json:"Buffers"`
	Cached             int                                `json:"Cached"`
}

type SystemMemoryOperOperSystemMemory struct {
	ObjectSize int `json:"Object-size"`
	Allocated  int `json:"Allocated"`
	Max        int `json:"Max"`
}

type SystemMemoryOperOperAflexMemory struct {
	ObjectSize int `json:"Object-size"`
	Allocated  int `json:"Allocated"`
	Max        int `json:"Max"`
}

type SystemMemoryOperOperSslMemory struct {
	ObjectSize int `json:"Object-size"`
	Allocated  int `json:"Allocated"`
	Max        int `json:"Max"`
}

type SystemMemoryOperOperN2Memory struct {
	ObjectSize int `json:"Object-size"`
	Allocated  int `json:"Allocated"`
	Max        int `json:"Max"`
}

type SystemMemoryOperOperTcpMemory struct {
	ObjectSize int `json:"Object-size"`
	Allocated  int `json:"Allocated"`
	Max        int `json:"Max"`
}

func (p *SystemMemoryOper) GetId() string {
	return "1"
}

func (p *SystemMemoryOper) getPath() string {
	return "system/memory/oper"
}

func (p *SystemMemoryOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (SystemMemoryy, error) {
	logger.Println("SystemMemoryOper::Get")
	headers := axapi.GenRequestHeader(authToken)
	_, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)

	var sm SystemMemoryy

	if err == nil {
		err = json.Unmarshal(axResp, &sm)
		if err != nil {
			logger.Println("json.Unmarshal() failed with error", err)
		}
	}
	return sm, err
}
