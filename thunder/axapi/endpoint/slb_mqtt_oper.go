

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbMqttOper struct {
    
    Oper SlbMqttOperOper `json:"oper"`

}
type DataSlbMqttOper struct {
    DtSlbMqttOper SlbMqttOper `json:"mqtt"`
}


type SlbMqttOperOper struct {
    MqttCpuList []SlbMqttOperOperMqttCpuList `json:"mqtt-cpu-list"`
    CpuCount int `json:"cpu-count"`
}


type SlbMqttOperOperMqttCpuList struct {
    Recv_mqtt_connect int `json:"recv_mqtt_connect"`
    Recv_mqtt_connack int `json:"recv_mqtt_connack"`
    Recv_mqtt_publish int `json:"recv_mqtt_publish"`
    Recv_mqtt_puback int `json:"recv_mqtt_puback"`
    Recv_mqtt_pubrec int `json:"recv_mqtt_pubrec"`
    Recv_mqtt_pubrel int `json:"recv_mqtt_pubrel"`
    Recv_mqtt_pubcomp int `json:"recv_mqtt_pubcomp"`
    Recv_mqtt_subscribe int `json:"recv_mqtt_subscribe"`
    Recv_mqtt_suback int `json:"recv_mqtt_suback"`
    Recv_mqtt_unsubscribe int `json:"recv_mqtt_unsubscribe"`
    Recv_mqtt_unsuback int `json:"recv_mqtt_unsuback"`
    Recv_mqtt_pingreq int `json:"recv_mqtt_pingreq"`
    Recv_mqtt_pingresp int `json:"recv_mqtt_pingresp"`
    Recv_mqtt_disconnect int `json:"recv_mqtt_disconnect"`
    Recv_mqtt_auth int `json:"recv_mqtt_auth"`
    Recv_mqtt_other int `json:"recv_mqtt_other"`
    Curr_proxy int `json:"curr_proxy"`
    Total_proxy int `json:"total_proxy"`
    Request int `json:"request"`
    Parse_connect_fail int `json:"parse_connect_fail"`
    Parse_publish_fail int `json:"parse_publish_fail"`
    Parse_subscribe_fail int `json:"parse_subscribe_fail"`
    Parse_unsubscribe_fail int `json:"parse_unsubscribe_fail"`
    Tuple_not_linked int `json:"tuple_not_linked"`
    Tuple_already_linked int `json:"tuple_already_linked"`
    Conn_null int `json:"conn_null"`
    Client_id_null int `json:"client_id_null"`
    Session_exist int `json:"session_exist"`
    Insertion_failed int `json:"insertion_failed"`
    Insertion_successful int `json:"insertion_successful"`
}

func (p *SlbMqttOper) GetId() string{
    return "1"
}

func (p *SlbMqttOper) getPath() string{
    return "slb/mqtt/oper"
}

func (p *SlbMqttOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbMqttOper,error) {
logger.Println("SlbMqttOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbMqttOper
    if err == nil {
        if len(axResp) > 0{
        err = json.Unmarshal(axResp, &p)
        }
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }
    return payload,err
}
