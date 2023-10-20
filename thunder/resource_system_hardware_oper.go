package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemHardwareOper() *schema.Resource {
	return &schema.Resource{
		Description: "`thunder_system_hardware_oper`: Operational Status for the object hardware\n\n__PLACEHOLDER__",
		ReadContext: resourceSystemHardwareOperRead,
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
						"nvm_firmware_version": {
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

func resourceSystemHardwareOperRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSystemHardwareOperRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSystemHardwareOper(d)
		res, err := obj.Get(client.Token, client.Host, d.Id(), logger)
		items := setObjectSystemHardwareOperOper(res)
		d.SetId(obj.GetId())
		d.Set("oper", items)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func setObjectSystemHardwareOperOper(res edpt.SystemHardwaree) []interface{} {
	return []interface{}{
		map[string]interface{}{
			"platform_description": res.DataSystemHardware.Oper.PlatformDescription,
			"serial":               res.DataSystemHardware.Oper.Serial,
			"cpu":                  res.DataSystemHardware.Oper.Cpu,
			"cpu_cores":            res.DataSystemHardware.Oper.CpuCores,
			"cpu_stepping":         res.DataSystemHardware.Oper.CpuStepping,
			"storage":              res.DataSystemHardware.Oper.Storage,
			"memory":               res.DataSystemHardware.Oper.Memory,
			"ssl_cards":            setObjectSystemHardwareOperOperSslCards(res.DataSystemHardware.Oper.SslCards),
			"octeon":               res.DataSystemHardware.Oper.Octeon,
			"compression_cards":    setObjectSystemHardwareOperOperCompressionCards(res.DataSystemHardware.Oper.CompressionCards),
			"l23_asic":             res.DataSystemHardware.Oper.L23Asic,
			"ipmi":                 res.DataSystemHardware.Oper.Ipmi,
			"ports":                res.DataSystemHardware.Oper.Ports,
			"plat_flag":            res.DataSystemHardware.Oper.PlatFlag,
			"bios_version":         res.DataSystemHardware.Oper.BiosVersion,
			"bios_release_date":    res.DataSystemHardware.Oper.BiosReleaseDate,
			"nvm_firmware_version": res.DataSystemHardware.Oper.NvmFirmwareVersion,
			"fpga_summary":         res.DataSystemHardware.Oper.FpgaSummary,
			"fpga_date":            res.DataSystemHardware.Oper.FpgaDate,
			"disk_total":           res.DataSystemHardware.Oper.DiskTotal,
			"disk_used":            res.DataSystemHardware.Oper.DiskUsed,
			"disk_free":            res.DataSystemHardware.Oper.DiskFree,
			"disk_percentage":      res.DataSystemHardware.Oper.DiskPercentage,
			"disk1_status":         res.DataSystemHardware.Oper.Disk1Status,
			"disk2_status":         res.DataSystemHardware.Oper.Disk2Status,
			"num_disks":            res.DataSystemHardware.Oper.NumDisks,
			"raid_present":         res.DataSystemHardware.Oper.Raid_present,
			"raid_list":            setSliceSystemHardwareOperOperRaidList(res.DataSystemHardware.Oper.RaidList),
			"psu1_np15":            res.DataSystemHardware.Oper.Psu1_np15,
			"psu2_np15":            res.DataSystemHardware.Oper.Psu2_np15,
			"spe_present":          res.DataSystemHardware.Oper.Spe_present,
			"bypass_pr":            res.DataSystemHardware.Oper.BypassPr,
			"bypass_list":          setSliceSystemHardwareOperOperBypassList(res.DataSystemHardware.Oper.BypassList),
			"alldynamic":           res.DataSystemHardware.Oper.Alldynamic,
			"mcpld_type":           res.DataSystemHardware.Oper.McpldType,
			"mcpld_date":           res.DataSystemHardware.Oper.McpldDate,
		},
	}

}

func setObjectSystemHardwareOperOperSslCards(d []edpt.SystemHardwareOperOperSslCards) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["ssl_devices"] = item.SslDevices
		in["nitroxpx"] = item.Nitroxpx
		in["nitrox3"] = item.Nitrox3
		in["nitrox3_cores"] = item.Nitrox3Cores
		in["nitrox5"] = item.Nitrox5
		in["nitrox5_cores"] = item.Nitrox5Cores
		in["nitrox2"] = item.Nitrox2
		in["nitrox1"] = item.Nitrox1
		in["hsm"] = item.Hsm
		in["unknown_ssl_cards"] = item.UnknownSslCards
		in["coleto_ssl_cards"] = item.ColetoSslCards
		result = append(result, in)
	}
	return result
}

func setObjectSystemHardwareOperOperCompressionCards(d []edpt.SystemHardwareOperOperCompressionCards) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["gzip_devices"] = item.GzipDevices
		in["aha363"] = item.Aha363
		in["unknown_compression"] = item.UnknownCompression
		result = append(result, in)
	}
	return result
}

