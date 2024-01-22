package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceHsmTemplate() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_hsm_template`: HSM Template\n\n__PLACEHOLDER__",
		CreateContext: resourceHsmTemplateCreate,
		UpdateContext: resourceHsmTemplateUpdate,
		ReadContext:   resourceHsmTemplateRead,
		DeleteContext: resourceHsmTemplateDelete,

		Schema: map[string]*schema.Schema{
			"enroll_timeout": {
				Type: schema.TypeInt, Optional: true, Description: "Specify Enroll Timeout",
			},
			"health_check_interval": {
				Type: schema.TypeInt, Optional: true, Description: "Specify Thales HSM Health Check Interval",
			},
			"hsm_dev": {
				Type: schema.TypeList, Optional: true, Description: "",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hsm_ip": {
							Type: schema.TypeString, Optional: true, Description: "Specify HSM Device IP Address",
						},
						"hsm_port": {
							Type: schema.TypeInt, Optional: true, Description: "Specify Port",
						},
						"hsm_priority": {
							Type: schema.TypeInt, Optional: true, Description: "Specify Priority",
						},
					},
				},
			},
			"password": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify HSM Passphrase",
			},
			"password_string": {
				Type: schema.TypeString, Optional: true, Description: "Password (minimum 4 characters)",
			},
			"protection": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Specify Protection Method",
			},
			"protection_module": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Module",
			},
			"protection_ocs": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Operator Card Set",
			},
			"protection_softcard_hash": {
				Type: schema.TypeString, Optional: true, Description: "Hash",
			},
			"rfs_ip": {
				Type: schema.TypeString, Optional: true, Description: "Specify Thales Remote File System",
			},
			"rfs_port": {
				Type: schema.TypeInt, Optional: true, Description: "Specify Port",
			},
			"sec_world": {
				Type: schema.TypeString, Optional: true, Description: "Security World Name",
			},
			"softcard": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Softcard",
			},
			"softhsm_enum": {
				Type: schema.TypeString, Optional: true, Description: "'softHSM': software implementation of a cryptographic store; 'thalesHSM': Thales HSM;",
			},
			"template_name": {
				Type: schema.TypeString, Required: true, Description: "Specify Template name",
			},
			"user_tag": {
				Type: schema.TypeString, Optional: true, Description: "Customized tag",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"worker": {
				Type: schema.TypeInt, Optional: true, Description: "Specify number of workers for each data CPU",
			},
		},
	}
}
func resourceHsmTemplateCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHsmTemplateCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHsmTemplate(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHsmTemplateRead(ctx, d, meta)
	}
	return diags
}

func resourceHsmTemplateUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHsmTemplateUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHsmTemplate(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceHsmTemplateRead(ctx, d, meta)
	}
	return diags
}
func resourceHsmTemplateDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHsmTemplateDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHsmTemplate(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceHsmTemplateRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceHsmTemplateRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointHsmTemplate(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func getSliceHsmTemplateHsmDev(d []interface{}) []edpt.HsmTemplateHsmDev {

	count1 := len(d)
	ret := make([]edpt.HsmTemplateHsmDev, 0, count1)
	for _, item := range d {
		in := item.(map[string]interface{})
		var oi edpt.HsmTemplateHsmDev
		oi.HsmIp = in["hsm_ip"].(string)
		oi.HsmPort = in["hsm_port"].(int)
		oi.HsmPriority = in["hsm_priority"].(int)
		ret = append(ret, oi)
	}
	return ret
}

func dataToEndpointHsmTemplate(d *schema.ResourceData) edpt.HsmTemplate {
	var ret edpt.HsmTemplate
	//omit encrypted
	ret.Inst.EnrollTimeout = d.Get("enroll_timeout").(int)
	ret.Inst.HealthCheckInterval = d.Get("health_check_interval").(int)
	ret.Inst.HsmDev = getSliceHsmTemplateHsmDev(d.Get("hsm_dev").([]interface{}))
	ret.Inst.Password = d.Get("password").(int)
	ret.Inst.PasswordString = d.Get("password_string").(string)
	ret.Inst.Protection = d.Get("protection").(int)
	ret.Inst.ProtectionModule = d.Get("protection_module").(int)
	ret.Inst.ProtectionOcs = d.Get("protection_ocs").(int)
	ret.Inst.ProtectionSoftcardHash = d.Get("protection_softcard_hash").(string)
	ret.Inst.RfsIp = d.Get("rfs_ip").(string)
	ret.Inst.RfsPort = d.Get("rfs_port").(int)
	ret.Inst.SecWorld = d.Get("sec_world").(string)
	ret.Inst.Softcard = d.Get("softcard").(int)
	ret.Inst.SofthsmEnum = d.Get("softhsm_enum").(string)
	ret.Inst.TemplateName = d.Get("template_name").(string)
	ret.Inst.UserTag = d.Get("user_tag").(string)
	//omit uuid
	ret.Inst.Worker = d.Get("worker").(int)
	return ret
}
