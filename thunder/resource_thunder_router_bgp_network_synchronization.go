package thunder

//Thunder resource RouterBgpNetworkSynchronization

import (
	"context"
	"strconv"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterBgpNetworkSynchronization() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceRouterBgpNetworkSynchronizationCreate,
		UpdateContext: resourceRouterBgpNetworkSynchronizationUpdate,
		ReadContext:   resourceRouterBgpNetworkSynchronizationRead,
		DeleteContext: resourceRouterBgpNetworkSynchronizationDelete,
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

func resourceRouterBgpNetworkSynchronizationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating RouterBgpNetworkSynchronization (Inside resourceRouterBgpNetworkSynchronizationCreate) ")
		name1 := d.Get("as_number").(int)
		name := strconv.Itoa(name1)
		data := dataToRouterBgpNetworkSynchronization(d)
		logger.Println("[INFO] received formatted data from method data to RouterBgpNetworkSynchronization --")
		d.SetId(strconv.Itoa(name1))
		err := go_thunder.PostRouterBgpNetworkSynchronization(client.Token, name, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceRouterBgpNetworkSynchronizationRead(ctx, d, meta)

	}
	return diags
}

func resourceRouterBgpNetworkSynchronizationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading RouterBgpNetworkSynchronization (Inside resourceRouterBgpNetworkSynchronizationRead)")

	if client.Host != "" {
		name1 := d.Id()
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetRouterBgpNetworkSynchronization(client.Token, name1, client.Host)
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

func resourceRouterBgpNetworkSynchronizationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceRouterBgpNetworkSynchronizationRead(ctx, d, meta)
}

func resourceRouterBgpNetworkSynchronizationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceRouterBgpNetworkSynchronizationRead(ctx, d, meta)
}
func dataToRouterBgpNetworkSynchronization(d *schema.ResourceData) go_thunder.RouterBgpNetworkSynchronization {
	var vc go_thunder.RouterBgpNetworkSynchronization
	var c go_thunder.RouterBgpNetworkSynchronizationInstance
	c.NetworkSynchronization = d.Get("network_synchronization").(int)

	vc.NetworkSynchronization = c
	return vc
}
