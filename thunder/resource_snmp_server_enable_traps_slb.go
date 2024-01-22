package thunder

import (
	"context"
	edpt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSnmpServerEnableTrapsSlb() *schema.Resource {
	return &schema.Resource{
		Description:   "`thunder_snmp_server_enable_traps_slb`: Enable SLB group traps\n\n__PLACEHOLDER__",
		CreateContext: resourceSnmpServerEnableTrapsSlbCreate,
		UpdateContext: resourceSnmpServerEnableTrapsSlbUpdate,
		ReadContext:   resourceSnmpServerEnableTrapsSlbRead,
		DeleteContext: resourceSnmpServerEnableTrapsSlbDelete,

		Schema: map[string]*schema.Schema{
			"all": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable all SLB traps",
			},
			"application_buffer_limit": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable application buffer reach limit trap",
			},
			"bw_rate_limit_exceed": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SLB server/port bandwidth rate limit exceed trap",
			},
			"bw_rate_limit_resume": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SLB server/port bandwidth rate limit resume trap",
			},
			"gateway_down": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SLB server gateway down trap",
			},
			"gateway_up": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SLB server gateway up trap",
			},
			"server_conn_limit": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SLB server connection limit trap",
			},
			"server_conn_resume": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SLB server connection resume trap",
			},
			"server_disabled": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SLB server-disabled trap",
			},
			"server_down": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SLB server-down trap",
			},
			"server_selection_failure": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SLB server selection failure trap",
			},
			"server_up": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable slb server up trap",
			},
			"service_conn_limit": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SLB service connection limit trap",
			},
			"service_conn_resume": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SLB service connection resume trap",
			},
			"service_down": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SLB service-down trap",
			},
			"service_group_down": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SLB service-group-down trap",
			},
			"service_group_member_down": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SLB service-group-member-down trap",
			},
			"service_group_member_up": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SLB service-group-member-up trap",
			},
			"service_group_up": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SLB service-group-up trap",
			},
			"service_up": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SLB service-up trap",
			},
			"uuid": {
				Type: schema.TypeString, Optional: true, Computed: true, Description: "uuid of the object",
			},
			"vip_connlimit": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable the virtual server reach conn-limit trap",
			},
			"vip_connratelimit": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable the virtual server reach conn-rate-limit trap",
			},
			"vip_down": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SLB virtual server down trap",
			},
			"vip_port_connlimit": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable the virtual port reach conn-limit trap",
			},
			"vip_port_connratelimit": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable the virtual port reach conn-rate-limit trap",
			},
			"vip_port_down": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SLB virtual port down trap",
			},
			"vip_port_up": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SLB virtual port up trap",
			},
			"vip_up": {
				Type: schema.TypeInt, Optional: true, Default: 0, Description: "Enable SLB virtual server up trap",
			},
		},
	}
}
func resourceSnmpServerEnableTrapsSlbCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsSlbCreate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsSlb(d)
		d.SetId(obj.GetId())
		err := obj.Post(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSnmpServerEnableTrapsSlbRead(ctx, d, meta)
	}
	return diags
}

func resourceSnmpServerEnableTrapsSlbUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsSlbUpdate()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsSlb(d)
		err := obj.Put(client.Token, client.Host, logger)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceSnmpServerEnableTrapsSlbRead(ctx, d, meta)
	}
	return diags
}
func resourceSnmpServerEnableTrapsSlbDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsSlbDelete()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsSlb(d)
		err := obj.Delete(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceSnmpServerEnableTrapsSlbRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(Thunder)
	logger := client.log
	logger.Println("resourceSnmpServerEnableTrapsSlbRead()")
	var diags diag.Diagnostics
	if client.Host != "" {
		obj := dataToEndpointSnmpServerEnableTrapsSlb(d)
		err := obj.Get(client.Token, client.Host, d.Id(), logger)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func dataToEndpointSnmpServerEnableTrapsSlb(d *schema.ResourceData) edpt.SnmpServerEnableTrapsSlb {
	var ret edpt.SnmpServerEnableTrapsSlb
	ret.Inst.All = d.Get("all").(int)
	ret.Inst.ApplicationBufferLimit = d.Get("application_buffer_limit").(int)
	ret.Inst.BwRateLimitExceed = d.Get("bw_rate_limit_exceed").(int)
	ret.Inst.BwRateLimitResume = d.Get("bw_rate_limit_resume").(int)
	ret.Inst.GatewayDown = d.Get("gateway_down").(int)
	ret.Inst.GatewayUp = d.Get("gateway_up").(int)
	ret.Inst.ServerConnLimit = d.Get("server_conn_limit").(int)
	ret.Inst.ServerConnResume = d.Get("server_conn_resume").(int)
	ret.Inst.ServerDisabled = d.Get("server_disabled").(int)
	ret.Inst.ServerDown = d.Get("server_down").(int)
	ret.Inst.ServerSelectionFailure = d.Get("server_selection_failure").(int)
	ret.Inst.ServerUp = d.Get("server_up").(int)
	ret.Inst.ServiceConnLimit = d.Get("service_conn_limit").(int)
	ret.Inst.ServiceConnResume = d.Get("service_conn_resume").(int)
	ret.Inst.ServiceDown = d.Get("service_down").(int)
	ret.Inst.ServiceGroupDown = d.Get("service_group_down").(int)
	ret.Inst.ServiceGroupMemberDown = d.Get("service_group_member_down").(int)
	ret.Inst.ServiceGroupMemberUp = d.Get("service_group_member_up").(int)
	ret.Inst.ServiceGroupUp = d.Get("service_group_up").(int)
	ret.Inst.ServiceUp = d.Get("service_up").(int)
	//omit uuid
	ret.Inst.VipConnlimit = d.Get("vip_connlimit").(int)
	ret.Inst.VipConnratelimit = d.Get("vip_connratelimit").(int)
	ret.Inst.VipDown = d.Get("vip_down").(int)
	ret.Inst.VipPortConnlimit = d.Get("vip_port_connlimit").(int)
	ret.Inst.VipPortConnratelimit = d.Get("vip_port_connratelimit").(int)
	ret.Inst.VipPortDown = d.Get("vip_port_down").(int)
	ret.Inst.VipPortUp = d.Get("vip_port_up").(int)
	ret.Inst.VipUp = d.Get("vip_up").(int)
	return ret
}
