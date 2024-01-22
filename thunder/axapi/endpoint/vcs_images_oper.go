

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VcsImagesOper struct {
    
    Oper VcsImagesOperOper `json:"oper"`

}
type DataVcsImagesOper struct {
    DtVcsImagesOper VcsImagesOper `json:"images"`
}


type VcsImagesOperOper struct {
    ImagesList []VcsImagesOperOperImagesList `json:"images-list"`
}


type VcsImagesOperOperImagesList struct {
    Type string `json:"type"`
    ImagesName string `json:"images-name"`
}

func (p *VcsImagesOper) GetId() string{
    return "1"
}

func (p *VcsImagesOper) getPath() string{
    return "vcs/images/oper"
}

func (p *VcsImagesOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataVcsImagesOper,error) {
logger.Println("VcsImagesOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataVcsImagesOper
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
