

package endpoint
import (
        "io/ioutil"
        "os"
        "net/http"
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_0-P2_22
type FileWebServiceCertKey struct {
	Inst struct {

    FileHandle string `json:"file-handle"`
    FileContent []byte `json:"-"`
    File string `json:"file"`

	} `json:"web-service-cert-key"`
}
type DeleteFileWebServiceCertKey struct {
        Inst struct {
            FileName string `json:"filename"`
        } `json:"web-service-cert-key"`
}

func (p *FileWebServiceCertKey) GetId() string{
    return "1"
}

func (p *FileWebServiceCertKey) getPath() string{
    return "file/web-service-cert-key"
}

func (p *FileWebServiceCertKey) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FileWebServiceCertKey::Post")
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
        s := &FileWebServiceCertKey{}
            s.Inst.FileHandle =  p.Inst.File
            s.Inst.FileContent = data
            s.Inst.File = p.Inst.File
        _, err := axapi.NormalizeMultipartObject(http.MethodPost, p.getPath(), s.Inst.File, s.Inst.FileContent, s, headers, host, logger)
    return err
}

func (p *FileWebServiceCertKey) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FileWebServiceCertKey::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, "file/web-service-cert-key/oper", "", nil, headers, logger)
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
func (p *FileWebServiceCertKey) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FileWebServiceCertKey::Put")
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

func (p *FileWebServiceCertKey) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FileWebServiceCertKey::Delete")
    headers := axapi.GenRequestHeader(authToken)
        s := &DeleteFileWebServiceCertKey{}
        s.Inst.FileName = p.Inst.File
        payloadBytes, err := json.Marshal(s)
        logger.Println(s)
        if err != nil {
            logger.Println("json.Marshal() failed with error", err)
            return err
        }
        logger.Println("payload:", string(payloadBytes))
        _, _, err = axapi.SendPost(host, "delete/web-service-cert-key", payloadBytes, headers, logger)
    return err
}
