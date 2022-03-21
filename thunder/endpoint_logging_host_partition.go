package thunder
import (
    "bytes"
    "github.com/clarketm/json"
    "io/ioutil"
    "util"
)

//based on ACOS 5_2_1-P3_70
type EndpointLoggingHostPartition struct { Inst LoggingHostPartition `json:"partition"` }
type LoggingHostPartition struct {
    PartitionName string `json:"partition-name,omitempty"`
    Shared int `json:"shared,omitempty"`
    Uuid string `json:"uuid,omitempty"`
}

func (p *EndpointLoggingHostPartition) GetId() string{
    return "1"
}

func PostEndpointLoggingHostPartition(auth_token string, inst EndpointLoggingHostPartition, host string) error {
    logger := util.GetLoggerInstance()
    logger.Println("PostLoggingHostPartition")
    headers := GetReqHeader(auth_token)
    payloadBytes, err := json.Marshal(inst)
    if err != nil {
        logger.Println("json.Marshal() failed with error", err)
        return err
    }
    logger.Println("payload:", string(payloadBytes))
    resp, err := doHttp("POST", "https://" + host + "/axapi/v3/logging/host/partition", bytes.NewReader(payloadBytes), headers)
    if err != nil {
        logger.Println("HTTP request failed with error", err)
    } else {
        logger.Println("HTTP status", resp.Status)
        data, _ := ioutil.ReadAll(resp.Body)
        logger.Println("response:", string(data))
        var m EndpointLoggingHostPartition
        err = json.Unmarshal(data, &m)
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        } else {
            err = checkAxApiResponseStatus("PostLoggingHostPartition", data)
        }
    }
    return err
}

func GetEndpointLoggingHostPartition(auth_token string, host string) (*EndpointLoggingHostPartition, error) {
    logger := util.GetLoggerInstance()
    logger.Println("GetLoggingHostPartition")
    headers := GetReqHeader(auth_token)
    resp, err := doHttp("GET", "https://"+ host + "/axapi/v3/logging/host/partition/", nil, headers)
    if err != nil {
        logger.Println("HTTP request failed with error", err)
    } else {
        logger.Println("HTTP status", resp.Status)
        data, _ := ioutil.ReadAll(resp.Body)
        logger.Println("response:", string(data))
        var m EndpointLoggingHostPartition
        err = json.Unmarshal(data, &m)
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        } else {
            err = checkAxApiResponseStatus("GetLoggingHostPartition", data)
            if err == nil {
                return &m, nil
            }
        }
    }
    return nil, err
}

func PutEndpointLoggingHostPartition(auth_token string, inst EndpointLoggingHostPartition, host string) error {
    logger := util.GetLoggerInstance()
    logger.Println("PutLoggingHostPartition")
    headers := GetReqHeader(auth_token)
    payloadBytes, err := json.Marshal(inst)
    if err != nil {
        logger.Println("json.Marshal() failed with error", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))

    resp, err := doHttp("PUT", "https://" + host + "/axapi/v3/logging/host/partition/", bytes.NewReader(payloadBytes), headers)
    if err != nil {
        logger.Println("HTTP request failed with error", err)
    } else {
        logger.Println("HTTP status", resp.Status)
        data, _ := ioutil.ReadAll(resp.Body)
        logger.Println("response:", string(data))
        var m EndpointLoggingHostPartition
        err = json.Unmarshal(data, &m)
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        } else {
            err = checkAxApiResponseStatus("PutLoggingHostPartition", data)
        }
    }
    return err
}

func DeleteEndpointLoggingHostPartition(auth_token string, host string) error {
    logger := util.GetLoggerInstance()
    logger.Println("DeleteLoggingHostPartition")
    headers := GetReqHeader(auth_token)
    resp, err := doHttp("DELETE", "https://" + host + "/axapi/v3/logging/host/partition/", nil, headers)
    if err != nil {
        logger.Println("HTTP request failed with error", err)
    } else {
        logger.Println("HTTP status", resp.Status)
        data, _ := ioutil.ReadAll(resp.Body)
        logger.Println("response:", data)
        var m EndpointLoggingHostPartition
        err = json.Unmarshal(data, &m)
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        } else {
            err = checkAxApiResponseStatus("DeleteLoggingHostPartition", data)
        }
    }
    return err
}
