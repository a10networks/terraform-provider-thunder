package thunder

//Thunder resource RouterBgpAddressFamilyIpv6NetworkSynchronization

import (
	"context"
	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
	"util"
)

func resourceRouterBgpAddressFamilyIpv6NetworkSynchronization() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceRouterBgpAddressFamilyIpv6NetworkSynchronizationCreate,
		UpdateContext: resourceRouterBgpAddressFamilyIpv6NetworkSynchronizationUpdate,
		ReadContext:   resourceRouterBgpAddressFamilyIpv6NetworkSynchronizationRead,
		DeleteContext: resourceRouterBgpAddressFamilyIpv6NetworkSynchronizationDelete,
		Schema: map[string]*schema.Schema{
			"network_synchronization": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"as_number": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"process_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"sequence": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceRouterBgpAddressFamilyIpv6NetworkSynchronizationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating RouterBgpAddressFamilyIpv6NetworkSynchronization (Inside resourceRouterBgpAddressFamilyIpv6NetworkSynchronizationCreate) ")
		name1 := d.Get("as_number").(int)
		logger.Println(name1)
		name := strconv.Itoa(name1)

		logger.Println(name)
		data := dataToRouterBgpAddressFamilyIpv6NetworkSynchronization(d)
		logger.Println("[INFO] received formatted data from method data to RouterBgpAddressFamilyIpv6NetworkSynchronization --")
		d.SetId(strconv.Itoa(name1))
		err := go_thunder.PostRouterBgpAddressFamilyIpv6NetworkSynchronization(client.Token, name, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceRouterBgpAddressFamilyIpv6NetworkSynchronizationRead(ctx, d, meta)

	}
	return diags
}

func resourceRouterBgpAddressFamilyIpv6NetworkSynchronizationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading RouterBgpAddressFamilyIpv6NetworkSynchronization (Inside resourceRouterBgpAddressFamilyIpv6NetworkSynchronizationRead)")

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetRouterBgpAddressFamilyIpv6NetworkSynchronization(client.Token, name1, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if data == nil {
			logger.Println("[INFO] No data found " + name1)
			return nil
		}
		return diags
	}
	return nil
}

func resourceRouterBgpAddressFamilyIpv6NetworkSynchronizationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceRouterBgpAddressFamilyIpv6NetworkSynchronizationRead(ctx, d, meta)
}

func resourceRouterBgpAddressFamilyIpv6NetworkSynchronizationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceRouterBgpAddressFamilyIpv6NetworkSynchronizationRead(ctx, d, meta)
}
func dataToRouterBgpAddressFamilyIpv6NetworkSynchronization(d *schema.ResourceData) go_thunder.RouterBgpAddressFamilyIpv6NetworkSynchronization {
	var vc go_thunder.RouterBgpAddressFamilyIpv6NetworkSynchronization
	var c go_thunder.RouterBgpAddressFamilyIpv6NetworkSynchronizationInstance
	c.NetworkSynchronization = d.Get("network_synchronization").(int)

	vc.NetworkSynchronization = c
	return vc
}
