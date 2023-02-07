package thunder
import (
    "bytes"
    "github.com/clarketm/json"
    "io/ioutil"
    "util"
)

//based on ACOS 5_2_1-P3_70
type EndpointLoggingConsole struct { Inst LoggingConsole `json:"console"` }
type LoggingConsole struct {
    ConsoleLevelname string `json:"console-levelname,omitempty"`
    Uuid string `json:"uuid,omitempty"`
}

func (p *EndpointLoggingConsole) GetId() string{
    return "1"
}

func PostEndpointLoggingConsole(auth_token string, inst EndpointLoggingConsole, host string) error {
    logger := util.GetLoggerInstance()
    logger.Println("PostLoggingConsole")
    headers := GetReqHeader(auth_token)
    payloadBytes, err := json.Marshal(inst)
    if err != nil {
        logger.Println("json.Marshal() failed with error", err)
        return err
    }
    logger.Println("payload:", string(payloadBytes))
    resp, err := doHttp("POST", "https://" + host + "/axapi/v3/logging/console", bytes.NewReader(payloadBytes), headers)
    if err != nil {
        logger.Println("HTTP request failed with error", err)
    } else {
        logger.Println("HTTP status", resp.Status)
        data, _ := ioutil.ReadAll(resp.Body)
        logger.Println("response:", string(data))
        var m EndpointLoggingConsole
        err = json.Unmarshal(data, &m)
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        } else {
            err = checkAxApiResponseStatus("PostLoggingConsole", data)
        }
    }
    return err
}

func GetEndpointLoggingConsole(auth_token string, host string) (*EndpointLoggingConsole, error) {
    logger := util.GetLoggerInstance()
    logger.Println("GetLoggingConsole")
    headers := GetReqHeader(auth_token)
    resp, err := doHttp("GET", "https://"+ host + "/axapi/v3/logging/console/", nil, headers)
    if err != nil {
        logger.Println("HTTP request failed with error", err)
    } else {
        logger.Println("HTTP status", resp.Status)
        data, _ := ioutil.ReadAll(resp.Body)
        logger.Println("response:", string(data))
        var m EndpointLoggingConsole
        err = json.Unmarshal(data, &m)
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        } else {
            err = checkAxApiResponseStatus("GetLoggingConsole", data)
            if err == nil {
                return &m, nil
            }
        }
    }
    return nil, err
}

func PutEndpointLoggingConsole(auth_token string, inst EndpointLoggingConsole, host string) error {
    logger := util.GetLoggerInstance()
    logger.Println("PutLoggingConsole")
    headers := GetReqHeader(auth_token)
    payloadBytes, err := json.Marshal(inst)
    if err != nil {
        logger.Println("json.Marshal() failed with error", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))

    resp, err := doHttp("PUT", "https://" + host + "/axapi/v3/logging/console/", bytes.NewReader(payloadBytes), headers)
    if err != nil {
        logger.Println("HTTP request failed with error", err)
    } else {
        logger.Println("HTTP status", resp.Status)
        data, _ := ioutil.ReadAll(resp.Body)
        logger.Println("response:", string(data))
        var m EndpointLoggingConsole
        err = json.Unmarshal(data, &m)
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        } else {
            err = checkAxApiResponseStatus("PutLoggingConsole", data)
        }
    }
    return err
}

func DeleteEndpointLoggingConsole(auth_token string, host string) error {
    logger := util.GetLoggerInstance()
    logger.Println("DeleteLoggingConsole")
    headers := GetReqHeader(auth_token)
    resp, err := doHttp("DELETE", "https://" + host + "/axapi/v3/logging/console/", nil, headers)
    if err != nil {
        logger.Println("HTTP request failed with error", err)
    } else {
        logger.Println("HTTP status", resp.Status)
        data, _ := ioutil.ReadAll(resp.Body)
        logger.Println("response:", data)
        var m EndpointLoggingConsole
        err = json.Unmarshal(data, &m)
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        } else {
            err = checkAxApiResponseStatus("DeleteLoggingConsole", data)
        }
    }
    return err
}
