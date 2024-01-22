package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVersionOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_version_oper`: Operational Status for the object version\n\n__PLACEHOLDER__",
		ReadContext: resourceVersionOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hw_platform": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"copyright": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"sw_version": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"plat_features": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"boot_from": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"serial_number": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"aflex_version": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"axapi_version": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"pri_gui_version": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"sec_gui_version": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"firmware_version": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"hd_pri": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"hd_sec": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"cf_pri": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"cf_sec": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"last_config_saved_time": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"virtualization_type": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"sys_poll_mode": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"product": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"hw_code": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"current_time": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"up_time": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"nun_ctrl_cpus": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"buff_size": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"io_buff_enabled": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"build_type": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"cots_sys_mfg": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"cots_sys_name": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"cots_sys_ver": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"series_name": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"hostname": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"alldynamic": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceVersionOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceVersionOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointVersionOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		VersionOperOper := setObjectVersionOperOper(res)
		d.Set("oper", VersionOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectVersionOperOper(ret edpt.DataVersionOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"hw_platform":            ret.DtVersionOper.Oper.HwPlatform,
			"copyright":              ret.DtVersionOper.Oper.Copyright,
			"sw_version":             ret.DtVersionOper.Oper.SwVersion,
			"plat_features":          ret.DtVersionOper.Oper.PlatFeatures,
			"boot_from":              ret.DtVersionOper.Oper.BootFrom,
			"serial_number":          ret.DtVersionOper.Oper.SerialNumber,
			"aflex_version":          ret.DtVersionOper.Oper.AflexVersion,
			"axapi_version":          ret.DtVersionOper.Oper.AxapiVersion,
			"pri_gui_version":        ret.DtVersionOper.Oper.PriGuiVersion,
			"sec_gui_version":        ret.DtVersionOper.Oper.SecGuiVersion,
			"firmware_version":       ret.DtVersionOper.Oper.FirmwareVersion,
			"hd_pri":                 ret.DtVersionOper.Oper.HdPri,
			"hd_sec":                 ret.DtVersionOper.Oper.HdSec,
			"cf_pri":                 ret.DtVersionOper.Oper.CfPri,
			"cf_sec":                 ret.DtVersionOper.Oper.CfSec,
			"last_config_saved_time": ret.DtVersionOper.Oper.LastConfigSavedTime,
			"virtualization_type":    ret.DtVersionOper.Oper.VirtualizationType,
			"sys_poll_mode":          ret.DtVersionOper.Oper.SysPollMode,
			"product":                ret.DtVersionOper.Oper.Product,
			"hw_code":                ret.DtVersionOper.Oper.HwCode,
			"current_time":           ret.DtVersionOper.Oper.CurrentTime,
			"up_time":                ret.DtVersionOper.Oper.UpTime,
			"nun_ctrl_cpus":          ret.DtVersionOper.Oper.NunCtrlCpus,
			"buff_size":              ret.DtVersionOper.Oper.BuffSize,
			"io_buff_enabled":        ret.DtVersionOper.Oper.IoBuffEnabled,
			"build_type":             ret.DtVersionOper.Oper.BuildType,
			"cots_sys_mfg":           ret.DtVersionOper.Oper.CotsSysMfg,
			"cots_sys_name":          ret.DtVersionOper.Oper.CotsSysName,
			"cots_sys_ver":           ret.DtVersionOper.Oper.CotsSysVer,
			"series_name":            ret.DtVersionOper.Oper.SeriesName,
			"hostname":               ret.DtVersionOper.Oper.Hostname,
			"alldynamic":             ret.DtVersionOper.Oper.Alldynamic,
		},
	}
}

func getObjectVersionOperOper(d []interface{}) edpt.VersionOperOper {

	count1 := len(d)
	var ret edpt.VersionOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.HwPlatform = in["hw_platform"].(string)
		ret.Copyright = in["copyright"].(string)
		ret.SwVersion = in["sw_version"].(string)
		ret.PlatFeatures = in["plat_features"].(string)
		ret.BootFrom = in["boot_from"].(string)
		ret.SerialNumber = in["serial_number"].(string)
		ret.AflexVersion = in["aflex_version"].(string)
		ret.AxapiVersion = in["axapi_version"].(string)
		ret.PriGuiVersion = in["pri_gui_version"].(string)
		ret.SecGuiVersion = in["sec_gui_version"].(string)
		ret.FirmwareVersion = in["firmware_version"].(string)
		ret.HdPri = in["hd_pri"].(string)
		ret.HdSec = in["hd_sec"].(string)
		ret.CfPri = in["cf_pri"].(string)
		ret.CfSec = in["cf_sec"].(string)
		ret.LastConfigSavedTime = in["last_config_saved_time"].(string)
		ret.VirtualizationType = in["virtualization_type"].(string)
		ret.SysPollMode = in["sys_poll_mode"].(string)
		ret.Product = in["product"].(string)
		ret.HwCode = in["hw_code"].(string)
		ret.CurrentTime = in["current_time"].(string)
		ret.UpTime = in["up_time"].(string)
		ret.NunCtrlCpus = in["nun_ctrl_cpus"].(int)
		ret.BuffSize = in["buff_size"].(int)
		ret.IoBuffEnabled = in["io_buff_enabled"].(string)
		ret.BuildType = in["build_type"].(string)
		ret.CotsSysMfg = in["cots_sys_mfg"].(string)
		ret.CotsSysName = in["cots_sys_name"].(string)
		ret.CotsSysVer = in["cots_sys_ver"].(string)
		ret.SeriesName = in["series_name"].(string)
		ret.Hostname = in["hostname"].(string)
		ret.Alldynamic = in["alldynamic"].(int)
	}
	return ret
}

func dataToEndpointVersionOper(d *schema.ResourceData) edpt.VersionOper {
	var ret edpt.VersionOper

	ret.Oper = getObjectVersionOperOper(d.Get("oper").([]interface{}))
	return ret
}
