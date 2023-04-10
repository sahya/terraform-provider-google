// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: DCL     ***
//
// ----------------------------------------------------------------------------
//
//     This file is managed by Magic Modules (https://github.com/GoogleCloudPlatform/magic-modules)
//     and is based on the DCL (https://github.com/GoogleCloudPlatform/declarative-resource-client-library).
//     Changes will need to be made to the DCL or Magic Modules instead of here.
//
//     We are not currently able to accept contributions to this file. If changes
//     are required, please file an issue at https://github.com/hashicorp/terraform-provider-google/issues/new/choose
//
// ----------------------------------------------------------------------------

package google

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	dcl "github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
	firebaserules "github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/firebaserules"
)

func ResourceFirebaserulesRuleset() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirebaserulesRulesetCreate,
		Read:   resourceFirebaserulesRulesetRead,
		Delete: resourceFirebaserulesRulesetDelete,

		Importer: &schema.ResourceImporter{
			State: resourceFirebaserulesRulesetImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"source": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: "`Source` for the `Ruleset`.",
				MaxItems:    1,
				Elem:        FirebaserulesRulesetSourceSchema(),
			},

			"project": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      "The project for the resource",
			},

			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Output only. Time the `Ruleset` was created.",
			},

			"metadata": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Output only. The metadata for this ruleset.",
				Elem:        FirebaserulesRulesetMetadataSchema(),
			},

			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Output only. Name of the `Ruleset`. The ruleset_id is auto generated by the service. Format: `projects/{project_id}/rulesets/{ruleset_id}`",
			},
		},
	}
}

func FirebaserulesRulesetSourceSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"files": {
				Type:        schema.TypeList,
				Required:    true,
				ForceNew:    true,
				Description: "`File` set constituting the `Source` bundle.",
				Elem:        FirebaserulesRulesetSourceFilesSchema(),
			},

			"language": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "`Language` of the `Source` bundle. If unspecified, the language will default to `FIREBASE_RULES`. Possible values: LANGUAGE_UNSPECIFIED, FIREBASE_RULES, EVENT_FLOW_TRIGGERS",
			},
		},
	}
}

func FirebaserulesRulesetSourceFilesSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"content": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Textual Content.",
			},

			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "File name.",
			},

			"fingerprint": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Fingerprint (e.g. github sha) associated with the `File`.",
			},
		},
	}
}

func FirebaserulesRulesetMetadataSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"services": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Services that this ruleset has declarations for (e.g., \"cloud.firestore\"). There may be 0+ of these.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func resourceFirebaserulesRulesetCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &firebaserules.Ruleset{
		Source:  expandFirebaserulesRulesetSource(d.Get("source")),
		Project: dcl.String(project),
	}

	id, err := obj.ID()
	if err != nil {
		return fmt.Errorf("error constructing id: %s", err)
	}
	d.SetId(id)
	directive := CreateDirective
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}
	billingProject := project
	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}
	client := NewDCLFirebaserulesClient(config, userAgent, billingProject, d.Timeout(schema.TimeoutCreate))
	if bp, err := ReplaceVars(d, config, client.Config.BasePath); err != nil {
		d.SetId("")
		return fmt.Errorf("Could not format %q: %w", client.Config.BasePath, err)
	} else {
		client.Config.BasePath = bp
	}
	res, err := client.ApplyRuleset(context.Background(), obj, directive...)

	if _, ok := err.(dcl.DiffAfterApplyError); ok {
		log.Printf("[DEBUG] Diff after apply returned from the DCL: %s", err)
	} else if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error creating Ruleset: %s", err)
	}

	if err = d.Set("name", res.Name); err != nil {
		return fmt.Errorf("error setting name in state: %s", err)
	}
	// ID has a server-generated value, set again after creation.

	id, err = res.ID()
	if err != nil {
		return fmt.Errorf("error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Ruleset %q: %#v", d.Id(), res)

	return resourceFirebaserulesRulesetRead(d, meta)
}

func resourceFirebaserulesRulesetRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &firebaserules.Ruleset{
		Source:  expandFirebaserulesRulesetSource(d.Get("source")),
		Project: dcl.String(project),
		Name:    dcl.StringOrNil(d.Get("name").(string)),
	}

	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}
	billingProject := project
	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}
	client := NewDCLFirebaserulesClient(config, userAgent, billingProject, d.Timeout(schema.TimeoutRead))
	if bp, err := ReplaceVars(d, config, client.Config.BasePath); err != nil {
		d.SetId("")
		return fmt.Errorf("Could not format %q: %w", client.Config.BasePath, err)
	} else {
		client.Config.BasePath = bp
	}
	res, err := client.GetRuleset(context.Background(), obj)
	if err != nil {
		resourceName := fmt.Sprintf("FirebaserulesRuleset %q", d.Id())
		return handleNotFoundDCLError(err, d, resourceName)
	}

	if err = d.Set("source", flattenFirebaserulesRulesetSource(res.Source)); err != nil {
		return fmt.Errorf("error setting source in state: %s", err)
	}
	if err = d.Set("project", res.Project); err != nil {
		return fmt.Errorf("error setting project in state: %s", err)
	}
	if err = d.Set("create_time", res.CreateTime); err != nil {
		return fmt.Errorf("error setting create_time in state: %s", err)
	}
	if err = d.Set("metadata", flattenFirebaserulesRulesetMetadata(res.Metadata)); err != nil {
		return fmt.Errorf("error setting metadata in state: %s", err)
	}
	if err = d.Set("name", res.Name); err != nil {
		return fmt.Errorf("error setting name in state: %s", err)
	}

	return nil
}

func resourceFirebaserulesRulesetDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := &firebaserules.Ruleset{
		Source:  expandFirebaserulesRulesetSource(d.Get("source")),
		Project: dcl.String(project),
		Name:    dcl.StringOrNil(d.Get("name").(string)),
	}

	log.Printf("[DEBUG] Deleting Ruleset %q", d.Id())
	userAgent, err := generateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}
	billingProject := project
	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}
	client := NewDCLFirebaserulesClient(config, userAgent, billingProject, d.Timeout(schema.TimeoutDelete))
	if bp, err := ReplaceVars(d, config, client.Config.BasePath); err != nil {
		d.SetId("")
		return fmt.Errorf("Could not format %q: %w", client.Config.BasePath, err)
	} else {
		client.Config.BasePath = bp
	}
	if err := client.DeleteRuleset(context.Background(), obj); err != nil {
		return fmt.Errorf("Error deleting Ruleset: %s", err)
	}

	log.Printf("[DEBUG] Finished deleting Ruleset %q", d.Id())
	return nil
}

func resourceFirebaserulesRulesetImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)

	if err := ParseImportId([]string{
		"projects/(?P<project>[^/]+)/rulesets/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVarsForId(d, config, "projects/{{project}}/rulesets/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func expandFirebaserulesRulesetSource(o interface{}) *firebaserules.RulesetSource {
	if o == nil {
		return firebaserules.EmptyRulesetSource
	}
	objArr := o.([]interface{})
	if len(objArr) == 0 || objArr[0] == nil {
		return firebaserules.EmptyRulesetSource
	}
	obj := objArr[0].(map[string]interface{})
	return &firebaserules.RulesetSource{
		Files:    expandFirebaserulesRulesetSourceFilesArray(obj["files"]),
		Language: firebaserules.RulesetSourceLanguageEnumRef(obj["language"].(string)),
	}
}

func flattenFirebaserulesRulesetSource(obj *firebaserules.RulesetSource) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"files":    flattenFirebaserulesRulesetSourceFilesArray(obj.Files),
		"language": obj.Language,
	}

	return []interface{}{transformed}

}
func expandFirebaserulesRulesetSourceFilesArray(o interface{}) []firebaserules.RulesetSourceFiles {
	if o == nil {
		return make([]firebaserules.RulesetSourceFiles, 0)
	}

	objs := o.([]interface{})
	if len(objs) == 0 || objs[0] == nil {
		return make([]firebaserules.RulesetSourceFiles, 0)
	}

	items := make([]firebaserules.RulesetSourceFiles, 0, len(objs))
	for _, item := range objs {
		i := expandFirebaserulesRulesetSourceFiles(item)
		items = append(items, *i)
	}

	return items
}

func expandFirebaserulesRulesetSourceFiles(o interface{}) *firebaserules.RulesetSourceFiles {
	if o == nil {
		return firebaserules.EmptyRulesetSourceFiles
	}

	obj := o.(map[string]interface{})
	return &firebaserules.RulesetSourceFiles{
		Content:     dcl.String(obj["content"].(string)),
		Name:        dcl.String(obj["name"].(string)),
		Fingerprint: dcl.String(obj["fingerprint"].(string)),
	}
}

func flattenFirebaserulesRulesetSourceFilesArray(objs []firebaserules.RulesetSourceFiles) []interface{} {
	if objs == nil {
		return nil
	}

	items := []interface{}{}
	for _, item := range objs {
		i := flattenFirebaserulesRulesetSourceFiles(&item)
		items = append(items, i)
	}

	return items
}

func flattenFirebaserulesRulesetSourceFiles(obj *firebaserules.RulesetSourceFiles) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"content":     obj.Content,
		"name":        obj.Name,
		"fingerprint": obj.Fingerprint,
	}

	return transformed

}

func flattenFirebaserulesRulesetMetadata(obj *firebaserules.RulesetMetadata) interface{} {
	if obj == nil || obj.Empty() {
		return nil
	}
	transformed := map[string]interface{}{
		"services": obj.Services,
	}

	return []interface{}{transformed}

}
