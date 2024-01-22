

package endpoint
import (
        "io/ioutil"
        "os"
        "net/http"
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type FileStartupConfigOper struct {
	Inst struct {

    Oper FileStartupConfigOperOper `json:"oper"`
    FileContent []byte `json:"-"`
    File string `json:"file"`
    FileHandle string `json:"file-handle"`

	} `json:"oper"`
}
type DeleteFileStartupConfigOper struct {
        Inst struct {
            FileName string `json:"filename"`
        } `json:"oper"`
}


type FileStartupConfigOperOper struct {
    Dirty int `json:"dirty"`
    CurrentStartupConfig string `json:"current-startup-config"`
    PriStartupConfig string `json:"pri-startup-config"`
    SecStartupConfig string `json:"sec-startup-config"`
    FileList []FileStartupConfigOperOperFileList `json:"file-list"`
}


type FileStartupConfigOperOperFileList struct {
    ProfileName string `json:"Profile-Name"`
    Size int `json:"Size"`
    UpdateTime string `json:"update-time"`
}

func (p *FileStartupConfigOper) GetId() string{
    return "1"
}

func (p *FileStartupConfigOper) getPath() string{
    return "file/startup-config/oper"
}

func (p *FileStartupConfigOper) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FileStartupConfigOper::Post")
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
        s := &FileStartupConfigOper{}
            s.Inst.Oper = p.Inst.Oper
            s.Inst.FileContent = data
            s.Inst.File = p.Inst.File
            s.Inst.FileHandle =  p.Inst.File
        _, err := axapi.NormalizeMultipartObject(http.MethodPost, p.getPath(), s.Inst.File, s.Inst.FileContent, s, headers, host, logger)
    return err
}

func (p *FileStartupConfigOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FileStartupConfigOper::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, "file/oper/oper", "", nil, headers, logger)
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
func (p *FileStartupConfigOper) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FileStartupConfigOper::Put")
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

func (p *FileStartupConfigOper) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FileStartupConfigOper::Delete")
    headers := axapi.GenRequestHeader(authToken)
        s := &DeleteFileStartupConfigOper{}
        s.Inst.FileName = p.Inst.File
        payloadBytes, err := json.Marshal(s)
        logger.Println(s)
        if err != nil {
            logger.Println("json.Marshal() failed with error", err)
            return err
        }
        logger.Println("payload:", string(payloadBytes))
        _, _, err = axapi.SendPost(host, "delete/oper", payloadBytes, headers, logger)
    return err
}