func setSliceSystemHardwareOperOperRaidList(d []edpt.SystemHardwareOperOperRaidList) []map[string]interface{} {
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

func setSliceSystemHardwareOperOperBypassList(d []edpt.SystemHardwareOperOperBypassList) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, item := range d {
		in := make(map[string]interface{})
		in["bypass_name"] = item.BypassName
		in["bypass_info"] = item.BypassInfo
		result = append(result, in)
	}
	return result
}

func getObjectSystemHardwareOperOper(d []interface{}) edpt.SystemHardwareOperOper {
	count := len(d)
	var ret edpt.SystemHardwareOperOper
	if count > 0 {
		in := d[0].(map[string]interface{})
		ret.PlatformDescription = in["platform_description"].(string)
		ret.Serial = in["serial"].(string)
		ret.Cpu = in["cpu"].(string)
		ret.CpuCores = in["cpu_cores"].(int)
		ret.CpuStepping = in["cpu_stepping"].(int)
		ret.Storage = in["storage"].(string)
		ret.Memory = in["memory"].(string)
		ret.SslCards = getObjectSystemHardwareOperOperSslCards(in["ssl_cards"].([]interface{}))
		ret.Octeon = in["octeon"].(int)
		ret.CompressionCards = getObjectSystemHardwareOperOperCompressionCards(in["compression_cards"].([]interface{}))
		ret.L23Asic = in["l23_asic"].(string)
		ret.Ipmi = in["ipmi"].(string)
		ret.Ports = in["ports"].(string)
		ret.PlatFlag = in["plat_flag"].(string)
		ret.BiosVersion = in["bios_version"].(string)
		ret.BiosReleaseDate = in["bios_release_date"].(string)
		ret.NvmFirmwareVersion = in["nvm_firmware_version"].(string)
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
		ret.RaidList = getSliceSystemHardwareOperOperRaidList(in["raid_list"].([]interface{}))
		ret.Psu1_np15 = in["psu1_np15"].(string)
		ret.Psu2_np15 = in["psu2_np15"].(string)
		ret.Spe_present = in["spe_present"].(string)
		ret.BypassPr = in["bypass_pr"].(int)
		ret.BypassList = getSliceSystemHardwareOperOperBypassList(in["bypass_list"].([]interface{}))
		ret.Alldynamic = in["alldynamic"].(int)
		ret.McpldType = in["mcpld_type"].(int)
		ret.McpldDate = in["mcpld_date"].(string)
	}
	return ret
}

func getObjectSystemHardwareOperOperSslCards(d []interface{}) []edpt.SystemHardwareOperOperSslCards {
	count := len(d)
	ret := make([]edpt.SystemHardwareOperOperSslCards, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemHardwareOperOperSslCards
		oi.SslDevices = in["ssl_devices"].(int)
		oi.Nitroxpx = in["nitroxpx"].(int)
		oi.Nitrox3 = in["nitrox3"].(int)
		oi.Nitrox3Cores = in["nitrox3_cores"].(int)
		oi.Nitrox5 = in["nitrox5"].(int)
		oi.Nitrox5Cores = in["nitrox5_cores"].(int)
		oi.Nitrox2 = in["nitrox2"].(int)
		oi.Nitrox1 = in["nitrox1"].(int)
		oi.Hsm = in["hsm"].(int)
		oi.UnknownSslCards = in["unknown_ssl_cards"].(int)
		oi.ColetoSslCards = in["coleto_ssl_cards"].(int)
	}
	return ret
}

func getObjectSystemHardwareOperOperCompressionCards(d []interface{}) []edpt.SystemHardwareOperOperCompressionCards {
	count := len(d)
	ret := make([]edpt.SystemHardwareOperOperCompressionCards, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemHardwareOperOperCompressionCards
		oi.GzipDevices = in["gzip_devices"].(int)
		oi.Aha363 = in["aha363"].(int)
		oi.UnknownCompression = in["unknown_compression"].(int)
	}
	return ret
}

func getSliceSystemHardwareOperOperRaidList(d []interface{}) []edpt.SystemHardwareOperOperRaidList {
	count := len(d)
	ret := make([]edpt.SystemHardwareOperOperRaidList, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemHardwareOperOperRaidList
		oi.MdName = in["md_name"].(string)
		oi.MdPri = in["md_pri"].(string)
		oi.MdSec = in["md_sec"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func getSliceSystemHardwareOperOperBypassList(d []interface{}) []edpt.SystemHardwareOperOperBypassList {
	count := len(d)
	ret := make([]edpt.SystemHardwareOperOperBypassList, 0, count)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.SystemHardwareOperOperBypassList
		oi.BypassName = in["bypass_name"].(string)
		oi.BypassInfo = in["bypass_info"].(string)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointSystemHardwareOper(d *schema.ResourceData) edpt.SystemHardwareOper {
	var ret edpt.SystemHardwareOper
	ret.Oper = getObjectSystemHardwareOperOper(d.Get("oper").([]interface{}))
	return ret
}
