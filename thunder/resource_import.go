package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceImport() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_import`: Get files from remote site\n\n__PLACEHOLDER__",
		CreateContext: resourceImportCreate,
		UpdateContext: resourceImportUpdate,
		ReadContext:   resourceImportRead,
		DeleteContext: resourceImportDelete,

		Schema: map[string]*schema.Schema{
			"aflex": {
				Type: schema.TypeString, Optional: true, Description: "aFleX Script Source File",
			},
			"auth_jwks": {
				Type: schema.TypeString, Optional: true, Description: "JSON web key",
			},
			"auth_portal": {
				Type: schema.TypeString, Optional: true, Description: "Portal file for http authentication",
			},
			"auth_portal_image": {
				Type: schema.TypeString, Optional: true, Description: "Image file for default portal",
			},
			"auth_saml_idp": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"saml_idp_name": {
							Type: schema.TypeString, Optional: true, Description: "Metadata name",
						},
						"verify_xml_signature": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Verify metadata's XML signature",
						},
						"overwrite": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Overwrite existing file",
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
			"background": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Background mode for importing class-list",
			},
			"bios_file": {
				Type: schema.TypeString, Optional: true, Description: "BIOS Image file",
			},
			"bw_list": {
				Type: schema.TypeString, Optional: true, Description: "Black white List File",
			},
			"ca_cert": {
				Type: schema.TypeString, Optional: true, Description: "CA Cert File(enter bulk when import an archive file)",
			},
			"certificate_type": {
				Type: schema.TypeString, Optional: true, Description: "'pem': pem; 'der': der; 'pfx': pfx; 'p7b': p7b;",
			},
			"class_list": {
				Type: schema.TypeString, Optional: true, Description: "Class List File",
			},
			"class_list_convert": {
				Type: schema.TypeString, Optional: true, Description: "Convert Class List File to A10 format",
			},
			"class_list_type": {
				Type: schema.TypeString, Optional: true, Description: "'ac': ac; 'ipv4': ipv4; 'ipv6': ipv6; 'string': string; 'string-case-insensitive': string-case-insensitive;",
			},
			"cloud_config": {
				Type: schema.TypeString, Optional: true, Description: "Cloud Configuration File",
			},
			"cloud_creds": {
				Type: schema.TypeString, Optional: true, Description: "Cloud Credentials File",
			},
			"csr_generate": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Generate CSR file",
			},
			"ddos_script": {
				Type: schema.TypeString, Optional: true, Description: "Python/Perl/Bash/TCL script to be used with ddos action-list",
			},
			"digest": {
				Type: schema.TypeString, Optional: true, Description: "'sha1': sha1; 'sha256': sha256; 'sha384': sha384; 'sha512': sha512;",
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
			"geo_location": {
				Type: schema.TypeString, Optional: true, Description: "Geo-location CSV File",
			},
			"geo_location_archive": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"geo_location_archive_format": {
							Type: schema.TypeString, Optional: true, Description: "'GeoLite2-ASN': GeoLite2-ASN CSV Zipped File; 'GeoLite2-City': GeoLite2-City CSV Zipped File; 'GeoLite2-Country': GeoLite2-Country CSV Zipped File;",
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
			"glm_cert": {
				Type: schema.TypeString, Optional: true, Description: "GLM certificate",
			},
			"glm_license": {
				Type: schema.TypeString, Optional: true, Description: "License File",
			},
			"health_external": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"externalfilename": {
							Type: schema.TypeString, Optional: true, Description: "Specify the Program Name",
						},
						"description": {
							Type: schema.TypeString, Optional: true, Description: "Describe the Program Function briefly",
						},
						"overwrite": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Overwrite existing file",
						},
						"use_mgmt_port": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management port as source port",
						},
						"remote_file": {
							Type: schema.TypeString, Optional: true, Description: "",
						},
						"password": {
							Type: schema.TypeString, Optional: true, Description: "password for the remote site",
						},
					},
				},
			},
			"health_postfile": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"postfilename": {
							Type: schema.TypeString, Optional: true, Description: "Specify the File Name",
						},
						"overwrite": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Overwrite existing file",
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
			"local_uri_file": {
				Type: schema.TypeString, Optional: true, Description: "Local URI files for http response",
			},
			"lw_4o6": {
				Type: schema.TypeString, Optional: true, Description: "LW-4over6 Binding Table File",
			},
			"ng_waf_custom_page": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"custom_page": {
							Type: schema.TypeString, Optional: true, Description: "Custom html file for ng-waf",
						},
						"overwrite": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Overwrite existing file",
						},
						"use_mgmt_port": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management interface for reachability",
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
			"ng_waf_module": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"overwrite": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Overwrite existing file",
						},
						"use_mgmt_port": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management interface for reachability",
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
			"overwrite": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Overwrite existing file",
			},
			"password": {
				Type: schema.TypeString, Optional: true, Description: "password for the remote site",
			},
			"pfx_password": {
				Type: schema.TypeString, Optional: true, Description: "The password for certificate file (pfx type only)",
			},
			"remote_file": {
				Type: schema.TypeString, Optional: true, Description: "profile name for remote url",
			},
			"remote_file_zone_transfer": {
				Type: schema.TypeString, Optional: true, Description: "profile name for remote url",
			},
			"rpz": {
				Type: schema.TypeString, Optional: true, Description: "Response Policy Zone File",
			},
			"secured": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Mark as non-exportable",
			},
			"ssl_cert": {
				Type: schema.TypeString, Optional: true, Description: "SSL Cert File(enter bulk when import an archive file)",
			},
			"ssl_cert_key": {
				Type: schema.TypeString, Optional: true, Description: "'bulk': import an archive file;",
			},
			"ssl_crl": {
				Type: schema.TypeString, Optional: true, Description: "SSL Crl File",
			},
			"ssl_key": {
				Type: schema.TypeString, Optional: true, Description: "SSL Key File(enter bulk when import an archive file)",
			},
			"store": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"delete": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Delete an import store profile",
						},
						"create": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Create an import store profile",
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
				Type: schema.TypeString, Optional: true, Description: "Import store name",
			},
			"terminal": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "terminal vi",
			},
			"thales_kmdata": {
				Type: schema.TypeString, Optional: true, Description: "Thales Kmdata files",
			},
			"thales_secworld": {
				Type: schema.TypeString, Optional: true, Description: "Thales security world files",
			},
			"to_device": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"device": {
							Type: schema.TypeInt, Optional: true, Description: "Device (Device ID)",
						},
						"glm_license": {
							Type: schema.TypeString, Optional: true, Description: "License File",
						},
						"glm_cert": {
							Type: schema.TypeString, Optional: true, Description: "GLM certificate",
						},
						"web_category_license": {
							Type: schema.TypeString, Optional: true, Description: "License file to enable web-category feature",
						},
						"overwrite": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Overwrite existing file",
						},
						"use_mgmt_port": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management port as source port",
						},
						"remote_file": {
							Type: schema.TypeString, Optional: true, Description: "profile name for remote url",
						},
					},
				},
			},
			"tsig": {
				Type: schema.TypeString, Optional: true, Description: "Transaction SIGnatures Key file",
			},
			"usb_license": {
				Type: schema.TypeString, Optional: true, Description: "USB License File",
			},
			"use_mgmt_port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management port as source port",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"web_category_license": {
				Type: schema.TypeString, Optional: true, Description: "License file to enable web-category feature",
			},
			"xml_schema": {
				Type: schema.TypeString, Optional: true, Description: "XML-Schema File",
			},
			"zone_transfer": {
				Type: schema.TypeString, Optional: true, Description: "'zone-transfer': zone-transfer;",
			},
		},
	}
}
func resourceImportCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImport(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceImportRead(ctx, d, meta)
	}
	return diags
}

func resourceImportUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImport(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceImportRead(ctx, d, meta)
	}
	return diags
}
func resourceImportDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImport(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceImportRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceImportRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointImport(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectImportAuthSamlIdp439(d []interface{}) edpt.ImportAuthSamlIdp439 {

	count1 := len(d)
	var ret edpt.ImportAuthSamlIdp439
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.SamlIdpName = in["saml_idp_name"].(string)
		ret.VerifyXmlSignature = in["verify_xml_signature"].(int)
		ret.Overwrite = in["overwrite"].(int)
		ret.UseMgmtPort = in["use_mgmt_port"].(int)
		ret.RemoteFile = in["remote_file"].(string)
		ret.Password = in["password"].(string)
	}
	return ret
}

func getObjectImportGeoLocationArchive440(d []interface{}) edpt.ImportGeoLocationArchive440 {

	count1 := len(d)
	var ret edpt.ImportGeoLocationArchive440
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.GeoLocationArchiveFormat = in["geo_location_archive_format"].(string)
		ret.UseMgmtPort = in["use_mgmt_port"].(int)
		ret.RemoteFile = in["remote_file"].(string)
		ret.Password = in["password"].(string)
	}
	return ret
}

func getObjectImportHealthExternal441(d []interface{}) edpt.ImportHealthExternal441 {

	count1 := len(d)
	var ret edpt.ImportHealthExternal441
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Externalfilename = in["externalfilename"].(string)
		ret.Description = in["description"].(string)
		ret.Overwrite = in["overwrite"].(int)
		ret.UseMgmtPort = in["use_mgmt_port"].(int)
		ret.RemoteFile = in["remote_file"].(string)
		ret.Password = in["password"].(string)
	}
	return ret
}

func getObjectImportHealthPostfile442(d []interface{}) edpt.ImportHealthPostfile442 {

	count1 := len(d)
	var ret edpt.ImportHealthPostfile442
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Postfilename = in["postfilename"].(string)
		ret.Overwrite = in["overwrite"].(int)
		ret.UseMgmtPort = in["use_mgmt_port"].(int)
		ret.RemoteFile = in["remote_file"].(string)
		ret.Password = in["password"].(string)
	}
	return ret
}

func getObjectImportNgWafCustomPage443(d []interface{}) edpt.ImportNgWafCustomPage443 {

	count1 := len(d)
	var ret edpt.ImportNgWafCustomPage443
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CustomPage = in["custom_page"].(string)
		ret.Overwrite = in["overwrite"].(int)
		ret.UseMgmtPort = in["use_mgmt_port"].(int)
		ret.RemoteFile = in["remote_file"].(string)
		ret.Password = in["password"].(string)
	}
	return ret
}

func getObjectImportNgWafModule444(d []interface{}) edpt.ImportNgWafModule444 {

	count1 := len(d)
	var ret edpt.ImportNgWafModule444
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Overwrite = in["overwrite"].(int)
		ret.UseMgmtPort = in["use_mgmt_port"].(int)
		ret.RemoteFile = in["remote_file"].(string)
		ret.Password = in["password"].(string)
	}
	return ret
}

func getObjectImportStore445(d []interface{}) edpt.ImportStore445 {

	count1 := len(d)
	var ret edpt.ImportStore445
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Delete = in["delete"].(int)
		ret.Create = in["create"].(int)
		ret.Name = in["name"].(string)
		ret.RemoteFile = in["remote_file"].(string)
	}
	return ret
}

func getObjectImportToDevice446(d []interface{}) edpt.ImportToDevice446 {

	count1 := len(d)
	var ret edpt.ImportToDevice446
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Device = in["device"].(int)
		ret.GlmLicense = in["glm_license"].(string)
		ret.GlmCert = in["glm_cert"].(string)
		ret.WebCategoryLicense = in["web_category_license"].(string)
		ret.Overwrite = in["overwrite"].(int)
		ret.UseMgmtPort = in["use_mgmt_port"].(int)
		ret.RemoteFile = in["remote_file"].(string)
	}
	return ret
}

func dataToEndpointImport(d *schema.ResourceData) edpt.Import {
	var ret edpt.Import
	ret.Inst.Aflex = d.Get("aflex").(string)
	ret.Inst.AuthJwks = d.Get("auth_jwks").(string)
	ret.Inst.AuthPortal = d.Get("auth_portal").(string)
	ret.Inst.AuthPortalImage = d.Get("auth_portal_image").(string)
	ret.Inst.AuthSamlIdp = getObjectImportAuthSamlIdp439(d.Get("auth_saml_idp").([]interface{}))
	ret.Inst.Background = d.Get("background").(int)
	ret.Inst.BiosFile = d.Get("bios_file").(string)
	ret.Inst.BwList = d.Get("bw_list").(string)
	ret.Inst.CaCert = d.Get("ca_cert").(string)
	ret.Inst.CertificateType = d.Get("certificate_type").(string)
	ret.Inst.ClassList = d.Get("class_list").(string)
	ret.Inst.ClassListConvert = d.Get("class_list_convert").(string)
	ret.Inst.ClassListType = d.Get("class_list_type").(string)
	ret.Inst.CloudConfig = d.Get("cloud_config").(string)
	ret.Inst.CloudCreds = d.Get("cloud_creds").(string)
	ret.Inst.CsrGenerate = d.Get("csr_generate").(int)
	ret.Inst.DdosScript = d.Get("ddos_script").(string)
	ret.Inst.Digest = d.Get("digest").(string)
	ret.Inst.DnssecDnskey = d.Get("dnssec_dnskey").(string)
	ret.Inst.DnssecDs = d.Get("dnssec_ds").(string)
	ret.Inst.DomainList = d.Get("domain_list").(string)
	ret.Inst.GeoLocation = d.Get("geo_location").(string)
	ret.Inst.GeoLocationArchive = getObjectImportGeoLocationArchive440(d.Get("geo_location_archive").([]interface{}))
	ret.Inst.GlmCert = d.Get("glm_cert").(string)
	ret.Inst.GlmLicense = d.Get("glm_license").(string)
	ret.Inst.HealthExternal = getObjectImportHealthExternal441(d.Get("health_external").([]interface{}))
	ret.Inst.HealthPostfile = getObjectImportHealthPostfile442(d.Get("health_postfile").([]interface{}))
	ret.Inst.IpMapList = d.Get("ip_map_list").(string)
	ret.Inst.LocalUriFile = d.Get("local_uri_file").(string)
	ret.Inst.Lw4o6 = d.Get("lw_4o6").(string)
	ret.Inst.NgWafCustomPage = getObjectImportNgWafCustomPage443(d.Get("ng_waf_custom_page").([]interface{}))
	ret.Inst.NgWafModule = getObjectImportNgWafModule444(d.Get("ng_waf_module").([]interface{}))
	ret.Inst.Overwrite = d.Get("overwrite").(int)
	ret.Inst.Password = d.Get("password").(string)
	ret.Inst.PfxPassword = d.Get("pfx_password").(string)
	ret.Inst.RemoteFile = d.Get("remote_file").(string)
	ret.Inst.RemoteFileZoneTransfer = d.Get("remote_file_zone_transfer").(string)
	ret.Inst.Rpz = d.Get("rpz").(string)
	ret.Inst.Secured = d.Get("secured").(int)
	ret.Inst.SslCert = d.Get("ssl_cert").(string)
	ret.Inst.SslCertKey = d.Get("ssl_cert_key").(string)
	ret.Inst.SslCrl = d.Get("ssl_crl").(string)
	ret.Inst.SslKey = d.Get("ssl_key").(string)
	ret.Inst.Store = getObjectImportStore445(d.Get("store").([]interface{}))
	ret.Inst.StoreName = d.Get("store_name").(string)
	ret.Inst.Terminal = d.Get("terminal").(int)
	ret.Inst.ThalesKmdata = d.Get("thales_kmdata").(string)
	ret.Inst.ThalesSecworld = d.Get("thales_secworld").(string)
	ret.Inst.ToDevice = getObjectImportToDevice446(d.Get("to_device").([]interface{}))
	ret.Inst.Tsig = d.Get("tsig").(string)
	ret.Inst.UsbLicense = d.Get("usb_license").(string)
	ret.Inst.UseMgmtPort = d.Get("use_mgmt_port").(int)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	ret.Inst.WebCategoryLicense = d.Get("web_category_license").(string)
	ret.Inst.XmlSchema = d.Get("xml_schema").(string)
	ret.Inst.ZoneTransfer = d.Get("zone_transfer").(string)
	return ret
}
