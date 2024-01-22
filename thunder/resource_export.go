package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceExport() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_export`: Put files to remote site\n\n__PLACEHOLDER__",
		CreateContext: resourceExportCreate,
		UpdateContext: resourceExportUpdate,
		ReadContext:   resourceExportRead,
		DeleteContext: resourceExportDelete,

		Schema: map[string]*schema.Schema{
			"aflex": {
				Type: schema.TypeString, Optional: true, Description: "aFleX Script Source File",
			},
			"auth_jwks": {
				Type: schema.TypeString, Optional: true, Description: "Json web key",
			},
			"auth_portal": {
				Type: schema.TypeString, Optional: true, Description: "Portal file for http authentication",
			},
			"auth_portal_image": {
				Type: schema.TypeString, Optional: true, Description: "Image file for default portal",
			},
			"axdebug": {
				Type: schema.TypeString, Optional: true, Description: "AX Debug Packet File",
			},
			"bw_list": {
				Type: schema.TypeString, Optional: true, Description: "Black white List File",
			},
			"ca_cert": {
				Type: schema.TypeString, Optional: true, Description: "CA Cert File",
			},
			"capture_config": {
				Type: schema.TypeString, Optional: true, Description: "Capture-config pcapng file",
			},
			"capture_config_realtime": {
				Type: schema.TypeString, Optional: true, Description: "Capture-config pcapng real-time file (For GUI)",
			},
			"class_list": {
				Type: schema.TypeString, Optional: true, Description: "Class List File",
			},
			"csr": {
				Type: schema.TypeString, Optional: true, Description: "Certificate Signing Request",
			},
			"debug_monitor": {
				Type: schema.TypeString, Optional: true, Description: "Debug Monitor Output",
			},
			"debug_traffic_capture": {
				Type: schema.TypeString, Optional: true, Description: "Debug-Traffic-Capture pcapng file",
			},
			"debug_traffic_capture_chassis": {
				Type: schema.TypeString, Optional: true, Description: "Debug-Traffic-Capture pcapng file",
			},
			"debug_traffic_capture_chassis_slot": {
				Type: schema.TypeInt, Optional: true, Description: "specify slot id in chassis",
			},
			"dnssec_dnskey": {
				Type: schema.TypeString, Optional: true, Description: "DNSSEC DNSKEY(KSK) file for child zone",
			},
			"dnssec_ds": {
				Type: schema.TypeString, Optional: true, Description: "DNSSEC DS file for child zone",
			},
			"domain_list": {
				Type: schema.TypeString, Optional: true, Description: "Domain List File",
			},
			"externalfilename": {
				Type: schema.TypeString, Optional: true, Description: "Export the External Program from the System",
			},
			"fixed_nat": {
				Type: schema.TypeString, Optional: true, Description: "Fixed NAT Port Mapping File",
			},
			"fixed_nat_archive": {
				Type: schema.TypeString, Optional: true, Description: "Fixed NAT Port Mapping Archive File",
			},
			"geo_location": {
				Type: schema.TypeString, Optional: true, Description: "Geo-location CSV File",
			},
			"geo_location_archive": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"geo_location_archive_name": {
							Type: schema.TypeString, Optional: true, Description: "'GeoLite2-ASN-Archive': GeoLite2-ASN CSV Zipped File; 'GeoLite2-City-Archive': GeoLite2-City CSV Zipped File; 'GeoLite2-Country-Archive': GeoLite2-Country CSV Zipped File;",
						},
						"use_mgmt_port": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management port as source port",
						},
						"remote_file": {
							Type: schema.TypeString, Optional: true, Description: "Profile name for remote url",
						},
						"password": {
							Type: schema.TypeString, Optional: true, Description: "password for the remote site",
						},
					},
				},
			},
			"ip_map_list": {
				Type: schema.TypeString, Optional: true, Description: "IP Map List File",
			},
			"ipsec_error_dump": {
				Type: schema.TypeString, Optional: true, Description: "IPsec error dump File",
			},
			"local_uri_file": {
				Type: schema.TypeString, Optional: true, Description: "Local URI files for http response",
			},
			"lw_4o6": {
				Type: schema.TypeString, Optional: true, Description: "LW-4over6 Binding Table File",
			},
			"lw_4o6_binding_table_validation_log": {
				Type: schema.TypeString, Optional: true, Description: "LW-4over6 Binding Table Validation Log File",
			},
			"merged_pcap": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Export the merged pcap file when there are multiple Export sessions",
			},
			"mon_entity_debug_file": {
				Type: schema.TypeString, Optional: true, Description: "Enter Mon entity debug file name",
			},
			"password": {
				Type: schema.TypeString, Optional: true, Description: "password for the remote site",
			},
			"per_cpu": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Export the per-cpu files along with the merged pcap file in .tgz format",
			},
			"pkt_count": {
				Type: schema.TypeInt, Optional: true, Description: "Specify number of latest packets to export",
			},
			"pktcapture_file": {
				Type: schema.TypeString, Optional: true, Description: "Enter Pktcapture file name",
			},
			"profile": {
				Type: schema.TypeString, Optional: true, Description: "Startup-config Profile",
			},
			"remote_file": {
				Type: schema.TypeString, Optional: true, Description: "profile name for remote url",
			},
			"rpz": {
				Type: schema.TypeString, Optional: true, Description: "Response Policy Zone File",
			},
			"running_config": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Running Config",
			},
			"saml_idp_name": {
				Type: schema.TypeString, Optional: true, Description: "SAML metadata of identity provider",
			},
			"ssl_cert": {
				Type: schema.TypeString, Optional: true, Description: "SSL Cert File",
			},
			"ssl_cert_key": {
				Type: schema.TypeString, Optional: true, Description: "Local SSL Key/Certificate file name",
			},
			"ssl_crl": {
				Type: schema.TypeString, Optional: true, Description: "SSL Crl File",
			},
			"ssl_key": {
				Type: schema.TypeString, Optional: true, Description: "SSL Key File",
			},
			"startup_config": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Startup Config",
			},
			"status_check": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "check export task status",
			},
			"store": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"delete": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Delete an export store profile",
						},
						"create": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Create an export store profile",
						},
						"name": {
							Type: schema.TypeString, Optional: true, Description: "profile name to store remote url",
						},
						"remote_file": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
					},
				},
			},
			"store_name": {
				Type: schema.TypeString, Optional: true, Description: "Export store name",
			},
			"syslog": {
				Type: schema.TypeString, Optional: true, Description: "Enter \"messages\" as the default syslog file name",
			},
			"tgz": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Export the merged pcap in .tgz format",
			},
			"thales_kmdata": {
				Type: schema.TypeString, Optional: true, Description: "Thales Kmdata files",
			},
			"thales_secworld": {
				Type: schema.TypeString, Optional: true, Description: "Thales security world files",
			},
			"tsig": {
				Type: schema.TypeString, Optional: true, Description: "Transaction SIGnatures Key file",
			},
			"use_mgmt_port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management port as source port",
			},
			"visibility": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Export Visibility module related files",
			},
			"xml_schema": {
				Type: schema.TypeString, Optional: true, Description: "XML-Schema File",
			},
		},
	}
}
func resourceExportCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceExportCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointExport(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceExportRead(ctx, d, meta)
	}
	return diags
}

func resourceExportUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceExportUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointExport(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceExportRead(ctx, d, meta)
	}
	return diags
}
func resourceExportDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceExportDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointExport(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceExportRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceExportRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointExport(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectExportGeoLocationArchive347(d []interface{}) edpt.ExportGeoLocationArchive347 {

	count1 := len(d)
	var ret edpt.ExportGeoLocationArchive347
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GeoLocationArchiveName = in["geo_location_archive_name"].(string)
		ret.UseMgmtPort = in["use_mgmt_port"].(int)
		ret.RemoteFile = in["remote_file"].(string)
		ret.Password = in["password"].(string)
	}
	return ret
}

func getObjectExportStore348(d []interface{}) edpt.ExportStore348 {

	count1 := len(d)
	var ret edpt.ExportStore348
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Delete = in["delete"].(int)
		ret.Create = in["create"].(int)
		ret.Name = in["name"].(string)
		ret.RemoteFile = in["remote_file"].(string)
	}
	return ret
}

func dataToEndpointExport(d *schema.ResourceData) edpt.Export {
	var ret edpt.Export
	ret.Inst.Aflex = d.Get("aflex").(string)
	ret.Inst.AuthJwks = d.Get("auth_jwks").(string)
	ret.Inst.AuthPortal = d.Get("auth_portal").(string)
	ret.Inst.AuthPortalImage = d.Get("auth_portal_image").(string)
	ret.Inst.Axdebug = d.Get("axdebug").(string)
	ret.Inst.BwList = d.Get("bw_list").(string)
	ret.Inst.CaCert = d.Get("ca_cert").(string)
	ret.Inst.CaptureConfig = d.Get("capture_config").(string)
	ret.Inst.CaptureConfigRealtime = d.Get("capture_config_realtime").(string)
	ret.Inst.ClassList = d.Get("class_list").(string)
	ret.Inst.Csr = d.Get("csr").(string)
	ret.Inst.DebugMonitor = d.Get("debug_monitor").(string)
	ret.Inst.DebugTrafficCapture = d.Get("debug_traffic_capture").(string)
	ret.Inst.DebugTrafficCaptureChassis = d.Get("debug_traffic_capture_chassis").(string)
	ret.Inst.DebugTrafficCaptureChassisSlot = d.Get("debug_traffic_capture_chassis_slot").(int)
	ret.Inst.DnssecDnskey = d.Get("dnssec_dnskey").(string)
	ret.Inst.DnssecDs = d.Get("dnssec_ds").(string)
	ret.Inst.DomainList = d.Get("domain_list").(string)
	ret.Inst.Externalfilename = d.Get("externalfilename").(string)
	ret.Inst.FixedNat = d.Get("fixed_nat").(string)
	ret.Inst.FixedNatArchive = d.Get("fixed_nat_archive").(string)
	ret.Inst.GeoLocation = d.Get("geo_location").(string)
	ret.Inst.GeoLocationArchive = getObjectExportGeoLocationArchive347(d.Get("geo_location_archive").([]interface{}))
	ret.Inst.IpMapList = d.Get("ip_map_list").(string)
	ret.Inst.IpsecErrorDump = d.Get("ipsec_error_dump").(string)
	ret.Inst.LocalUriFile = d.Get("local_uri_file").(string)
	ret.Inst.Lw4o6 = d.Get("lw_4o6").(string)
	ret.Inst.Lw4o6BindingTableValidationLog = d.Get("lw_4o6_binding_table_validation_log").(string)
	ret.Inst.MergedPcap = d.Get("merged_pcap").(int)
	ret.Inst.MonEntityDebugFile = d.Get("mon_entity_debug_file").(string)
	ret.Inst.Password = d.Get("password").(string)
	ret.Inst.PerCpu = d.Get("per_cpu").(int)
	ret.Inst.PktCount = d.Get("pkt_count").(int)
	ret.Inst.PktcaptureFile = d.Get("pktcapture_file").(string)
	ret.Inst.Profile = d.Get("profile").(string)
	ret.Inst.RemoteFile = d.Get("remote_file").(string)
	ret.Inst.Rpz = d.Get("rpz").(string)
	ret.Inst.RunningConfig = d.Get("running_config").(int)
	ret.Inst.SamlIdpName = d.Get("saml_idp_name").(string)
	ret.Inst.SslCert = d.Get("ssl_cert").(string)
	ret.Inst.SslCertKey = d.Get("ssl_cert_key").(string)
	ret.Inst.SslCrl = d.Get("ssl_crl").(string)
	ret.Inst.SslKey = d.Get("ssl_key").(string)
	ret.Inst.StartupConfig = d.Get("startup_config").(int)
	ret.Inst.StatusCheck = d.Get("status_check").(int)
	ret.Inst.Store = getObjectExportStore348(d.Get("store").([]interface{}))
	ret.Inst.StoreName = d.Get("store_name").(string)
	ret.Inst.Syslog = d.Get("syslog").(string)
	ret.Inst.Tgz = d.Get("tgz").(int)
	ret.Inst.ThalesKmdata = d.Get("thales_kmdata").(string)
	ret.Inst.ThalesSecworld = d.Get("thales_secworld").(string)
	ret.Inst.Tsig = d.Get("tsig").(string)
	ret.Inst.UseMgmtPort = d.Get("use_mgmt_port").(int)
	ret.Inst.Visibility = d.Get("visibility").(int)
	ret.Inst.XmlSchema = d.Get("xml_schema").(string)
	return ret
}
