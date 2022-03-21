package thunder
import (
    "bytes"
    "github.com/clarketm/json"
    "io/ioutil"
    "util"
)

//based on ACOS 5_2_1-P3_70
type EndpointLoggingAuditlog struct { Inst LoggingAuditlog `json:"auditlog"` }
type LoggingAuditlog struct {
    AuditFacility string `json:"audit-facility,omitempty"`
    Host4 string `json:"host4,omitempty"`
    Host6 string `json:"host6,omitempty"`
    PartitionName string `json:"partition-name,omitempty"`
    Port int `json:"port,omitempty"`
    Shared int `json:"shared,omitempty"`
    Uuid string `json:"uuid,omitempty"`
}

func (p *EndpointLoggingAuditlog) GetId() string{
    return "1"
}

func PostEndpointLoggingAuditlog(auth_token string, inst EndpointLoggingAuditlog, host string) error {
    logger := util.GetLoggerInstance()
    logger.Println("PostLoggingAuditlog")
    headers := GetReqHeader(auth_token)
    payloadBytes, err := json.Marshal(inst)
    if err != nil {
        logger.Println("json.Marshal() failed with error", err)
        return err
    }
    logger.Println("payload:", string(payloadBytes))
    resp, err := doHttp("POST", "https://" + host + "/axapi/v3/logging/auditlog", bytes.NewReader(payloadBytes), headers)
    if err != nil {
        logger.Println("HTTP request failed with error", err)
    } else {
        logger.Println("HTTP status", resp.Status)
        data, _ := ioutil.ReadAll(resp.Body)
        logger.Println("response:", string(data))
        var m EndpointLoggingAuditlog
        err = json.Unmarshal(data, &m)
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        } else {
            err = checkAxApiResponseStatus("PostLoggingAuditlog", data)
        }
    }
    return err
}

func GetEndpointLoggingAuditlog(auth_token string, host string) (*EndpointLoggingAuditlog, error) {
    logger := util.GetLoggerInstance()
    logger.Println("GetLoggingAuditlog")
    headers := GetReqHeader(auth_token)
    resp, err := doHttp("GET", "https://"+ host + "/axapi/v3/logging/auditlog/", nil, headers)
    if err != nil {
        logger.Println("HTTP request failed with error", err)
    } else {
        logger.Println("HTTP status", resp.Status)
        data, _ := ioutil.ReadAll(resp.Body)
        logger.Println("response:", string(data))
        var m EndpointLoggingAuditlog
        err = json.Unmarshal(data, &m)
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        } else {
            err = checkAxApiResponseStatus("GetLoggingAuditlog", data)
            if err == nil {
                return &m, nil
            }
        }
    }
    return nil, err
}

func PutEndpointLoggingAuditlog(auth_token string, inst EndpointLoggingAuditlog, host string) error {
    logger := util.GetLoggerInstance()
    logger.Println("PutLoggingAuditlog")
    headers := GetReqHeader(auth_token)
    payloadBytes, err := json.Marshal(inst)
    if err != nil {
        logger.Println("json.Marshal() failed with error", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))

    resp, err := doHttp("PUT", "https://" + host + "/axapi/v3/logging/auditlog/", bytes.NewReader(payloadBytes), headers)
    if err != nil {
        logger.Println("HTTP request failed with error", err)
    } else {
        logger.Println("HTTP status", resp.Status)
        data, _ := ioutil.ReadAll(resp.Body)
        logger.Println("response:", string(data))
        var m EndpointLoggingAuditlog
        err = json.Unmarshal(data, &m)
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        } else {
            err = checkAxApiResponseStatus("PutLoggingAuditlog", data)
        }
    }
    return err
}

func DeleteEndpointLoggingAuditlog(auth_token string, host string) error {
    logger := util.GetLoggerInstance()
    logger.Println("DeleteLoggingAuditlog")
    headers := GetReqHeader(auth_token)
    resp, err := doHttp("DELETE", "https://" + host + "/axapi/v3/logging/auditlog/", nil, headers)
    if err != nil {
        logger.Println("HTTP request failed with error", err)
    } else {
        logger.Println("HTTP status", resp.Status)
        data, _ := ioutil.ReadAll(resp.Body)
        logger.Println("response:", data)
        var m EndpointLoggingAuditlog
        err = json.Unmarshal(data, &m)
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        } else {
            err = checkAxApiResponseStatus("DeleteLoggingAuditlog", data)
        }
    }
    return err
}
