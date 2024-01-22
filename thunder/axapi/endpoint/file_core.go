

package endpoint
import (
        "io/ioutil"
        "os"
        "net/http"
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_0-P2_22
type FileCore struct {
	Inst struct {

    Action string `json:"action"`
    File string `json:"file"`
    Slot int `json:"slot"`
    FileContent []byte `json:"-"`
    FileHandle string `json:"file-handle"`

	} `json:"core"`
}
type DeleteFileCore struct {
        Inst struct {
            FileName string `json:"filename"`
        } `json:"core"`
}

func (p *FileCore) GetId() string{
    return "1"
}

func (p *FileCore) getPath() string{
    return "file/core"
}

func (p *FileCore) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FileCore::Post")
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
        s := &FileCore{}
            s.Inst.Action = p.Inst.Action
            s.Inst.File = p.Inst.File
            s.Inst.Slot = p.Inst.Slot
            s.Inst.FileContent = data
            s.Inst.FileHandle =  p.Inst.File
        _, err := axapi.NormalizeMultipartObject(http.MethodPost, p.getPath(), s.Inst.File, s.Inst.FileContent, s, headers, host, logger)
    return err
}

func (p *FileCore) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FileCore::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, "file/core/oper", "", nil, headers, logger)
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
func (p *FileCore) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FileCore::Put")
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

func (p *FileCore) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FileCore::Delete")
    headers := axapi.GenRequestHeader(authToken)
        s := &DeleteFileCore{}
        s.Inst.FileName = p.Inst.File
        payloadBytes, err := json.Marshal(s)
        logger.Println(s)
        if err != nil {
            logger.Println("json.Marshal() failed with error", err)
            return err
        }
        logger.Println("payload:", string(payloadBytes))
        _, _, err = axapi.SendPost(host, "delete/core", payloadBytes, headers, logger)
    return err
}
