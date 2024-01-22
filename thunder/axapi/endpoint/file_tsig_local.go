

package endpoint
import (
        "io/ioutil"
        "os"
        "net/http"
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type FileTsigLocal struct {
	Inst struct {

    Action string `json:"action"`
    Device int `json:"device"`
    File string `json:"file"`
    FileHandle string `json:"file-handle"`
    Uuid string `json:"uuid"`
    FileContent []byte `json:"-"`

	} `json:"tsig"`
}
type DeleteFileTsigLocal struct {
        Inst struct {
            FileName string `json:"filename"`
        } `json:"tsig"`
}

func (p *FileTsigLocal) GetId() string{
    return "1"
}

func (p *FileTsigLocal) getPath() string{
    return "file/tsig"
}

func (p *FileTsigLocal) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FileTsigLocal::Post")
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
        s := &FileTsigLocal{}
            s.Inst.Action = p.Inst.Action
            s.Inst.Device = p.Inst.Device
            s.Inst.File = p.Inst.File
            s.Inst.FileHandle =  p.Inst.File
            s.Inst.Uuid = p.Inst.Uuid
            s.Inst.FileContent = data
        _, err := axapi.NormalizeMultipartObject(http.MethodPost, p.getPath(), s.Inst.File, s.Inst.FileContent, s, headers, host, logger)
    return err
}

func (p *FileTsigLocal) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FileTsigLocal::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, "file/tsig/oper", "", nil, headers, logger)
    if err == nil {
        if len(axResp) > 0{
        err = json.Unmarshal(axResp, &p)
        }
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }
    return err
}
func (p *FileTsigLocal) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FileTsigLocal::Put")
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

func (p *FileTsigLocal) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FileTsigLocal::Delete")
    headers := axapi.GenRequestHeader(authToken)
        s := &DeleteFileTsigLocal{}
        s.Inst.FileName = p.Inst.File
        payloadBytes, err := json.Marshal(s)
        logger.Println(s)
        if err != nil {
            logger.Println("json.Marshal() failed with error", err)
            return err
        }
        logger.Println("payload:", string(payloadBytes))
        _, _, err = axapi.SendPost(host, "delete/tsig", payloadBytes, headers, logger)
    return err
}
