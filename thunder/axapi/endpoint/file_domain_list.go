

package endpoint
import (
        "io/ioutil"
        "os"
        "net/http"
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_0-P2_22
type FileDomainList struct {
	Inst struct {

    Action string `json:"action"`
    DstFile string `json:"dst-file"`
    File string `json:"file"`
    FileHandle string `json:"file-handle"`
    Size int `json:"size"`
    Uuid string `json:"uuid"`
    FileContent []byte `json:"-"`

	} `json:"domain-list"`
}
type DeleteFileDomainList struct {
        Inst struct {
            FileName string `json:"filename"`
        } `json:"domain-list"`
}

func (p *FileDomainList) GetId() string{
    return "1"
}

func (p *FileDomainList) getPath() string{
    return "file/domain-list"
}

func (p *FileDomainList) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FileDomainList::Post")
    headers := axapi.GenRequestHeader(authToken)
        f, error := os.Open(p.Inst.FileHandle)
        if error != nil {
            logger.Println("Failed to open a file: ", error)
            return error
        }
        data, error := ioutil.ReadAll(f)
        defer f.Close()
        if error != nil {
            logger.Println("Failed to read file: ", error)
            return error
        }
        s := &FileDomainList{}
            s.Inst.Action = p.Inst.Action
            s.Inst.DstFile = p.Inst.DstFile
            s.Inst.File = p.Inst.File
            s.Inst.FileHandle =  p.Inst.File
            s.Inst.Size = p.Inst.Size
            s.Inst.Uuid = p.Inst.Uuid
            s.Inst.FileContent = data
        _, err := axapi.NormalizeMultipartObject(http.MethodPost, p.getPath(), s.Inst.File, s.Inst.FileContent, s, headers, host, logger)
    return err
}

func (p *FileDomainList) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FileDomainList::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, "file/domain-list/oper", "", nil, headers, logger)
    if err == nil {
        if len(axResp) != 0 {
        err = json.Unmarshal(axResp, &p)
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
                     }
            }
   }
    return err
}
func (p *FileDomainList) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FileDomainList::Put")
    headers := axapi.GenRequestHeader(authToken)
    payloadBytes, err := axapi.SerializeToJson(p)
    if err != nil {
        logger.Println("Failed to serialize struct as json", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))
    _, _, err = axapi.SendPut(host, p.getPath(), "", payloadBytes, headers, logger)
    return err
}

func (p *FileDomainList) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FileDomainList::Delete")
    headers := axapi.GenRequestHeader(authToken)
        s := &DeleteFileDomainList{}
        s.Inst.FileName = p.Inst.File
        payloadBytes, err := json.Marshal(s)
        logger.Println(s)
        if err != nil {
            logger.Println("json.Marshal() failed with error", err)
            return err
        }
        logger.Println("payload:", string(payloadBytes))
        _, _, err = axapi.SendPost(host, "delete/domain-list", payloadBytes, headers, logger)
    return err
}
