

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemFpgaCoreCrcOper struct {
    
    Oper SystemFpgaCoreCrcOperOper `json:"oper"`

}
type DataSystemFpgaCoreCrcOper struct {
    DtSystemFpgaCoreCrcOper SystemFpgaCoreCrcOper `json:"fpga-core-crc"`
}


type SystemFpgaCoreCrcOperOper struct {
    Enable_val string `json:"enable_val"`
    Reboot_val string `json:"reboot_val"`
    Crc_error_present string `json:"crc_error_present"`
}

func (p *SystemFpgaCoreCrcOper) GetId() string{
    return "1"
}

func (p *SystemFpgaCoreCrcOper) getPath() string{
    return "system/fpga-core-crc/oper"
}

func (p *SystemFpgaCoreCrcOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemFpgaCoreCrcOper,error) {
logger.Println("SystemFpgaCoreCrcOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemFpgaCoreCrcOper
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
