

package endpoint
import (
        "io/ioutil"
        "os"
        "net/http"
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type FileGeoLocationOper struct {
	Inst struct {

    Oper FileGeoLocationOperOper `json:"oper"`
    FileContent []byte `json:"-"`
    File string `json:"file"`
    FileHandle string `json:"file-handle"`

	} `json:"oper"`
}
type DeleteFileGeoLocationOper struct {
        Inst struct {
            FileName string `json:"filename"`
        } `json:"oper"`
}


type FileGeoLocationOperOper struct {
    FileList []FileGeoLocationOperOperFileList `json:"file-list"`
}


type FileGeoLocationOperOperFileList struct {
    File string `json:"file"`
}

func (p *FileGeoLocationOper) GetId() string{
    return "1"
}

func (p *FileGeoLocationOper) getPath() string{
    return "file/geo-location/oper"
}

func (p *FileGeoLocationOper) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FileGeoLocationOper::Post")
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
        s := &FileGeoLocationOper{}
            s.Inst.Oper = p.Inst.Oper
            s.Inst.FileContent = data
            s.Inst.File = p.Inst.File
            s.Inst.FileHandle =  p.Inst.File
        _, err := axapi.NormalizeMultipartObject(http.MethodPost, p.getPath(), s.Inst.File, s.Inst.FileContent, s, headers, host, logger)
    return err
}

func (p *FileGeoLocationOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FileGeoLocationOper::Get")
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
func (p *FileGeoLocationOper) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FileGeoLocationOper::Put")
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

func (p *FileGeoLocationOper) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FileGeoLocationOper::Delete")
    headers := axapi.GenRequestHeader(authToken)
        s := &DeleteFileGeoLocationOper{}
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
