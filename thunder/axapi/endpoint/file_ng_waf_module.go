

package endpoint
import (
        "io/ioutil"
        "os"
        "net/http"
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_0-P2_22
type FileNgWafModule struct {
	Inst struct {

    Action string `json:"action"`
    File string `json:"file"`
    FileHandle string `json:"file-handle"`
    FileContent []byte `json:"-"`

	} `json:"ng-waf-module"`
}
type DeleteFileNgWafModule struct {
        Inst struct {
            FileName string `json:"filename"`
        } `json:"ng-waf-module"`
}

func (p *FileNgWafModule) GetId() string{
    return "1"
}

func (p *FileNgWafModule) getPath() string{
    return "file/ng-waf-module"
}

func (p *FileNgWafModule) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FileNgWafModule::Post")
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
        s := &FileNgWafModule{}
            s.Inst.Action = p.Inst.Action
            s.Inst.File = p.Inst.File
            s.Inst.FileHandle =  p.Inst.File
            s.Inst.FileContent = data
        _, err := axapi.NormalizeMultipartObject(http.MethodPost, p.getPath(), s.Inst.File, s.Inst.FileContent, s, headers, host, logger)
    return err
}

func (p *FileNgWafModule) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FileNgWafModule::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, "file/ng-waf-module/oper", "", nil, headers, logger)
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
func (p *FileNgWafModule) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FileNgWafModule::Put")
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

func (p *FileNgWafModule) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FileNgWafModule::Delete")
    headers := axapi.GenRequestHeader(authToken)
        s := &DeleteFileNgWafModule{}
        s.Inst.FileName = p.Inst.File
        payloadBytes, err := json.Marshal(s)
        logger.Println(s)
        if err != nil {
            logger.Println("json.Marshal() failed with error", err)
            return err
        }
        logger.Println("payload:", string(payloadBytes))
        _, _, err = axapi.SendPost(host, "delete/ng-waf-module", payloadBytes, headers, logger)
    return err
}
