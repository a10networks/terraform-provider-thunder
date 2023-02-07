package thunder

//Thunder resource SnmpServerEnableTrapsSlb

import (
	"context"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"util"
)

func resourceSnmpServerEnableTrapsSlb() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSnmpServerEnableTrapsSlbCreate,
		UpdateContext: resourceSnmpServerEnableTrapsSlbUpdate,
		ReadContext:   resourceSnmpServerEnableTrapsSlbRead,
		DeleteContext: resourceSnmpServerEnableTrapsSlbDelete,
		Schema: map[string]*schema.Schema{
			"all": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"server_down": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"vip_port_connratelimit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"server_selection_failure": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"service_group_down": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"server_conn_limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"service_group_member_up": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"server_conn_resume": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"service_up": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"service_conn_limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"gateway_up": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"service_group_up": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"application_buffer_limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"vip_connratelimit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"vip_connlimit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"service_group_member_down": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"service_down": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"bw_rate_limit_exceed": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"server_disabled": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"server_up": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"vip_port_connlimit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"vip_port_down": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"bw_rate_limit_resume": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"gateway_down": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"vip_up": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"vip_port_up": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"vip_down": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"service_conn_resume": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceSnmpServerEnableTrapsSlbCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating SnmpServerEnableTrapsSlb (Inside resourceSnmpServerEnableTrapsSlbCreate) ")

		data := dataToSnmpServerEnableTrapsSlb(d)
		logger.Println("[INFO] received formatted data from method data to SnmpServerEnableTrapsSlb --")
		d.SetId("1")
		err := go_thunder.PostSnmpServerEnableTrapsSlb(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceSnmpServerEnableTrapsSlbRead(ctx, d, meta)

	}
	return diags
}

func resourceSnmpServerEnableTrapsSlbRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading SnmpServerEnableTrapsSlb (Inside resourceSnmpServerEnableTrapsSlbRead)")

	if client.Host != "" {
		logger.Println("[INFO] Fetching service Read")
		data, err := go_thunder.GetSnmpServerEnableTrapsSlb(client.Token, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if data == nil {
			logger.Println("[INFO] No data found ")
			return nil
		}
		return diags
	}
	return nil
}

func resourceSnmpServerEnableTrapsSlbUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSnmpServerEnableTrapsSlbRead(ctx, d, meta)
}

func resourceSnmpServerEnableTrapsSlbDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceSnmpServerEnableTrapsSlbRead(ctx, d, meta)
}
func dataToSnmpServerEnableTrapsSlb(d *schema.ResourceData) go_thunder.SnmpServerEnableTrapsSlb {
	var vc go_thunder.SnmpServerEnableTrapsSlb
	var c go_thunder.SnmpServerEnableTrapsSlbInstance
	c.All = d.Get("all").(int)
	c.ServerDown = d.Get("server_down").(int)
	c.VipPortConnratelimit = d.Get("vip_port_connratelimit").(int)
	c.ServerSelectionFailure = d.Get("server_selection_failure").(int)
	c.ServiceGroupDown = d.Get("service_group_down").(int)
	c.ServerConnLimit = d.Get("server_conn_limit").(int)
	c.ServiceGroupMemberUp = d.Get("service_group_member_up").(int)
	c.ServerConnResume = d.Get("server_conn_resume").(int)
	c.ServiceUp = d.Get("service_up").(int)
	c.ServiceConnLimit = d.Get("service_conn_limit").(int)
	c.GatewayUp = d.Get("gateway_up").(int)
	c.ServiceGroupUp = d.Get("service_group_up").(int)
	c.ApplicationBufferLimit = d.Get("application_buffer_limit").(int)
	c.VipConnratelimit = d.Get("vip_connratelimit").(int)
	c.VipConnlimit = d.Get("vip_connlimit").(int)
	c.ServiceGroupMemberDown = d.Get("service_group_member_down").(int)
	c.ServiceDown = d.Get("service_down").(int)
	c.BwRateLimitExceed = d.Get("bw_rate_limit_exceed").(int)
	c.ServerDisabled = d.Get("server_disabled").(int)
	c.ServerUp = d.Get("server_up").(int)
	c.VipPortConnlimit = d.Get("vip_port_connlimit").(int)
	c.VipPortDown = d.Get("vip_port_down").(int)
	c.BwRateLimitResume = d.Get("bw_rate_limit_resume").(int)
	c.GatewayDown = d.Get("gateway_down").(int)
	c.VipUp = d.Get("vip_up").(int)
	c.VipPortUp = d.Get("vip_port_up").(int)
	c.VipDown = d.Get("vip_down").(int)
	c.ServiceConnResume = d.Get("service_conn_resume").(int)

	vc.All = c
	return vc
}
