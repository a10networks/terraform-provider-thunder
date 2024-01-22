

package endpoint
import (
        "io/ioutil"
        "os"
        "net/http"
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_0-P2_22
type FileHealthPostfile struct {
	Inst struct {

    Action string `json:"action"`
    DstFile string `json:"dst-file"`
    File string `json:"file"`
    FileHandle string `json:"file-handle"`
    Uuid string `json:"uuid"`
    FileContent []byte `json:"-"`

	} `json:"health-postfile"`
}
type DeleteFileHealthPostfile struct {
        Inst struct {
            FileName string `json:"filename"`
        } `json:"health-postfile"`
}

func (p *FileHealthPostfile) GetId() string{
    return "1"
}

func (p *FileHealthPostfile) getPath() string{
    return "file/health-postfile"
}

func (p *FileHealthPostfile) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FileHealthPostfile::Post")
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
        s := &FileHealthPostfile{}
            s.Inst.Action = p.Inst.Action
            s.Inst.DstFile = p.Inst.DstFile
            s.Inst.File = p.Inst.File
            s.Inst.FileHandle =  p.Inst.File
            s.Inst.Uuid = p.Inst.Uuid
            s.Inst.FileContent = data
        _, err := axapi.NormalizeMultipartObject(http.MethodPost, p.getPath(), s.Inst.File, s.Inst.FileContent, s, headers, host, logger)
    return err
}

func (p *FileHealthPostfile) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FileHealthPostfile::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, "file/health-postfile/oper", "", nil, headers, logger)
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
func (p *FileHealthPostfile) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FileHealthPostfile::Put")
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

func (p *FileHealthPostfile) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FileHealthPostfile::Delete")
    headers := axapi.GenRequestHeader(authToken)
        s := &DeleteFileHealthPostfile{}
        s.Inst.FileName = p.Inst.File
        payloadBytes, err := json.Marshal(s)
        logger.Println(s)
        if err != nil {
            logger.Println("json.Marshal() failed with error", err)
            return err
        }
        logger.Println("payload:", string(payloadBytes))
        _, _, err = axapi.SendPost(host, "delete/health-postfile", payloadBytes, headers, logger)
    return err
}
