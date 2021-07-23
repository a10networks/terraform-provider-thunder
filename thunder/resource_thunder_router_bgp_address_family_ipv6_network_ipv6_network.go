package thunder

//Thunder resource RouterBgpAddressFamilyIpv6NetworkIpv6Network

import (
	"context"
	"strconv"
	"strings"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterBgpAddressFamilyIpv6NetworkIpv6Network() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceRouterBgpAddressFamilyIpv6NetworkIpv6NetworkCreate,
		UpdateContext: resourceRouterBgpAddressFamilyIpv6NetworkIpv6NetworkUpdate,
		ReadContext:   resourceRouterBgpAddressFamilyIpv6NetworkIpv6NetworkRead,
		DeleteContext: resourceRouterBgpAddressFamilyIpv6NetworkIpv6NetworkDelete,
		Schema: map[string]*schema.Schema{
			"network_ipv6": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"route_map": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"backdoor": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"comm_value": {
				Type:        schema.TypeString,
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

func resourceRouterBgpAddressFamilyIpv6NetworkIpv6NetworkCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating RouterBgpAddressFamilyIpv6NetworkIpv6Network (Inside resourceRouterBgpAddressFamilyIpv6NetworkIpv6NetworkCreate) ")
		name1 := d.Get("as_number").(int)
		name2 := d.Get("network_ipv6").(string)
		name := strconv.Itoa(name1)
		data := dataToRouterBgpAddressFamilyIpv6NetworkIpv6Network(d)
		logger.Println("[INFO] received formatted data from method data to RouterBgpAddressFamilyIpv6NetworkIpv6Network --")
		d.SetId(strconv.Itoa(name1) + "," + name2)
		err := go_thunder.PostRouterBgpAddressFamilyIpv6NetworkIpv6Network(client.Token, name, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceRouterBgpAddressFamilyIpv6NetworkIpv6NetworkRead(ctx, d, meta)

	}
	return diags
}

func resourceRouterBgpAddressFamilyIpv6NetworkIpv6NetworkRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading RouterBgpAddressFamilyIpv6NetworkIpv6Network (Inside resourceRouterBgpAddressFamilyIpv6NetworkIpv6NetworkRead)")

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name1 := id[0]
		name2 := id[1]
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetRouterBgpAddressFamilyIpv6NetworkIpv6Network(client.Token, name1, name2, client.Host)
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

func resourceRouterBgpAddressFamilyIpv6NetworkIpv6NetworkUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name1 := id[0]
		name2 := id[1]
		logger.Println("[INFO] Modifying RouterBgpAddressFamilyIpv6NetworkIpv6Network   (Inside resourceRouterBgpAddressFamilyIpv6NetworkIpv6NetworkUpdate) ")
		data := dataToRouterBgpAddressFamilyIpv6NetworkIpv6Network(d)
		logger.Println("[INFO] received formatted data from method data to RouterBgpAddressFamilyIpv6NetworkIpv6Network ")
		err := go_thunder.PutRouterBgpAddressFamilyIpv6NetworkIpv6Network(client.Token, name1, name2, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceRouterBgpAddressFamilyIpv6NetworkIpv6NetworkRead(ctx, d, meta)

	}
	return diags
}

func resourceRouterBgpAddressFamilyIpv6NetworkIpv6NetworkDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name1 := id[0]
		name2 := id[1]
		logger.Println("[INFO] Deleting instance (Inside resourceRouterBgpAddressFamilyIpv6NetworkIpv6NetworkDelete) " + name1)
		err := go_thunder.DeleteRouterBgpAddressFamilyIpv6NetworkIpv6Network(client.Token, name1, name2, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return diag.FromErr(err)
		}
		return nil
	}
	return diags
}

func dataToRouterBgpAddressFamilyIpv6NetworkIpv6Network(d *schema.ResourceData) go_thunder.RouterBgpAddressFamilyIpv6NetworkIpv6Network {
	var vc go_thunder.RouterBgpAddressFamilyIpv6NetworkIpv6Network
	var c go_thunder.RouterBgpAddressFamilyIpv6NetworkIpv6NetworkInstance
	c.NetworkIpv6 = d.Get("network_ipv6").(string)
	c.RouteMap = d.Get("route_map").(string)
	c.Backdoor = d.Get("backdoor").(int)
	c.Description = d.Get("description").(string)
	c.CommValue = d.Get("comm_value").(string)

	vc.NetworkIpv6 = c
	return vc
}
