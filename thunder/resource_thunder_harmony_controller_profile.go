package thunder

//Thunder resource harmony controller profile

import (
	"context"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceHarmonyControllerProfile() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceProfileCreate,
		UpdateContext: resourceProfileUpdate,
		ReadContext:   resourceProfileRead,
		DeleteContext: resourceProfileDelete,
		Schema: map[string]*schema.Schema{
			"thunder_mgmt_ip": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip_address": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"use_mgmt_port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"user_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"provider2": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"host": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"secret_value": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"region": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"availability_zone": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceProfileCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	logger.Println("[INFO] Creating Profile (Inside resourceProfileCreate)")

	if client.Host != "" {
		name := d.Get("user_name").(string)
		vc := dataToProfile(d)
		d.SetId(name)
		err := go_thunder.PostProfile(client.Token, vc, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		return resourceProfileRead(ctx, d, meta)
	}
	return diags
}

func resourceProfileRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	logger.Println("[INFO] Reading Profile (Inside resourceProfileRead)")

	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {

		name := d.Id()

		vc, err := go_thunder.GetProfile(client.Token, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if vc == nil {
			logger.Println("[INFO] No Profile found" + name)
			d.SetId("")
			return nil
		}

		return diags
	}
	return nil
}

func resourceProfileUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceProfileRead(ctx, d, meta)
}

func resourceProfileDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceProfileRead(ctx, d, meta)
}

//Utility method to instantiate Profile Structure
//Configuration for basic required parameters
//TODO: Conf. for all the parameters

func dataToProfile(d *schema.ResourceData) go_thunder.Profile {
	var vc go_thunder.ProfileInstance

	var c go_thunder.Profile

	vc.Host = d.Get("host").(string)
	vc.Port = d.Get("port").(int)
	vc.UserName = d.Get("user_name").(string)
	vc.SecretValue = d.Get("secret_value").(string)
	vc.Provider2 = d.Get("provider2").(string)
	vc.Action = d.Get("action").(string)
	vc.UseMgmtPort = d.Get("use_mgmt_port").(int)
	vc.Region = d.Get("region").(string)
	vc.AvailabilityZone = d.Get("availability_zone").(string)

	var mgmtIp go_thunder.ThunderMgmtIP

	mgmtIp.IPAddress = d.Get("thunder_mgmt_ip.0.ip_address").(string)

	vc.IPAddress = mgmtIp

	c.Host = vc

	return c
}
