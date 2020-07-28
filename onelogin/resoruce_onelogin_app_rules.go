package onelogin

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/onelogin/onelogin-go-sdk/pkg/client"
	"github.com/onelogin/terraform-provider-onelogin/ol_schema/app_rules"
	"github.com/onelogin/terraform-provider-onelogin/ol_schema/app_rules/actions"
	"github.com/onelogin/terraform-provider-onelogin/ol_schema/app_rules/conditions"
)

// Apps returns a resource with the CRUD methods and Terraform Schema defined
func AppRules() *schema.Resource {
	return &schema.Resource{
		Create:   appRuleCreate,
		Read:     appRuleRead,
		Update:   appRuleUpdate,
		Delete:   appRuleDelete,
		Importer: &schema.ResourceImporter{},
		Schema:   apprulesschema.Schema(),
	}
}

func appRuleCreate(d *schema.ResourceData, m interface{}) error {
	appRule := apprulesschema.Inflate(map[string]interface{}{
		"name":       d.Get("name"),
		"app_id":     d.Get("app_id"),
		"match":      d.Get("match"),
		"enabled":    d.Get("enabled"),
		"position":   d.Get("position"),
		"conditions": d.Get("conditions"),
		"actions":    d.Get("actions"),
	})
	client := m.(*client.APIClient)
	resp, err := client.Services.AppRulesV2.Create(&appRule)
	if err != nil {
		log.Println("[ERROR] There was a problem creating the app rule!", err)
		return err
	}
	log.Printf("[CREATED] Created app rule with %d", *(resp.ID))

	d.SetId(fmt.Sprintf("%d", *(resp.ID)))
	return appRuleRead(d, m)
}

func appRuleRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*client.APIClient)
	rid, _ := strconv.Atoi(d.Id())
	appID := d.Get("app_id").(int)
	appRule, err := client.Services.AppRulesV2.GetOne(int32(appID), int32(rid))
	if err != nil {
		log.Printf("[ERROR] There was a problem reading the app rule!")
		log.Println(err)
		return err
	}
	if appRule == nil {
		d.SetId("")
		return nil
	}
	log.Printf("[READ] Reading app rule with %d", *(appRule.ID))

	d.Set("name", appRule.Name)
	d.Set("app_id", appID)
	d.Set("match", appRule.Match)
	d.Set("enabled", appRule.Enabled)
	d.Set("position", appRule.Position)
	d.Set("conditions", appruleconditionsschema.Flatten(appRule.Conditions))
	d.Set("actions", appruleactionsschema.Flatten(appRule.Actions))

	return nil
}

func appRuleUpdate(d *schema.ResourceData, m interface{}) error {
	appRule := apprulesschema.Inflate(map[string]interface{}{
		"name":       d.Get("name"),
		"app_id":     d.Get("app_id"),
		"match":      d.Get("match"),
		"enabled":    d.Get("enabled"),
		"position":   d.Get("position"),
		"conditions": d.Get("conditions"),
		"actions":    d.Get("actions"),
	})
	rid, _ := strconv.Atoi(d.Id())
	client := m.(*client.APIClient)
	resp, err := client.Services.AppRulesV2.Update(int32(rid), &appRule)
	if err != nil {
		log.Println("[ERROR] There was a problem updating the app rule!", err)
		return err
	}
	log.Printf("[UDATED] Updated app rule with %d", *(resp.ID))

	d.SetId(fmt.Sprintf("%d", *(resp.ID)))
	return appRuleRead(d, m)
}

func appRuleDelete(d *schema.ResourceData, m interface{}) error {
	aid, _ := strconv.Atoi(d.Id())
	appID := d.Get("app_id").(int)

	client := m.(*client.APIClient)
	err := client.Services.AppRulesV2.Destroy(int32(appID), int32(aid))
	if err != nil {
		log.Printf("[ERROR] There was a problem deleting the app rule!")
		log.Println(err)
	} else {
		log.Printf("[DELETED] Deleted app rule with %d", aid)
		d.SetId("")
	}

	return nil
}
