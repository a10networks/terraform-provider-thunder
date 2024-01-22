

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type ScaleoutDebugClusterDatabaseOper struct {
    
    Oper ScaleoutDebugClusterDatabaseOperOper `json:"oper"`

}
type DataScaleoutDebugClusterDatabaseOper struct {
    DtScaleoutDebugClusterDatabaseOper ScaleoutDebugClusterDatabaseOper `json:"database"`
}


type ScaleoutDebugClusterDatabaseOperOper struct {
    DatabaseList []ScaleoutDebugClusterDatabaseOperOperDatabaseList `json:"database-list"`
}


type ScaleoutDebugClusterDatabaseOperOperDatabaseList struct {
    Root string `json:"root"`
}

func (p *ScaleoutDebugClusterDatabaseOper) GetId() string{
    return "1"
}

func (p *ScaleoutDebugClusterDatabaseOper) getPath() string{
    return "scaleout/debug/cluster/database/oper"
}

func (p *ScaleoutDebugClusterDatabaseOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataScaleoutDebugClusterDatabaseOper,error) {
logger.Println("ScaleoutDebugClusterDatabaseOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataScaleoutDebugClusterDatabaseOper
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
