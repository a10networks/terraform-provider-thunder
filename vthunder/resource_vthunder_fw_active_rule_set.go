package vthunder

//vThunder resource FwActiveRuleSet

import (
    "github.com/go_vthunder/vthunder"
    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
    "util"
    "fmt"
    "strconv"
)

func resourceFwActiveRuleSet() *schema.Resource {
    return &schema.Resource{
        Create: resourceFwActiveRuleSetCreate,
        Update: resourceFwActiveRuleSetUpdate,
        Read:   resourceFwActiveRuleSetRead,
        Delete: resourceFwActiveRuleSetDelete,
         Schema: map[string]*schema.Schema{ 
"session_aging":{
 Type:schema.TypeString,
 Optional: true,
 Description: "",
 }, 
"override_nat_aging":{
 Type:schema.TypeInt,
 Optional: true,
 Description: "",
 }, 
"name":{
 Type:schema.TypeString,
 Optional: true,
 Description: "",
 }, 
"uuid":{
 Type:schema.TypeString,
 Optional: true,
 Description: "",
 }, 
},
    } }
 
func resourceFwActiveRuleSetCreate(d *schema.ResourceData, meta interface{}) error {
     logger := util.GetLoggerInstance()
     client := meta.(vThunder)
      
    if client.Host != "" {
         logger.Println("[INFO] Creating FwActiveRuleSet (Inside resourceFwActiveRuleSetCreate) ")
          
        data := dataToFwActiveRuleSet(d)
         logger.Println("[INFO] received formatted data from method data to FwActiveRuleSet --")
         d.SetId(strconv.Itoa('1'))
         go_vthunder.PostFwActiveRuleSet(client.Token, data, client.Host)
 
        return resourceFwActiveRuleSetRead(d, meta)

    }
     return nil
 }
    
func resourceFwActiveRuleSetRead(d *schema.ResourceData, meta interface{}) error {
     logger := util.GetLoggerInstance()
     client := meta.(vThunder)
     logger.Println("[INFO] Reading FwActiveRuleSet (Inside resourceFwActiveRuleSetRead)")
 
    if client.Host != "" {
         name := d.Id()
        logger.Println("[INFO] Fetching service Read" + name)
         data, err := go_vthunder.GetFwActiveRuleSet(client.Token, client.Host)
         if data == nil {
             logger.Println("[INFO] No data found " + name)
             d.SetId("")
             return nil
         }
        return err
    }
     return nil
 }
    
func resourceFwActiveRuleSetUpdate(d *schema.ResourceData, meta interface{}) error {

    return resourceFwActiveRuleSetRead(d, meta)
}

func resourceFwActiveRuleSetDelete(d *schema.ResourceData, meta interface{}) error {

    return resourceFwActiveRuleSetRead(d, meta)
}
func dataToFwActiveRuleSet(d *schema.ResourceData) go_vthunder.Fw {
    var vc go_vthunder.Fw
    var c go_vthunder.FwActiveRuleSetInstance
    c.SessionAging = d.Get("session_aging").(string)
c.OverrideNatAging = d.Get("override_nat_aging").(int)
c.Name = d.Get("name").(string)

    vc.SessionAging = c 
    return vc
}