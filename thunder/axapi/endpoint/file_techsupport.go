

package endpoint
import (
        "io/ioutil"
        "os"
        "net/http"
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_0-P2_22
type FileTechsupport struct {
	Inst struct {

    Action string `json:"action"`
    File string `json:"file"`
    Slot int `json:"slot"`
    Status FileTechsupportStatus332 `json:"status"`
    Uuid string `json:"uuid"`
    FileContent []byte `json:"-"`
    FileHandle string `json:"file-handle"`

	} `json:"techsupport"`
}
type DeleteFileTechsupport struct {
        Inst struct {
            FileName string `json:"filename"`
        } `json:"techsupport"`
}


type FileTechsupportStatus332 struct {
    Uuid string `json:"uuid"`
}

func (p *FileTechsupport) GetId() string{
    return "1"
}

func (p *FileTechsupport) getPath() string{
    return "file/techsupport"
}

func (p *FileTechsupport) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FileTechsupport::Post")
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
        s := &FileTechsupport{}
            s.Inst.Action = p.Inst.Action
            s.Inst.File = p.Inst.File
            s.Inst.Slot = p.Inst.Slot
            s.Inst.Status = p.Inst.Status
            s.Inst.Uuid = p.Inst.Uuid
            s.Inst.FileContent = data
            s.Inst.FileHandle =  p.Inst.File
        _, err := axapi.NormalizeMultipartObject(http.MethodPost, p.getPath(), s.Inst.File, s.Inst.FileContent, s, headers, host, logger)
    return err
}

func (p *FileTechsupport) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FileTechsupport::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, "file/techsupport/oper", "", nil, headers, logger)
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
func (p *FileTechsupport) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FileTechsupport::Put")
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

func (p *FileTechsupport) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FileTechsupport::Delete")
    headers := axapi.GenRequestHeader(authToken)
        s := &DeleteFileTechsupport{}
        s.Inst.FileName = p.Inst.File
        payloadBytes, err := json.Marshal(s)
        logger.Println(s)
        if err != nil {
            logger.Println("json.Marshal() failed with error", err)
            return err
        }
        logger.Println("payload:", string(payloadBytes))
        _, _, err = axapi.SendPost(host, "delete/techsupport", payloadBytes, headers, logger)
    return err
}
