

package endpoint
import (
        "io/ioutil"
        "os"
        "net/http"
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_0-P2_22
type FileAwsAccesskey struct {
	Inst struct {

    FileHandle string `json:"file-handle"`
    User string `json:"user"`
    Uuid string `json:"uuid"`
    FileContent []byte `json:"-"`
    File string `json:"file"`

	} `json:"aws-accesskey"`
}
type DeleteFileAwsAccesskey struct {
        Inst struct {
            FileName string `json:"filename"`
        } `json:"aws-accesskey"`
}

func (p *FileAwsAccesskey) GetId() string{
    return "1"
}

func (p *FileAwsAccesskey) getPath() string{
    return "file/aws-accesskey"
}

func (p *FileAwsAccesskey) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FileAwsAccesskey::Post")
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
        s := &FileAwsAccesskey{}
            s.Inst.FileHandle =  p.Inst.File
            s.Inst.User = p.Inst.User
            s.Inst.Uuid = p.Inst.Uuid
            s.Inst.FileContent = data
            s.Inst.File = p.Inst.File
        _, err := axapi.NormalizeMultipartObject(http.MethodPost, p.getPath(), s.Inst.File, s.Inst.FileContent, s, headers, host, logger)
    return err
}

func (p *FileAwsAccesskey) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FileAwsAccesskey::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, "file/aws-accesskey/oper", "", nil, headers, logger)
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
func (p *FileAwsAccesskey) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FileAwsAccesskey::Put")
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

func (p *FileAwsAccesskey) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FileAwsAccesskey::Delete")
    headers := axapi.GenRequestHeader(authToken)
        s := &DeleteFileAwsAccesskey{}
        s.Inst.FileName = p.Inst.File
        payloadBytes, err := json.Marshal(s)
        logger.Println(s)
        if err != nil {
            logger.Println("json.Marshal() failed with error", err)
            return err
        }
        logger.Println("payload:", string(payloadBytes))
        _, _, err = axapi.SendPost(host, "delete/aws-accesskey", payloadBytes, headers, logger)
    return err
}
