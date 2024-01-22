

package endpoint
import (
        "io/ioutil"
        "os"
        "net/http"
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type FileLicenseOper struct {
	Inst struct {

    Oper FileLicenseOperOper `json:"oper"`
    FileContent []byte `json:"-"`
    File string `json:"file"`
    FileHandle string `json:"file-handle"`

	} `json:"oper"`
}
type DeleteFileLicenseOper struct {
        Inst struct {
            FileName string `json:"filename"`
        } `json:"oper"`
}


type FileLicenseOperOper struct {
    HostId string `json:"host-id"`
    FeatureList []FileLicenseOperOperFeatureList `json:"feature-list"`
    FileList []FileLicenseOperOperFileList `json:"file-list"`
}


type FileLicenseOperOperFeatureList struct {
    FeatureInstalled string `json:"feature-installed"`
    Version string `json:"version"`
    ExpireDate string `json:"expire-date"`
    Notice string `json:"notice"`
    Temporary string `json:"Temporary"`
    Sn string `json:"SN"`
    Bandwidth string `json:"bandwidth"`
}


type FileLicenseOperOperFileList struct {
    File_name string `json:"file_name"`
}

func (p *FileLicenseOper) GetId() string{
    return "1"
}

func (p *FileLicenseOper) getPath() string{
    return "file/license/oper"
}

func (p *FileLicenseOper) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FileLicenseOper::Post")
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
        s := &FileLicenseOper{}
            s.Inst.Oper = p.Inst.Oper
            s.Inst.FileContent = data
            s.Inst.File = p.Inst.File
            s.Inst.FileHandle =  p.Inst.File
        _, err := axapi.NormalizeMultipartObject(http.MethodPost, p.getPath(), s.Inst.File, s.Inst.FileContent, s, headers, host, logger)
    return err
}

func (p *FileLicenseOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FileLicenseOper::Get")
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
func (p *FileLicenseOper) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FileLicenseOper::Put")
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

func (p *FileLicenseOper) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FileLicenseOper::Delete")
    headers := axapi.GenRequestHeader(authToken)
        s := &DeleteFileLicenseOper{}
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
