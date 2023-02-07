package thunder
import (
    "bytes"
    "github.com/clarketm/json"
    "io/ioutil"
    "util"
    "net/url"
)

//based on ACOS 5_2_1-P3_70
type EndpointSlbTemplateDnsLogging struct { Inst SlbTemplateDnsLogging `json:"dns-logging"` }
type SlbTemplateDnsLogging struct {
    Disable int `json:"disable,omitempty"`
    DnsLoggingProtocol string `json:"dns-logging-protocol,omitempty"`
    DnsLoggingRequestSection string `json:"dns-logging-request-section,omitempty"`
    DnsLoggingType string `json:"dns-logging-type,omitempty"`
    Name string `json:"name,omitempty"`
    UserTag string `json:"user-tag,omitempty"`
    Uuid string `json:"uuid,omitempty"`
}

func (p *EndpointSlbTemplateDnsLogging) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func PostEndpointSlbTemplateDnsLogging(auth_token string, inst EndpointSlbTemplateDnsLogging, host string) error {
    logger := util.GetLoggerInstance()
    logger.Println("PostSlbTemplateDnsLogging")
    headers := GetReqHeader(auth_token)
    payloadBytes, err := json.Marshal(inst)
    if err != nil {
        logger.Println("json.Marshal() failed with error", err)
        return err
    }
    logger.Println("payload:", string(payloadBytes))
    resp, err := doHttp("POST", "https://" + host + "/axapi/v3/slb/template/dns-logging", bytes.NewReader(payloadBytes), headers)
    if err != nil {
        logger.Println("HTTP request failed with error", err)
    } else {
        logger.Println("HTTP status", resp.Status)
        data, _ := ioutil.ReadAll(resp.Body)
        logger.Println("response:", string(data))
        var m EndpointSlbTemplateDnsLogging
        err = json.Unmarshal(data, &m)
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        } else {
            err = checkAxApiResponseStatus("PostSlbTemplateDnsLogging", data)
        }
    }
    return err
}

func GetEndpointSlbTemplateDnsLogging(auth_token string, host string, instId string) (*EndpointSlbTemplateDnsLogging, error) {
    logger := util.GetLoggerInstance()
    logger.Println("GetSlbTemplateDnsLogging")
    headers := GetReqHeader(auth_token)
    resp, err := doHttp("GET", "https://"+ host + "/axapi/v3/slb/template/dns-logging/" + instId, nil, headers)
    if err != nil {
        logger.Println("HTTP request failed with error", err)
    } else {
        logger.Println("HTTP status", resp.Status)
        data, _ := ioutil.ReadAll(resp.Body)
        logger.Println("response:", string(data))
        var m EndpointSlbTemplateDnsLogging
        err = json.Unmarshal(data, &m)
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        } else {
            err = checkAxApiResponseStatus("GetSlbTemplateDnsLogging", data)
            if err == nil {
                return &m, nil
            }
        }
    }
    return nil, err
}

func PutEndpointSlbTemplateDnsLogging(auth_token string, inst EndpointSlbTemplateDnsLogging, host string) error {
    logger := util.GetLoggerInstance()
    logger.Println("PutSlbTemplateDnsLogging")
    headers := GetReqHeader(auth_token)
    payloadBytes, err := json.Marshal(inst)
    if err != nil {
        logger.Println("json.Marshal() failed with error", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))

    resp, err := doHttp("PUT", "https://" + host + "/axapi/v3/slb/template/dns-logging/" + inst.GetId(), bytes.NewReader(payloadBytes), headers)
    if err != nil {
        logger.Println("HTTP request failed with error", err)
    } else {
        logger.Println("HTTP status", resp.Status)
        data, _ := ioutil.ReadAll(resp.Body)
        logger.Println("response:", string(data))
        var m EndpointSlbTemplateDnsLogging
        err = json.Unmarshal(data, &m)
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        } else {
            err = checkAxApiResponseStatus("PutSlbTemplateDnsLogging", data)
        }
    }
    return err
}

func DeleteEndpointSlbTemplateDnsLogging(auth_token string, host string, instId string) error {
    logger := util.GetLoggerInstance()
    logger.Println("DeleteSlbTemplateDnsLogging")
    headers := GetReqHeader(auth_token)
    resp, err := doHttp("DELETE", "https://" + host + "/axapi/v3/slb/template/dns-logging/" + instId, nil, headers)
    if err != nil {
        logger.Println("HTTP request failed with error", err)
    } else {
        logger.Println("HTTP status", resp.Status)
        data, _ := ioutil.ReadAll(resp.Body)
        logger.Println("response:", data)
        var m EndpointSlbTemplateDnsLogging
        err = json.Unmarshal(data, &m)
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        } else {
            err = checkAxApiResponseStatus("DeleteSlbTemplateDnsLogging", data)
        }
    }
    return err
}
