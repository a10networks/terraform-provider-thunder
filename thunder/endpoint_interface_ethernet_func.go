package thunder

import (
    "bytes"
    "github.com/clarketm/json"
    "io/ioutil"
    "util"
)

func PostInterfaceEthernet(id string, inst EndpointInterfaceEthernet, host string) error {
    logger := util.GetLoggerInstance()
    var headers = make(map[string]string)
    headers["Accept"] = "application/json"
    headers["Content-Type"] = "application/json"
    headers["Authorization"] = id
    logger.Println("[INFO] Inside PostInterfaceEthernet")
    payloadBytes, err := json.Marshal(inst)
    logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
    if err != nil {
        logger.Println("[INFO] Marshalling failed with error ", err)
        return err
    }

    resp, err := doHttp("POST", "https://"+host+"/axapi/v3/interface/ethernet", bytes.NewReader(payloadBytes), headers)

    if err != nil {
        logger.Println("The HTTP request failed with error ", err)
        return err
    } else {
        data, _ := ioutil.ReadAll(resp.Body)
        var m InterfaceEthernet
        err := json.Unmarshal(data, &m)
        if err != nil {
            logger.Println("Unmarshal error ", err)
            return err
        } else {
            logger.Println("[INFO] Post REQ RES..........................", m)
            err := checkAxApiResponseStatus("PostInterfaceEthernet", data)
            if err != nil {
                return err
            }

        }
    }
    return err
}

func PostInterfaceEthernetIfnum(id string, ifnum string, inst EndpointInterfaceEthernet, host string) error {
    logger := util.GetLoggerInstance()
    var headers = make(map[string]string)
    headers["Accept"] = "application/json"
    headers["Content-Type"] = "application/json"
    headers["Authorization"] = id
    logger.Println("[INFO] Inside PostInterfaceEthernet")
    payloadBytes, err := json.Marshal(inst)
    logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
    if err != nil {
        logger.Println("[INFO] Marshalling failed with error ", err)
        return err
    }

    resp, err := doHttp("POST", "https://"+host+"/axapi/v3/interface/ethernet/"+ifnum, bytes.NewReader(payloadBytes), headers)

    if err != nil {
        logger.Println("The HTTP request failed with error ", err)
        return err
    } else {
        data, _ := ioutil.ReadAll(resp.Body)
        var m InterfaceEthernet
        err := json.Unmarshal(data, &m)
        if err != nil {
            logger.Println("Unmarshal error ", err)
            return err
        } else {
            logger.Println("[INFO] Post REQ RES..........................", m)
            err := checkAxApiResponseStatus("PostInterfaceEthernet", data)
            if err != nil {
                return err
            }

        }
    }
    return err
}

func GetInterfaceEthernet(id string, name1 string, host string) (*InterfaceEthernet, error) {
    logger := util.GetLoggerInstance()
    var headers = make(map[string]string)
    headers["Accept"] = "application/json"
    headers["Content-Type"] = "application/json"
    headers["Authorization"] = id
    logger.Println("[INFO] Inside GetInterfaceEthernet")

    resp, err := doHttp("GET", "https://"+host+"/axapi/v3/interface/ethernet/"+name1, nil, headers)

    if err != nil {
        logger.Println("The HTTP request failed with error ", err)
        return nil, err
    } else {
        data, _ := ioutil.ReadAll(resp.Body)
        var m InterfaceEthernet
        err := json.Unmarshal(data, &m)
        if err != nil {
            logger.Println("Unmarshal error ", err)
            return nil, err
        } else {
            logger.Println("[INFO] Get REQ RES..........................", m)
            err := checkAxApiResponseStatus("GetInterfaceEthernet", data)
            if err != nil {
                return nil, err
            }
            return &m, nil
        }
    }

}

func PutInterfaceEthernet(id string, name1 string, inst EndpointInterfaceEthernet, host string) error {
    logger := util.GetLoggerInstance()
    var headers = make(map[string]string)
    headers["Accept"] = "application/json"
    headers["Content-Type"] = "application/json"
    headers["Authorization"] = id
    logger.Println("[INFO] Inside PutInterfaceEthernet")
    payloadBytes, err := json.Marshal(inst)
    logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
    if err != nil {
        logger.Println("[INFO] Marshalling failed with error ", err)
        return err
    }

    resp, err := doHttp("PUT", "https://"+host+"/axapi/v3/interface/ethernet/"+name1, bytes.NewReader(payloadBytes), headers)

    if err != nil {
        logger.Println("The HTTP request failed with error ", err)
        return err
    } else {
        data, _ := ioutil.ReadAll(resp.Body)
        var m InterfaceEthernet
        err := json.Unmarshal(data, &m)
        if err != nil {
            logger.Println("Unmarshal error ", err)
            return err
        } else {
            logger.Println("[INFO] Put REQ RES..........................", m)
            err := checkAxApiResponseStatus("PutInterfaceEthernet", data)
            if err != nil {
                return err
            }

        }
    }
    return err
}

func DeleteInterfaceEthernet(id string, name1 string, host string) error {
    logger := util.GetLoggerInstance()
    var headers = make(map[string]string)
    headers["Accept"] = "application/json"
    headers["Content-Type"] = "application/json"
    headers["Authorization"] = id
    logger.Println("[INFO] Inside DeleteInterfaceEthernet")

    resp, err := doHttp("DELETE", "https://"+host+"/axapi/v3/interface/ethernet/"+name1, nil, headers)

    if err != nil {
        logger.Println("The HTTP request failed with error ", err)
        return err
    } else {
        data, _ := ioutil.ReadAll(resp.Body)
        var m InterfaceEthernet
        err := json.Unmarshal(data, &m)
        if err != nil {
            logger.Println("Unmarshal error ", err)
            return err
        } else {
            logger.Println("[INFO] Delete REQ RES..........................", m)
            err := checkAxApiResponseStatus("DeleteInterfaceEthernet", data)
            if err != nil {
                return err
            }

        }
    }
    return err
}
