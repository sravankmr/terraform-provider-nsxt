/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: MPL-2.0 */

package nsxt

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNsxtPolicyMacDiscoveryProfile() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNsxtPolicyMacDiscoveryProfileRead,

		Schema: map[string]*schema.Schema{
			"id":           getDataSourceIDSchema(),
			"display_name": getDataSourceDisplayNameSchema(),
			"description":  getDataSourceDescriptionSchema(),
			"path":         getPathSchema(),
			"context":      getContextSchema(false, false, false),
		},
	}
}

func dataSourceNsxtPolicyMacDiscoveryProfileRead(d *schema.ResourceData, m interface{}) error {
	_, err := policyDataSourceResourceRead(d, getPolicyConnector(m), getSessionContext(d, m), "MacDiscoveryProfile", nil)
	if err != nil {
		return err
	}
	return nil
}
