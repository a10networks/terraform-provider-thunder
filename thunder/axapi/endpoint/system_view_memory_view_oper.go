

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemViewMemoryViewOper struct {
    
    Oper SystemViewMemoryViewOperOper `json:"oper"`

}
type DataSystemViewMemoryViewOper struct {
    DtSystemViewMemoryViewOper SystemViewMemoryViewOper `json:"memory-view"`
}


type SystemViewMemoryViewOperOper struct {
    Total int `json:"Total"`
    SystemMemory []SystemViewMemoryViewOperOperSystemMemory `json:"System-memory"`
    SystemMemoryCounts int `json:"system-memory-counts"`
    AflexMemory []SystemViewMemoryViewOperOperAflexMemory `json:"aFleX-memory"`
    AflexMemoryCounts int `json:"aflex-memory-counts"`
    SslMemory []SystemViewMemoryViewOperOperSslMemory `json:"SSL-memory"`
    SslMemoryCounts int `json:"ssl-memory-counts"`
    N2Memory []SystemViewMemoryViewOperOperN2Memory `json:"N2-memory"`
    N2MemoryCounts int `json:"n2-memory-counts"`
    TcpMemory []SystemViewMemoryViewOperOperTcpMemory `json:"TCP-memory"`
    TcpMemoryCounts int `json:"tcp-memory-counts"`
    Usage string `json:"Usage"`
    Used int `json:"Used"`
    Free int `json:"Free"`
    Shared int `json:"Shared"`
    Buffers int `json:"Buffers"`
    Cached int `json:"Cached"`
}


type SystemViewMemoryViewOperOperSystemMemory struct {
    ObjectSize int `json:"Object-size"`
    Allocated int `json:"Allocated"`
    Max int `json:"Max"`
}


type SystemViewMemoryViewOperOperAflexMemory struct {
    ObjectSize int `json:"Object-size"`
    Allocated int `json:"Allocated"`
    Max int `json:"Max"`
}


type SystemViewMemoryViewOperOperSslMemory struct {
    ObjectSize int `json:"Object-size"`
    Allocated int `json:"Allocated"`
    Max int `json:"Max"`
}


type SystemViewMemoryViewOperOperN2Memory struct {
    ObjectSize int `json:"Object-size"`
    Allocated int `json:"Allocated"`
    Max int `json:"Max"`
}


type SystemViewMemoryViewOperOperTcpMemory struct {
    ObjectSize int `json:"Object-size"`
    Allocated int `json:"Allocated"`
    Max int `json:"Max"`
}

func (p *SystemViewMemoryViewOper) GetId() string{
    return "1"
}

func (p *SystemViewMemoryViewOper) getPath() string{
    return "system-view/memory-view/oper"
}

func (p *SystemViewMemoryViewOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemViewMemoryViewOper,error) {
logger.Println("SystemViewMemoryViewOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemViewMemoryViewOper
    if err == nil {
        if len(axResp) > 0{
        err = json.Unmarshal(axResp, &p)
        }
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }
    return payload,err
}
