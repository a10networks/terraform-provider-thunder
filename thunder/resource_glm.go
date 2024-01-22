package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGlm() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_glm`: Set GLM Connection values\n\n__PLACEHOLDER__",
		CreateContext: resourceGlmCreate,
		UpdateContext: resourceGlmUpdate,
		ReadContext:   resourceGlmRead,
		DeleteContext: resourceGlmDelete,

		Schema: map[string]*schema.Schema{
			"allocate_bandwidth": {
				Type: schema.TypeInt, Optional: true, Description: "Enter the requested bandwidth in Mbps for Capacity Pool",
			},
			"appliance_name": {
				Type: schema.TypeString, Optional: true, Description: "Helpful identifier for this appliance",
			},
			"burst": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable Burst License",
			},
			"check_expiration": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
			},
			"create_license_request": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"create_license_request": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Create a GLM trial or license request",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"enable_requests": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Turn on periodic GLM license requests (default license retrieval interval is every 24 hours)",
			},
			"enterprise": {
				Type: schema.TypeString, Optional: true, Description: "Enter the ELM hostname, IP or [IPV6]",
			},
			"enterprise_ha_host_list": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"host_entry": {
							Type: schema.TypeString, Required: true, Description: "Enter the ELM hostname, IP or [IPV6]",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"enterprise_request_type": {
				Type: schema.TypeString, Optional: true, Description: "'fqdn': TLS verified with FQDN; 'self-signed': TLS verified with self signed certificate(Default); 'self-signed-pull-cert': Request and use self signed certificate;",
			},
			"host": {
				Type: schema.TypeString, Optional: true, Description: "",
			},
			"interval": {
				Type: schema.TypeInt, Optional: true, Description: "GLM license request interval (in hours)",
			},
			"new_license": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"existing_org": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use existing account with organization ID",
						},
						"org_id": {
							Type: schema.TypeInt, Optional: true, Description: "GLM organization id",
						},
						"existing_user": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use an existing account with email and password",
						},
						"glm_email": {
							Type: schema.TypeString, Optional: true, Description: "GLM email",
						},
						"glm_password": {
							Type: schema.TypeString, Optional: true, Description: "GLM password",
						},
						"new_user": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Create a new account",
						},
						"new_email": {
							Type: schema.TypeString, Optional: true, Description: "GLM email",
						},
						"new_password": {
							Type: schema.TypeString, Optional: true, Description: "GLM password",
						},
						"account_name": {
							Type: schema.TypeString, Optional: true, Description: "Account Name",
						},
						"first_name": {
							Type: schema.TypeString, Optional: true, Description: "First Name",
						},
						"last_name": {
							Type: schema.TypeString, Optional: true, Description: "Last Name",
						},
						"country": {
							Type: schema.TypeString, Optional: true, Description: "Country",
						},
						"phone": {
							Type: schema.TypeString, Optional: true, Description: "Phone",
						},
						"name": {
							Type: schema.TypeString, Optional: true, Description: "License name (Configure license name)",
						},
						"type": {
							Type: schema.TypeString, Optional: true, Description: "'webroot': webroot; 'webroot_trial': webroot_trial; 'webroot_ti': webroot_ti; 'webroot_ti_trial': webroot_ti_trial; 'qosmos': qosmos; 'qosmos_trial': qosmos_trial; 'ipsec_vpn': ipsec_vpn;",
						},
					},
				},
			},
			"port": {
				Type: schema.TypeInt, Optional: true, Description: "License request port (default 443)",
			},
			"proxy_server": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"host": {
							Type: schema.TypeString, Optional: true, Description: "Proxy server hostname or IP address",
						},
						"port": {
							Type: schema.TypeInt, Optional: true, Description: "Proxy server port",
						},
						"username": {
							Type: schema.TypeString, Optional: true, Description: "Username for proxy authentication",
						},
						"password": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Password for proxy authentication",
						},
						"secret_string": {
							Type: schema.TypeString, Optional: true, Description: "password value",
						},
						"uuid": {
							Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
						},
					},
				},
			},
			"send": {
				Type: schema.TypeList, MaxItems: 1, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"license_request": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Immediately send a single GLM license request",
						},
						"ha_status": {
							Type: schema.TypeInt, Optional: true, Default: 0, Description: "Send a ELM HA status request",
						},
					},
				},
			},
			"thunder_capacity_license": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "",
			},
			"token": {
				Type: schema.TypeString, Optional: true, Description: "License entitlement token",
			},
			"use_mgmt_port": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Use management port to connect to GLM",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
		},
	}
}
func resourceGlmCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGlmCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGlm(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGlmRead(ctx, d, meta)
	}
	return diags
}

func resourceGlmUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGlmUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGlm(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceGlmRead(ctx, d, meta)
	}
	return diags
}
func resourceGlmDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGlmDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGlm(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceGlmRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceGlmRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointGlm(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getObjectGlmCreateLicenseRequest380(d []interface{}) edpt.GlmCreateLicenseRequest380 {

	count1 := len(d)
	var ret edpt.GlmCreateLicenseRequest380
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.CreateLicenseRequest = in["create_license_request"].(int)
		//omit uuid
	}
	return ret
}

func getSliceGlmEnterpriseHaHostList(d []interface{}) []edpt.GlmEnterpriseHaHostList {

	count1 := len(d)
	ret := make([]edpt.GlmEnterpriseHaHostList, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.GlmEnterpriseHaHostList
		oi.HostEntry = in["host_entry"].(string)
		//omit uuid
		ret = append(ret, oi)
	}
	return ret
}

func getObjectGlmNewLicense381(d []interface{}) edpt.GlmNewLicense381 {

	count1 := len(d)
	var ret edpt.GlmNewLicense381
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.ExistingOrg = in["existing_org"].(int)
		ret.OrgId = in["org_id"].(int)
		ret.ExistingUser = in["existing_user"].(int)
		ret.GlmEmail = in["glm_email"].(string)
		ret.GlmPassword = in["glm_password"].(string)
		ret.NewUser = in["new_user"].(int)
		ret.NewEmail = in["new_email"].(string)
		ret.NewPassword = in["new_password"].(string)
		ret.AccountName = in["account_name"].(string)
		ret.FirstName = in["first_name"].(string)
		ret.LastName = in["last_name"].(string)
		ret.Country = in["country"].(string)
		ret.Phone = in["phone"].(string)
		ret.Name = in["name"].(string)
		ret.Type = in["type"].(string)
	}
	return ret
}

func getObjectGlmProxyServer382(d []interface{}) edpt.GlmProxyServer382 {

	count1 := len(d)
	var ret edpt.GlmProxyServer382
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.Host = in["host"].(string)
		ret.Port = in["port"].(int)
		ret.Username = in["username"].(string)
		ret.Password = in["password"].(int)
		ret.SecretString = in["secret_string"].(string)
		//omit encrypted
		//omit uuid
	}
	return ret
}

func getObjectGlmSend383(d []interface{}) edpt.GlmSend383 {

	count1 := len(d)
	var ret edpt.GlmSend383
	if count1 > 0 {
		in := d[0].(map[string]interface{})
		ret.LicenseRequest = in["license_request"].(int)
		ret.HaStatus = in["ha_status"].(int)
	}
	return ret
}

func dataToEndpointGlm(d *schema.ResourceData) edpt.Glm {
	var ret edpt.Glm
	ret.Inst.AllocateBandwidth = d.Get("allocate_bandwidth").(int)
	ret.Inst.ApplianceName = d.Get("appliance_name").(string)
	ret.Inst.Burst = d.Get("burst").(int)
	ret.Inst.CheckExpiration = d.Get("check_expiration").(int)
	ret.Inst.CreateLicenseRequest = getObjectGlmCreateLicenseRequest380(d.Get("create_license_request").([]interface{}))
	ret.Inst.EnableRequests = d.Get("enable_requests").(int)
	ret.Inst.Enterprise = d.Get("enterprise").(string)
	ret.Inst.EnterpriseHaHostList = getSliceGlmEnterpriseHaHostList(d.Get("enterprise_ha_host_list").([]interface{}))
	ret.Inst.EnterpriseRequestType = d.Get("enterprise_request_type").(string)
	ret.Inst.Host = d.Get("host").(string)
	ret.Inst.Interval = d.Get("interval").(int)
	ret.Inst.NewLicense = getObjectGlmNewLicense381(d.Get("new_license").([]interface{}))
	ret.Inst.Port = d.Get("port").(int)
	ret.Inst.ProxyServer = getObjectGlmProxyServer382(d.Get("proxy_server").([]interface{}))
	ret.Inst.Send = getObjectGlmSend383(d.Get("send").([]interface{}))
	ret.Inst.ThunderCapacityLicense = d.Get("thunder_capacity_license").(int)
	ret.Inst.Token = d.Get("token").(string)
	ret.Inst.UseMgmtPort = d.Get("use_mgmt_port").(int)
	//omit uuid
	return ret
}
