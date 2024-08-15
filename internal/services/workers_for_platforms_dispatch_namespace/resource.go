// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_for_platforms_dispatch_namespace

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/workers_for_platforms"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/apijson"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.ResourceWithConfigure = &WorkersForPlatformsDispatchNamespaceResource{}
var _ resource.ResourceWithModifyPlan = &WorkersForPlatformsDispatchNamespaceResource{}
var _ resource.ResourceWithImportState = &WorkersForPlatformsDispatchNamespaceResource{}

func NewResource() resource.Resource {
	return &WorkersForPlatformsDispatchNamespaceResource{}
}

// WorkersForPlatformsDispatchNamespaceResource defines the resource implementation.
type WorkersForPlatformsDispatchNamespaceResource struct {
	client *cloudflare.Client
}

func (r *WorkersForPlatformsDispatchNamespaceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_workers_for_platforms_dispatch_namespace"
}

func (r *WorkersForPlatformsDispatchNamespaceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*cloudflare.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"unexpected resource configure type",
			fmt.Sprintf("Expected *cloudflare.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *WorkersForPlatformsDispatchNamespaceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *WorkersForPlatformsDispatchNamespaceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	dataBytes, err := apijson.Marshal(data)
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize http request", err.Error())
		return
	}
	res := new(http.Response)
	env := WorkersForPlatformsDispatchNamespaceResultEnvelope{*data}
	_, err = r.client.WorkersForPlatforms.Dispatch.Namespaces.New(
		ctx,
		workers_for_platforms.DispatchNamespaceNewParams{
			AccountID: cloudflare.F(data.AccountID.ValueString()),
		},
		option.WithRequestBody("application/json", dataBytes),
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	bytes, _ := io.ReadAll(res.Body)
	err = apijson.Unmarshal(bytes, &env)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}
	data = &env.Result
	data.ID = data.NamespaceID

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WorkersForPlatformsDispatchNamespaceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *WorkersForPlatformsDispatchNamespaceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	res := new(http.Response)
	env := WorkersForPlatformsDispatchNamespaceResultEnvelope{*data}
	_, err := r.client.WorkersForPlatforms.Dispatch.Namespaces.Get(
		ctx,
		data.NamespaceID.ValueString(),
		workers_for_platforms.DispatchNamespaceGetParams{
			AccountID: cloudflare.F(data.AccountID.ValueString()),
		},
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	bytes, _ := io.ReadAll(res.Body)
	err = apijson.Unmarshal(bytes, &env)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}
	data = &env.Result
	data.ID = data.NamespaceID

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WorkersForPlatformsDispatchNamespaceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	var data *WorkersForPlatformsDispatchNamespaceModel

	path := strings.Split(req.ID, "/")
	if len(path) != 2 {
		resp.Diagnostics.AddError("Invalid ID", "expected urlencoded segments <account_id>/<dispatch_namespace>")
		return
	}
	path_account_id, err := url.PathUnescape(path[0])
	if err != nil {
		resp.Diagnostics.AddError("invalid urlencoded segment - <account_id>", fmt.Sprintf("%s -> %q", err.Error(), path[0]))
	}
	path_dispatch_namespace, err := url.PathUnescape(path[1])
	if err != nil {
		resp.Diagnostics.AddError("invalid urlencoded segment - <dispatch_namespace>", fmt.Sprintf("%s -> %q", err.Error(), path[1]))
	}
	if resp.Diagnostics.HasError() {
		return
	}

	res := new(http.Response)
	env := WorkersForPlatformsDispatchNamespaceResultEnvelope{*data}
	_, err = r.client.WorkersForPlatforms.Dispatch.Namespaces.Get(
		ctx,
		path_dispatch_namespace,
		workers_for_platforms.DispatchNamespaceGetParams{
			AccountID: cloudflare.F(path_account_id),
		},
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	bytes, _ := io.ReadAll(res.Body)
	err = apijson.Unmarshal(bytes, &env)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}
	data = &env.Result
	data.ID = data.NamespaceID

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WorkersForPlatformsDispatchNamespaceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *WorkersForPlatformsDispatchNamespaceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	var state *WorkersForPlatformsDispatchNamespaceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	dataBytes, err := apijson.MarshalForUpdate(data, state)
	if err != nil {
		resp.Diagnostics.AddError("failed to serialize http request", err.Error())
		return
	}
	res := new(http.Response)
	env := WorkersForPlatformsDispatchNamespaceResultEnvelope{*data}
	_, err = r.client.WorkersForPlatforms.Dispatch.Namespaces.New(
		ctx,
		workers_for_platforms.DispatchNamespaceNewParams{
			AccountID: cloudflare.F(data.NamespaceID.ValueString()),
		},
		option.WithRequestBody("application/json", dataBytes),
		option.WithResponseBodyInto(&res),
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	bytes, _ := io.ReadAll(res.Body)
	err = apijson.Unmarshal(bytes, &env)
	if err != nil {
		resp.Diagnostics.AddError("failed to deserialize http request", err.Error())
		return
	}
	data = &env.Result
	data.ID = data.NamespaceID

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WorkersForPlatformsDispatchNamespaceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *WorkersForPlatformsDispatchNamespaceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.WorkersForPlatforms.Dispatch.Namespaces.Delete(
		ctx,
		data.NamespaceID.ValueString(),
		workers_for_platforms.DispatchNamespaceDeleteParams{
			AccountID: cloudflare.F(data.AccountID.ValueString()),
		},
		option.WithMiddleware(logging.Middleware(ctx)),
	)
	if err != nil {
		resp.Diagnostics.AddError("failed to make http request", err.Error())
		return
	}
	data.ID = data.NamespaceID

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *WorkersForPlatformsDispatchNamespaceResource) ModifyPlan(_ context.Context, _ resource.ModifyPlanRequest, _ *resource.ModifyPlanResponse) {

}