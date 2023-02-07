package thunder

//Thunder resource RouterBgpNeighborEthernetNeighbor

import (
	"context"
	"strconv"
	"strings"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterBgpNeighborEthernetNeighbor() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceRouterBgpNeighborEthernetNeighborCreate,
		UpdateContext: resourceRouterBgpNeighborEthernetNeighborUpdate,
		ReadContext:   resourceRouterBgpNeighborEthernetNeighborRead,
		DeleteContext: resourceRouterBgpNeighborEthernetNeighborDelete,
		Schema: map[string]*schema.Schema{
			"ethernet": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"unnumbered": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
			},
			"peer_group_name": {
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

func resourceRouterBgpNeighborEthernetNeighborCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating RouterBgpNeighborEthernetNeighbor (Inside resourceRouterBgpNeighborEthernetNeighborCreate) ")
		name1 := d.Get("as_number").(int)
		name2 := d.Get("ethernet").(int)
		name := strconv.Itoa(name1)
		data := dataToRouterBgpNeighborEthernetNeighbor(d)
		logger.Println("[INFO] received formatted data from method data to RouterBgpNeighborEthernetNeighbor --")
		d.SetId(strconv.Itoa(name1) + "," + strconv.Itoa(name2))
		err := go_thunder.PostRouterBgpNeighborEthernetNeighbor(client.Token, name, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceRouterBgpNeighborEthernetNeighborRead(ctx, d, meta)

	}
	return diags
}

func resourceRouterBgpNeighborEthernetNeighborRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading RouterBgpNeighborEthernetNeighbor (Inside resourceRouterBgpNeighborEthernetNeighborRead)")

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name1 := id[0]
		name2 := id[1]
		logger.Println("[INFO] Fetching service Read" + name1)
		data, err := go_thunder.GetRouterBgpNeighborEthernetNeighbor(client.Token, name1, name2, client.Host)
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

func resourceRouterBgpNeighborEthernetNeighborUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name1 := id[0]
		name2 := id[1]
		logger.Println("[INFO] Modifying RouterBgpNeighborEthernetNeighbor   (Inside resourceRouterBgpNeighborEthernetNeighborUpdate) ")
		data := dataToRouterBgpNeighborEthernetNeighbor(d)
		logger.Println("[INFO] received formatted data from method data to RouterBgpNeighborEthernetNeighbor ")
		err := go_thunder.PutRouterBgpNeighborEthernetNeighbor(client.Token, name1, name2, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceRouterBgpNeighborEthernetNeighborRead(ctx, d, meta)

	}
	return diags
}

func resourceRouterBgpNeighborEthernetNeighborDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		id := strings.Split(d.Id(), ",")
		name1 := id[0]
		name2 := id[1]
		logger.Println("[INFO] Deleting instance (Inside resourceRouterBgpNeighborEthernetNeighborDelete) " + name1)
		err := go_thunder.DeleteRouterBgpNeighborEthernetNeighbor(client.Token, name1, name2, client.Host)
		if err != nil {
			logger.Printf("[ERROR] Unable to Delete resource instance  (%s) (%v)", name1, err)
			return diag.FromErr(err)
		}
		return nil
	}
	return diags
}

func dataToRouterBgpNeighborEthernetNeighbor(d *schema.ResourceData) go_thunder.RouterBgpNeighborEthernetNeighbor {
	var vc go_thunder.RouterBgpNeighborEthernetNeighbor
	var c go_thunder.RouterBgpNeighborEthernetNeighborInstance
	c.Ethernet = d.Get("ethernet").(int)
	c.Unnumbered = d.Get("unnumbered").(int)
	c.PeerGroupName = d.Get("peer_group_name").(string)

	vc.Ethernet = c
	return vc
}
