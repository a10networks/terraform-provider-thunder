

package endpoint
import (
        "io/ioutil"
        "os"
        "net/http"
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type FileTemplateLocal struct {
	Inst struct {

    ActType string `json:"act-type"`
    Action string `json:"action"`
    App FileTemplateApp352 `json:"app"`
    DstFile string `json:"dst-file"`
    File string `json:"file"`
    FileHandle string `json:"file-handle"`
    Uuid string `json:"uuid"`
    FileContent []byte `json:"-"`

	} `json:"template"`
}
type DeleteFileTemplateLocal struct {
        Inst struct {
            FileName string `json:"filename"`
        } `json:"template"`
}


type FileTemplateApp352 struct {
    Action string `json:"action"`
    ActName string `json:"act-name"`
    Version string `json:"version"`
}

func (p *FileTemplateLocal) GetId() string{
    return "1"
}

func (p *FileTemplateLocal) getPath() string{
    return "file/template"
}

func (p *FileTemplateLocal) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FileTemplateLocal::Post")
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
        s := &FileTemplateLocal{}
            s.Inst.ActType = p.Inst.ActType
            s.Inst.Action = p.Inst.Action
            s.Inst.App = p.Inst.App
            s.Inst.DstFile = p.Inst.DstFile
            s.Inst.File = p.Inst.File
            s.Inst.FileHandle =  p.Inst.File
            s.Inst.Uuid = p.Inst.Uuid
            s.Inst.FileContent = data
        _, err := axapi.NormalizeMultipartObject(http.MethodPost, p.getPath(), s.Inst.File, s.Inst.FileContent, s, headers, host, logger)
    return err
}

func (p *FileTemplateLocal) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FileTemplateLocal::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, "file/template/oper", "", nil, headers, logger)
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
func (p *FileTemplateLocal) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FileTemplateLocal::Put")
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

func (p *FileTemplateLocal) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FileTemplateLocal::Delete")
    headers := axapi.GenRequestHeader(authToken)
        s := &DeleteFileTemplateLocal{}
        s.Inst.FileName = p.Inst.File
        payloadBytes, err := json.Marshal(s)
        logger.Println(s)
        if err != nil {
            logger.Println("json.Marshal() failed with error", err)
            return err
        }
        logger.Println("payload:", string(payloadBytes))
        _, _, err = axapi.SendPost(host, "delete/template", payloadBytes, headers, logger)
    return err
}
