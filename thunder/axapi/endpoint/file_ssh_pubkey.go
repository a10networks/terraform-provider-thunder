

package endpoint
import (
        "io/ioutil"
        "os"
        "net/http"
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_0-P2_22
type FileSshPubkey struct {
	Inst struct {

    FileHandle string `json:"file-handle"`
    User string `json:"user"`
    Uuid string `json:"uuid"`
    FileContent []byte `json:"-"`
    File string `json:"file"`

	} `json:"ssh-pubkey"`
}
type DeleteFileSshPubkey struct {
        Inst struct {
            FileName string `json:"filename"`
        } `json:"ssh-pubkey"`
}

func (p *FileSshPubkey) GetId() string{
    return "1"
}

func (p *FileSshPubkey) getPath() string{
    return "file/ssh-pubkey"
}

func (p *FileSshPubkey) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FileSshPubkey::Post")
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
        s := &FileSshPubkey{}
            s.Inst.FileHandle =  p.Inst.File
            s.Inst.User = p.Inst.User
            s.Inst.Uuid = p.Inst.Uuid
            s.Inst.FileContent = data
            s.Inst.File = p.Inst.File
        _, err := axapi.NormalizeMultipartObject(http.MethodPost, p.getPath(), s.Inst.File, s.Inst.FileContent, s, headers, host, logger)
    return err
}

func (p *FileSshPubkey) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FileSshPubkey::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, "file/ssh-pubkey/oper", "", nil, headers, logger)
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
func (p *FileSshPubkey) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FileSshPubkey::Put")
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

func (p *FileSshPubkey) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FileSshPubkey::Delete")
    headers := axapi.GenRequestHeader(authToken)
        s := &DeleteFileSshPubkey{}
        s.Inst.FileName = p.Inst.File
        payloadBytes, err := json.Marshal(s)
        logger.Println(s)
        if err != nil {
            logger.Println("json.Marshal() failed with error", err)
            return err
        }
        logger.Println("payload:", string(payloadBytes))
        _, _, err = axapi.SendPost(host, "delete/ssh-pubkey", payloadBytes, headers, logger)
    return err
}
