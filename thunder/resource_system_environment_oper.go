package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemEnvironmentOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_environment_oper`: Operational Status for the object environment\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemEnvironmentOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"physical_temperature": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"physical_temperature2": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"cpu0_temperature": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"cpu1_temperature": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"fan1a_report": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"fan1a_value": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fan1b_report": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"fan1b_value": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fan2a_report": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"fan2a_value": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fan2b_report": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"fan2b_value": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fan3a_report": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"fan3a_value": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fan3b_report": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"fan3b_value": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fan4a_report": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"fan4a_value": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fan4b_report": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"fan4b_value": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fan5a_report": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"fan5a_value": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fan5b_report": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"fan5b_value": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fan6a_report": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"fan6a_value": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fan6b_report": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"fan6b_value": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fan7a_report": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"fan7a_value": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fan7b_report": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"fan7b_value": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fan8a_report": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"fan8a_value": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fan8b_report": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"fan8b_value": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fan9a_report": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"fan9a_value": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fan9b_report": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"fan9b_value": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fan10a_report": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"fan10a_value": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"fan10b_report": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"fan10b_value": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"voltage_label_1": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"voltage_label_2": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"voltage_label_3": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"voltage_label_4": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"voltage_label_5": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"voltage_label_6": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"voltage_label_7": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"voltage_label_8": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"voltage_label_9": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"voltage_label_10": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"voltage_label_11": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"voltage_label_12": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"voltage_label_13": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"voltage_label_14": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"voltage_label_15": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"voltage_label_16": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"voltage_label_17": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"power_unit1": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"power_unit2": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"power_unit3": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"power_unit4": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"pu2_physical_temperature": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"pu2_physical_temperature2": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"pu2_voltage_label_7": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"pu2_voltage_label_8": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"pu2_voltage_label_9": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"pu2_voltage_label_10": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSystemEnvironmentOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemEnvironmentOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemEnvironmentOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemEnvironmentOperOper := setObjectSystemEnvironmentOperOper(res)
		d.Set("oper", SystemEnvironmentOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemEnvironmentOperOper(ret edpt.DataSystemEnvironmentOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"physical_temperature":      ret.DtSystemEnvironmentOper.Oper.PhysicalTemperature,
			"physical_temperature2":     ret.DtSystemEnvironmentOper.Oper.PhysicalTemperature2,
			"cpu0_temperature":          ret.DtSystemEnvironmentOper.Oper.Cpu0Temperature,
			"cpu1_temperature":          ret.DtSystemEnvironmentOper.Oper.Cpu1Temperature,
			"fan1a_report":              ret.DtSystemEnvironmentOper.Oper.Fan1aReport,
			"fan1a_value":               ret.DtSystemEnvironmentOper.Oper.Fan1aValue,
			"fan1b_report":              ret.DtSystemEnvironmentOper.Oper.Fan1bReport,
			"fan1b_value":               ret.DtSystemEnvironmentOper.Oper.Fan1bValue,
			"fan2a_report":              ret.DtSystemEnvironmentOper.Oper.Fan2aReport,
			"fan2a_value":               ret.DtSystemEnvironmentOper.Oper.Fan2aValue,
			"fan2b_report":              ret.DtSystemEnvironmentOper.Oper.Fan2bReport,
			"fan2b_value":               ret.DtSystemEnvironmentOper.Oper.Fan2bValue,
			"fan3a_report":              ret.DtSystemEnvironmentOper.Oper.Fan3aReport,
			"fan3a_value":               ret.DtSystemEnvironmentOper.Oper.Fan3aValue,
			"fan3b_report":              ret.DtSystemEnvironmentOper.Oper.Fan3bReport,
			"fan3b_value":               ret.DtSystemEnvironmentOper.Oper.Fan3bValue,
			"fan4a_report":              ret.DtSystemEnvironmentOper.Oper.Fan4aReport,
			"fan4a_value":               ret.DtSystemEnvironmentOper.Oper.Fan4aValue,
			"fan4b_report":              ret.DtSystemEnvironmentOper.Oper.Fan4bReport,
			"fan4b_value":               ret.DtSystemEnvironmentOper.Oper.Fan4bValue,
			"fan5a_report":              ret.DtSystemEnvironmentOper.Oper.Fan5aReport,
			"fan5a_value":               ret.DtSystemEnvironmentOper.Oper.Fan5aValue,
			"fan5b_report":              ret.DtSystemEnvironmentOper.Oper.Fan5bReport,
			"fan5b_value":               ret.DtSystemEnvironmentOper.Oper.Fan5bValue,
			"fan6a_report":              ret.DtSystemEnvironmentOper.Oper.Fan6aReport,
			"fan6a_value":               ret.DtSystemEnvironmentOper.Oper.Fan6aValue,
			"fan6b_report":              ret.DtSystemEnvironmentOper.Oper.Fan6bReport,
			"fan6b_value":               ret.DtSystemEnvironmentOper.Oper.Fan6bValue,
			"fan7a_report":              ret.DtSystemEnvironmentOper.Oper.Fan7aReport,
			"fan7a_value":               ret.DtSystemEnvironmentOper.Oper.Fan7aValue,
			"fan7b_report":              ret.DtSystemEnvironmentOper.Oper.Fan7bReport,
			"fan7b_value":               ret.DtSystemEnvironmentOper.Oper.Fan7bValue,
			"fan8a_report":              ret.DtSystemEnvironmentOper.Oper.Fan8aReport,
			"fan8a_value":               ret.DtSystemEnvironmentOper.Oper.Fan8aValue,
			"fan8b_report":              ret.DtSystemEnvironmentOper.Oper.Fan8bReport,
			"fan8b_value":               ret.DtSystemEnvironmentOper.Oper.Fan8bValue,
			"fan9a_report":              ret.DtSystemEnvironmentOper.Oper.Fan9aReport,
			"fan9a_value":               ret.DtSystemEnvironmentOper.Oper.Fan9aValue,
			"fan9b_report":              ret.DtSystemEnvironmentOper.Oper.Fan9bReport,
			"fan9b_value":               ret.DtSystemEnvironmentOper.Oper.Fan9bValue,
			"fan10a_report":             ret.DtSystemEnvironmentOper.Oper.Fan10aReport,
			"fan10a_value":              ret.DtSystemEnvironmentOper.Oper.Fan10aValue,
			"fan10b_report":             ret.DtSystemEnvironmentOper.Oper.Fan10bReport,
			"fan10b_value":              ret.DtSystemEnvironmentOper.Oper.Fan10bValue,
			"voltage_label_1":           ret.DtSystemEnvironmentOper.Oper.VoltageLabel1,
			"voltage_label_2":           ret.DtSystemEnvironmentOper.Oper.VoltageLabel2,
			"voltage_label_3":           ret.DtSystemEnvironmentOper.Oper.VoltageLabel3,
			"voltage_label_4":           ret.DtSystemEnvironmentOper.Oper.VoltageLabel4,
			"voltage_label_5":           ret.DtSystemEnvironmentOper.Oper.VoltageLabel5,
			"voltage_label_6":           ret.DtSystemEnvironmentOper.Oper.VoltageLabel6,
			"voltage_label_7":           ret.DtSystemEnvironmentOper.Oper.VoltageLabel7,
			"voltage_label_8":           ret.DtSystemEnvironmentOper.Oper.VoltageLabel8,
			"voltage_label_9":           ret.DtSystemEnvironmentOper.Oper.VoltageLabel9,
			"voltage_label_10":          ret.DtSystemEnvironmentOper.Oper.VoltageLabel10,
			"voltage_label_11":          ret.DtSystemEnvironmentOper.Oper.VoltageLabel11,
			"voltage_label_12":          ret.DtSystemEnvironmentOper.Oper.VoltageLabel12,
			"voltage_label_13":          ret.DtSystemEnvironmentOper.Oper.VoltageLabel13,
			"voltage_label_14":          ret.DtSystemEnvironmentOper.Oper.VoltageLabel14,
			"voltage_label_15":          ret.DtSystemEnvironmentOper.Oper.VoltageLabel15,
			"voltage_label_16":          ret.DtSystemEnvironmentOper.Oper.VoltageLabel16,
			"voltage_label_17":          ret.DtSystemEnvironmentOper.Oper.VoltageLabel17,
			"power_unit1":               ret.DtSystemEnvironmentOper.Oper.PowerUnit1,
			"power_unit2":               ret.DtSystemEnvironmentOper.Oper.PowerUnit2,
			"power_unit3":               ret.DtSystemEnvironmentOper.Oper.PowerUnit3,
			"power_unit4":               ret.DtSystemEnvironmentOper.Oper.PowerUnit4,
			"pu2_physical_temperature":  ret.DtSystemEnvironmentOper.Oper.Pu2PhysicalTemperature,
			"pu2_physical_temperature2": ret.DtSystemEnvironmentOper.Oper.Pu2PhysicalTemperature2,
			"pu2_voltage_label_7":       ret.DtSystemEnvironmentOper.Oper.Pu2VoltageLabel7,
			"pu2_voltage_label_8":       ret.DtSystemEnvironmentOper.Oper.Pu2VoltageLabel8,
			"pu2_voltage_label_9":       ret.DtSystemEnvironmentOper.Oper.Pu2VoltageLabel9,
			"pu2_voltage_label_10":      ret.DtSystemEnvironmentOper.Oper.Pu2VoltageLabel10,
		},
	}
}

func getObjectSystemEnvironmentOperOper(d []interface{}) edpt.SystemEnvironmentOperOper {

	count1 := len(d)
	var ret edpt.SystemEnvironmentOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PhysicalTemperature = in["physical_temperature"].(string)
		ret.PhysicalTemperature2 = in["physical_temperature2"].(string)
		ret.Cpu0Temperature = in["cpu0_temperature"].(string)
		ret.Cpu1Temperature = in["cpu1_temperature"].(string)
		ret.Fan1aReport = in["fan1a_report"].(string)
		ret.Fan1aValue = in["fan1a_value"].(int)
		ret.Fan1bReport = in["fan1b_report"].(string)
		ret.Fan1bValue = in["fan1b_value"].(int)
		ret.Fan2aReport = in["fan2a_report"].(string)
		ret.Fan2aValue = in["fan2a_value"].(int)
		ret.Fan2bReport = in["fan2b_report"].(string)
		ret.Fan2bValue = in["fan2b_value"].(int)
		ret.Fan3aReport = in["fan3a_report"].(string)
		ret.Fan3aValue = in["fan3a_value"].(int)
		ret.Fan3bReport = in["fan3b_report"].(string)
		ret.Fan3bValue = in["fan3b_value"].(int)
		ret.Fan4aReport = in["fan4a_report"].(string)
		ret.Fan4aValue = in["fan4a_value"].(int)
		ret.Fan4bReport = in["fan4b_report"].(string)
		ret.Fan4bValue = in["fan4b_value"].(int)
		ret.Fan5aReport = in["fan5a_report"].(string)
		ret.Fan5aValue = in["fan5a_value"].(int)
		ret.Fan5bReport = in["fan5b_report"].(string)
		ret.Fan5bValue = in["fan5b_value"].(int)
		ret.Fan6aReport = in["fan6a_report"].(string)
		ret.Fan6aValue = in["fan6a_value"].(int)
		ret.Fan6bReport = in["fan6b_report"].(string)
		ret.Fan6bValue = in["fan6b_value"].(int)
		ret.Fan7aReport = in["fan7a_report"].(string)
		ret.Fan7aValue = in["fan7a_value"].(int)
		ret.Fan7bReport = in["fan7b_report"].(string)
		ret.Fan7bValue = in["fan7b_value"].(int)
		ret.Fan8aReport = in["fan8a_report"].(string)
		ret.Fan8aValue = in["fan8a_value"].(int)
		ret.Fan8bReport = in["fan8b_report"].(string)
		ret.Fan8bValue = in["fan8b_value"].(int)
		ret.Fan9aReport = in["fan9a_report"].(string)
		ret.Fan9aValue = in["fan9a_value"].(int)
		ret.Fan9bReport = in["fan9b_report"].(string)
		ret.Fan9bValue = in["fan9b_value"].(int)
		ret.Fan10aReport = in["fan10a_report"].(string)
		ret.Fan10aValue = in["fan10a_value"].(int)
		ret.Fan10bReport = in["fan10b_report"].(string)
		ret.Fan10bValue = in["fan10b_value"].(int)
		ret.VoltageLabel1 = in["voltage_label_1"].(string)
		ret.VoltageLabel2 = in["voltage_label_2"].(string)
		ret.VoltageLabel3 = in["voltage_label_3"].(string)
		ret.VoltageLabel4 = in["voltage_label_4"].(string)
		ret.VoltageLabel5 = in["voltage_label_5"].(string)
		ret.VoltageLabel6 = in["voltage_label_6"].(string)
		ret.VoltageLabel7 = in["voltage_label_7"].(string)
		ret.VoltageLabel8 = in["voltage_label_8"].(string)
		ret.VoltageLabel9 = in["voltage_label_9"].(string)
		ret.VoltageLabel10 = in["voltage_label_10"].(string)
		ret.VoltageLabel11 = in["voltage_label_11"].(string)
		ret.VoltageLabel12 = in["voltage_label_12"].(string)
		ret.VoltageLabel13 = in["voltage_label_13"].(string)
		ret.VoltageLabel14 = in["voltage_label_14"].(string)
		ret.VoltageLabel15 = in["voltage_label_15"].(string)
		ret.VoltageLabel16 = in["voltage_label_16"].(string)
		ret.VoltageLabel17 = in["voltage_label_17"].(string)
		ret.PowerUnit1 = in["power_unit1"].(string)
		ret.PowerUnit2 = in["power_unit2"].(string)
		ret.PowerUnit3 = in["power_unit3"].(string)
		ret.PowerUnit4 = in["power_unit4"].(string)
		ret.Pu2PhysicalTemperature = in["pu2_physical_temperature"].(string)
		ret.Pu2PhysicalTemperature2 = in["pu2_physical_temperature2"].(string)
		ret.Pu2VoltageLabel7 = in["pu2_voltage_label_7"].(string)
		ret.Pu2VoltageLabel8 = in["pu2_voltage_label_8"].(string)
		ret.Pu2VoltageLabel9 = in["pu2_voltage_label_9"].(string)
		ret.Pu2VoltageLabel10 = in["pu2_voltage_label_10"].(string)
	}
	return ret
}

func dataToEndpointSystemEnvironmentOper(d *schema.ResourceData) edpt.SystemEnvironmentOper {
	var ret edpt.SystemEnvironmentOper

	ret.Oper = getObjectSystemEnvironmentOperOper(d.Get("oper").([]interface{}))
	return ret
}
