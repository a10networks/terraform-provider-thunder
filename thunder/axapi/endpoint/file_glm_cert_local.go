

package endpoint
import (
        "io/ioutil"
        "os"
        "net/http"
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type FileGlmCertLocal struct {
	Inst struct {

    Action string `json:"action"`
    Device int `json:"device"`
    File string `json:"file"`
    FileHandle string `json:"file-handle"`
    FileContent []byte `json:"-"`

	} `json:"glm-cert"`
}
type DeleteFileGlmCertLocal struct {
        Inst struct {
            FileName string `json:"filename"`
        } `json:"glm-cert"`
}

func (p *FileGlmCertLocal) GetId() string{
    return "1"
}

func (p *FileGlmCertLocal) getPath() string{
    return "file/glm-cert"
}

func (p *FileGlmCertLocal) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FileGlmCertLocal::Post")
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
        s := &FileGlmCertLocal{}
            s.Inst.Action = p.Inst.Action
            s.Inst.Device = p.Inst.Device
            s.Inst.File = p.Inst.File
            s.Inst.FileHandle =  p.Inst.File
            s.Inst.FileContent = data
        _, err := axapi.NormalizeMultipartObject(http.MethodPost, p.getPath(), s.Inst.File, s.Inst.FileContent, s, headers, host, logger)
    return err
}

func (p *FileGlmCertLocal) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FileGlmCertLocal::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, "file/glm-cert/oper", "", nil, headers, logger)
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
func (p *FileGlmCertLocal) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FileGlmCertLocal::Put")
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

func (p *FileGlmCertLocal) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FileGlmCertLocal::Delete")
    headers := axapi.GenRequestHeader(authToken)
        s := &DeleteFileGlmCertLocal{}
        s.Inst.FileName = p.Inst.File
        payloadBytes, err := json.Marshal(s)
        logger.Println(s)
        if err != nil {
            logger.Println("json.Marshal() failed with error", err)
            return err
        }
        logger.Println("payload:", string(payloadBytes))
        _, _, err = axapi.SendPost(host, "delete/glm-cert", payloadBytes, headers, logger)
    return err
}
