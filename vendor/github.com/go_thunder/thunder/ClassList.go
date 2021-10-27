package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type ClassList struct {
	ClassListInstanceName ClassListInstance `json:"class-list,omitempty"`
}

type ClassListInstance struct {
	ClassListInstanceAcListAcMatchType []ClassListInstanceAcList   `json:"ac-list,omitempty"`
	ClassListInstanceDNSDNSMatchType   []ClassListInstanceDNS      `json:"dns,omitempty"`
	ClassListInstanceFile              int                         `json:"file,omitempty"`
	ClassListInstanceIpv4ListIpv4Addr  []ClassListInstanceIpv4List `json:"ipv4-list,omitempty"`
	ClassListInstanceIpv6ListIpv6Addr  []ClassListInstanceIpv6List `json:"ipv6-list,omitempty"`
	ClassListInstanceName              string                      `json:"name,omitempty"`
	ClassListInstanceStrListStr        []ClassListInstanceStrList  `json:"str-list,omitempty"`
	ClassListInstanceType              string                      `json:"type,omitempty"`
	ClassListInstanceUUID              string                      `json:"uuid,omitempty"`
	ClassListInstanceUserTag           string                      `json:"user-tag,omitempty"`
}

type ClassListInstanceAcList struct {
	ClassListInstanceAcListAcKeyString           string `json:"ac-key-string,omitempty"`
	ClassListInstanceAcListAcMatchType           string `json:"ac-match-type,omitempty"`
	ClassListInstanceAcListAcValue               string `json:"ac-value,omitempty"`
	ClassListInstanceAcListGtpRateLimitPolicyStr string `json:"gtp-rate-limit-policy-str,omitempty"`
}

type ClassListInstanceDNS struct {
	ClassListInstanceDNSDNSGlid                int    `json:"dns-glid,omitempty"`
	ClassListInstanceDNSDNSGlidShared          int    `json:"dns-glid-shared,omitempty"`
	ClassListInstanceDNSDNSLid                 int    `json:"dns-lid,omitempty"`
	ClassListInstanceDNSDNSMatchString         string `json:"dns-match-string,omitempty"`
	ClassListInstanceDNSDNSMatchType           string `json:"dns-match-type,omitempty"`
	ClassListInstanceDNSSharedPartitionDNSGlid int    `json:"shared-partition-dns-glid,omitempty"`
}

type ClassListInstanceIpv4List struct {
	ClassListInstanceIpv4ListAge                  int    `json:"age,omitempty"`
	ClassListInstanceIpv4ListGlid                 int    `json:"glid,omitempty"`
	ClassListInstanceIpv4ListGlidShared           int    `json:"glid-shared,omitempty"`
	ClassListInstanceIpv4ListGtpRateLimitPolicyV4 string `json:"gtp-rate-limit-policy-v4,omitempty"`
	ClassListInstanceIpv4ListIpv4Addr             string `json:"ipv4addr,omitempty"`
	ClassListInstanceIpv4ListLid                  int    `json:"lid,omitempty"`
	ClassListInstanceIpv4ListLsnLid               int    `json:"lsn-lid,omitempty"`
	ClassListInstanceIpv4ListLsnRadiusProfile     int    `json:"lsn-radius-profile,omitempty"`
	ClassListInstanceIpv4ListSharedPartitionGlid  int    `json:"shared-partition-glid,omitempty"`
}

type ClassListInstanceIpv6List struct {
	ClassListInstanceIpv6ListGtpRateLimitPolicyV6  string `json:"gtp-rate-limit-policy-v6,omitempty"`
	ClassListInstanceIpv6ListIpv6Addr              string `json:"ipv6-addr,omitempty"`
	ClassListInstanceIpv6ListSharedPartitionV6Glid int    `json:"shared-partition-v6-glid,omitempty"`
	ClassListInstanceIpv6ListV6Age                 int    `json:"v6-age,omitempty"`
	ClassListInstanceIpv6ListV6Glid                int    `json:"v6-glid,omitempty"`
	ClassListInstanceIpv6ListV6GlidShared          int    `json:"v6-glid-shared,omitempty"`
	ClassListInstanceIpv6ListV6Lid                 int    `json:"v6-lid,omitempty"`
	ClassListInstanceIpv6ListV6LsnLid              int    `json:"v6-lsn-lid,omitempty"`
	ClassListInstanceIpv6ListV6LsnRadiusProfile    int    `json:"v6-lsn-radius-profile,omitempty"`
}

type ClassListInstanceStrList struct {
	ClassListInstanceStrListSharedPartitionStrGlid int    `json:"shared-partition-str-glid,omitempty"`
	ClassListInstanceStrListStr                    string `json:"str,omitempty"`
	ClassListInstanceStrListStrGlid                int    `json:"str-glid,omitempty"`
	ClassListInstanceStrListStrGlidDummy           int    `json:"str-glid-dummy,omitempty"`
	ClassListInstanceStrListStrGlidShared          int    `json:"str-glid-shared,omitempty"`
	ClassListInstanceStrListStrLid                 int    `json:"str-lid,omitempty"`
	ClassListInstanceStrListStrLidDummy            int    `json:"str-lid-dummy,omitempty"`
	ClassListInstanceStrListValueStr               string `json:"value-str,omitempty"`
}

func PostClassList(id string, inst ClassList, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostClassList")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
		return err
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/class-list", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m ClassList
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostClassList", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetClassList(id string, name1 string, host string) (*ClassList, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetClassList")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/class-list/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m ClassList
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetClassList", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutClassList(id string, name1 string, inst ClassList, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutClassList")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
		return err
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/class-list/"+name1, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m ClassList
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Put REQ RES..........................", m)
			err := check_api_status("PutClassList", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func DeleteClassList(id string, name1 string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteClassList")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/class-list/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m ClassList
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Delete REQ RES..........................", m)
			err := check_api_status("DeleteClassList", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}
