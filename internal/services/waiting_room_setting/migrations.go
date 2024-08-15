// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waiting_room_setting

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.ResourceWithUpgradeState = &WaitingRoomSettingResource{}

func (r *WaitingRoomSettingResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{}
}