package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemViewHardwareViewOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_view_hardware_view_oper`: Operational Status for the object hardware-view\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemViewHardwareViewOperRead,

		Schema: map[string]*schema.Schema{
			"oper": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"platform_description": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"serial": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"cpu": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"cpu_cores": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"cpu_stepping": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"storage": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"memory": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ssl_cards": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ssl_devices": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"nitroxpx": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"nitrox3": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"nitrox3_cores": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"nitrox5": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"nitrox5_cores": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"nitrox2": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"nitrox1": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"hsm": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"unknown_ssl_cards": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"coleto_ssl_cards": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"octeon": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"compression_cards": {
							Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"gzip_devices": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"aha363": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
									"unknown_compression": {
										Type: schema.TypeInt, Optional: true, Description: "",
									},
								},
							},
						},
						"l23_asic": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ipmi": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"ports": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"plat_flag": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"bios_version": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"bios_release_date": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"nvm_firmware_versoin": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"fpga_summary": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"fpga_date": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"disk_total": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"disk_used": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"disk_free": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"disk_percentage": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"disk1_status": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"disk2_status": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"num_disks": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"raid_present": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"raid_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"md_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"md_pri": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"md_sec": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"psu1_np15": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"psu2_np15": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"spe_present": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"bypass_pr": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"bypass_list": {
							Type: schema.TypeList, Optional: true, Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"bypass_name": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
									"bypass_info": {
										Type: schema.TypeString, Optional: true, Description: "",
									},
								},
							},
						},
						"alldynamic": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"mcpld_type": {
							Type: schema.TypeInt, Optional: true, Description: "",
						},
						"mcpld_date": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
		},
	}
}

func resourceSystemViewHardwareViewOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemViewHardwareViewOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemViewHardwareViewOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		d.SetId(obj.GetId())
		logger.Println(res)
		SystemViewHardwareViewOperOper := setObjectSystemViewHardwareViewOperOper(res)
		d.Set("oper", SystemViewHardwareViewOperOper)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemViewHardwareViewOperOper(ret edpt.DataSystemViewHardwareViewOper) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"platform_description": ret.DtSystemViewHardwareViewOper.Oper.PlatformDescription,
			"serial":               ret.DtSystemViewHardwareViewOper.Oper.Serial,
			"cpu":                  ret.DtSystemViewHardwareViewOper.Oper.Cpu,
			"cpu_cores":            ret.DtSystemViewHardwareViewOper.Oper.CpuCores,
			"cpu_stepping":         ret.DtSystemViewHardwareViewOper.Oper.CpuStepping,
			"storage":              ret.DtSystemViewHardwareViewOper.Oper.Storage,
			"memory":               ret.DtSystemViewHardwareViewOper.Oper.Memory,
			"ssl_cards":            setObjectSystemViewHardwareViewOperOperSslCards(ret.DtSystemViewHardwareViewOper.Oper.SslCards),
			"octeon":               ret.DtSystemViewHardwareViewOper.Oper.Octeon,
			"compression_cards":    setObjectSystemViewHardwareViewOperOperCompressionCards(ret.DtSystemViewHardwareViewOper.Oper.CompressionCards),
			"l23_asic":             ret.DtSystemViewHardwareViewOper.Oper.L23Asic,
			"ipmi":                 ret.DtSystemViewHardwareViewOper.Oper.Ipmi,
			"ports":                ret.DtSystemViewHardwareViewOper.Oper.Ports,
			"plat_flag":            ret.DtSystemViewHardwareViewOper.Oper.PlatFlag,
			"bios_version":         ret.DtSystemViewHardwareViewOper.Oper.BiosVersion,
			"bios_release_date":    ret.DtSystemViewHardwareViewOper.Oper.BiosReleaseDate,
			"nvm_firmware_versoin": ret.DtSystemViewHardwareViewOper.Oper.NvmFirmwareVersoin,
			"fpga_summary":         ret.DtSystemViewHardwareViewOper.Oper.FpgaSummary,
			"fpga_date":            ret.DtSystemViewHardwareViewOper.Oper.FpgaDate,
			"disk_total":           ret.DtSystemViewHardwareViewOper.Oper.DiskTotal,
			"disk_used":            ret.DtSystemViewHardwareViewOper.Oper.DiskUsed,
			"disk_free":            ret.DtSystemViewHardwareViewOper.Oper.DiskFree,
			"disk_percentage":      ret.DtSystemViewHardwareViewOper.Oper.DiskPercentage,
			"disk1_status":         ret.DtSystemViewHardwareViewOper.Oper.Disk1Status,
			"disk2_status":         ret.DtSystemViewHardwareViewOper.Oper.Disk2Status,
			"num_disks":            ret.DtSystemViewHardwareViewOper.Oper.NumDisks,
			"raid_present":         ret.DtSystemViewHardwareViewOper.Oper.Raid_present,
			"raid_list":            setSliceSystemViewHardwareViewOperOperRaidList(ret.DtSystemViewHardwareViewOper.Oper.RaidList),
			"psu1_np15":            ret.DtSystemViewHardwareViewOper.Oper.Psu1_np15,
			"psu2_np15":            ret.DtSystemViewHardwareViewOper.Oper.Psu2_np15,
			"spe_present":          ret.DtSystemViewHardwareViewOper.Oper.Spe_present,
			"bypass_pr":            ret.DtSystemViewHardwareViewOper.Oper.BypassPr,
			"bypass_list":          setSliceSystemViewHardwareViewOperOperBypassList(ret.DtSystemViewHardwareViewOper.Oper.BypassList),
			"alldynamic":           ret.DtSystemViewHardwareViewOper.Oper.Alldynamic,
			"mcpld_type":           ret.DtSystemViewHardwareViewOper.Oper.McpldType,
			"mcpld_date":           ret.DtSystemViewHardwareViewOper.Oper.McpldDate,
		},
	}
}

func setObjectSystemViewHardwareViewOperOperSslCards(d edpt.SystemViewHardwareViewOperOperSslCards) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["ssl_devices"] = d.SslDevices

	in["nitroxpx"] = d.Nitroxpx

	in["nitrox3"] = d.Nitrox3

	in["nitrox3_cores"] = d.Nitrox3Cores

	in["nitrox5"] = d.Nitrox5

	in["nitrox5_cores"] = d.Nitrox5Cores

	in["nitrox2"] = d.Nitrox2

	in["nitrox1"] = d.Nitrox1

	in["hsm"] = d.Hsm

	in["unknown_ssl_cards"] = d.UnknownSslCards

	in["coleto_ssl_cards"] = d.ColetoSslCards
	result = append(result, in)
	return result
}

func setObjectSystemViewHardwareViewOperOperCompressionCards(d edpt.SystemViewHardwareViewOperOperCompressionCards) []map[string]interface{} {
	result := []map[string]interface{}{}
	in := make(map[string]interface{})

	in["gzip_devices"] = d.GzipDevices

	in["aha363"] = d.Aha363

	in["unknown_compression"] = d.UnknownCompression
	result = append(result, in)
	return result
}

func setSliceSystemViewHardwareViewOperOperRaidList(d []edpt.SystemViewHardwareViewOperOperRaidList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["md_name"] = item.MdName
		in["md_pri"] = item.MdPri
		in["md_sec"] = item.MdSec
		result = append(result, in)
	}
	return result
}

func setSliceSystemViewHardwareViewOperOperBypassList(d []edpt.SystemViewHardwareViewOperOperBypassList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["bypass_name"] = item.BypassName
		in["bypass_info"] = item.BypassInfo
		result = append(result, in)
	}
	return result
}

func getObjectSystemViewHardwareViewOperOper(d []interface{}) edpt.SystemViewHardwareViewOperOper {

	count1 := len(d)
	var ret edpt.SystemViewHardwareViewOperOper
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.PlatformDescription = in["platform_description"].(string)
		ret.Serial = in["serial"].(string)
		ret.Cpu = in["cpu"].(string)
		ret.CpuCores = in["cpu_cores"].(int)
		ret.CpuStepping = in["cpu_stepping"].(int)
		ret.Storage = in["storage"].(string)
		ret.Memory = in["memory"].(string)
		ret.SslCards = getObjectSystemViewHardwareViewOperOperSslCards(in["ssl_cards"].([]interface{}))
		ret.Octeon = in["octeon"].(int)
		ret.CompressionCards = getObjectSystemViewHardwareViewOperOperCompressionCards(in["compression_cards"].([]interface{}))
		ret.L23Asic = in["l23_asic"].(string)
		ret.Ipmi = in["ipmi"].(string)
		ret.Ports = in["ports"].(string)
		ret.PlatFlag = in["plat_flag"].(string)
		ret.BiosVersion = in["bios_version"].(string)
		ret.BiosReleaseDate = in["bios_release_date"].(string)
		ret.NvmFirmwareVersoin = in["nvm_firmware_versoin"].(string)
		ret.FpgaSummary = in["fpga_summary"].(string)
		ret.FpgaDate = in["fpga_date"].(string)
		ret.DiskTotal = in["disk_total"].(int)
		ret.DiskUsed = in["disk_used"].(int)
		ret.DiskFree = in["disk_free"].(int)
		ret.DiskPercentage = in["disk_percentage"].(int)
		ret.Disk1Status = in["disk1_status"].(string)
		ret.Disk2Status = in["disk2_status"].(string)
		ret.NumDisks = in["num_disks"].(int)
		ret.Raid_present = in["raid_present"].(int)
		ret.RaidList = getSliceSystemViewHardwareViewOperOperRaidList(in["raid_list"].([]interface{}))
		ret.Psu1_np15 = in["psu1_np15"].(string)
		ret.Psu2_np15 = in["psu2_np15"].(string)
		ret.Spe_present = in["spe_present"].(string)
		ret.BypassPr = in["bypass_pr"].(int)
		ret.BypassList = getSliceSystemViewHardwareViewOperOperBypassList(in["bypass_list"].([]interface{}))
		ret.Alldynamic = in["alldynamic"].(int)
		ret.McpldType = in["mcpld_type"].(int)
		ret.McpldDate = in["mcpld_date"].(string)
	}
	return ret
}

func getObjectSystemViewHardwareViewOperOperSslCards(d []interface{}) edpt.SystemViewHardwareViewOperOperSslCards {

	count1 := len(d)
	var ret edpt.SystemViewHardwareViewOperOperSslCards
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SslDevices = in["ssl_devices"].(int)
		ret.Nitroxpx = in["nitroxpx"].(int)
		ret.Nitrox3 = in["nitrox3"].(int)
		ret.Nitrox3Cores = in["nitrox3_cores"].(int)
		ret.Nitrox5 = in["nitrox5"].(int)
		ret.Nitrox5Cores = in["nitrox5_cores"].(int)
		ret.Nitrox2 = in["nitrox2"].(int)
		ret.Nitrox1 = in["nitrox1"].(int)
		ret.Hsm = in["hsm"].(int)
		ret.UnknownSslCards = in["unknown_ssl_cards"].(int)
		ret.ColetoSslCards = in["coleto_ssl_cards"].(int)
	}
	return ret
}

func getObjectSystemViewHardwareViewOperOperCompressionCards(d []interface{}) edpt.SystemViewHardwareViewOperOperCompressionCards {

	count1 := len(d)
	var ret edpt.SystemViewHardwareViewOperOperCompressionCards
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GzipDevices = in["gzip_devices"].(int)
		ret.Aha363 = in["aha363"].(int)
		ret.UnknownCompression = in["unknown_compression"].(int)
	}
	return ret
}

func getSliceSystemViewHardwareViewOperOperRaidList(d []interface{}) []edpt.SystemViewHardwareViewOperOperRaidList {

	count1 := len(d)
	ret := make([]edpt.SystemViewHardwareViewOperOperRaidList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemViewHardwareViewOperOperRaidList
		oi.MdName = in["md_name"].(string)
		oi.MdPri = in["md_pri"].(string)
		oi.MdSec = in["md_sec"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSystemViewHardwareViewOperOperBypassList(d []interface{}) []edpt.SystemViewHardwareViewOperOperBypassList {

	count1 := len(d)
	ret := make([]edpt.SystemViewHardwareViewOperOperBypassList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemViewHardwareViewOperOperBypassList
		oi.BypassName = in["bypass_name"].(string)
		oi.BypassInfo = in["bypass_info"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemViewHardwareViewOper(d *schema.ResourceData) edpt.SystemViewHardwareViewOper {
	var ret edpt.SystemViewHardwareViewOper

	ret.Oper = getObjectSystemViewHardwareViewOperOper(d.Get("oper").([]interface{}))
	return ret
}
