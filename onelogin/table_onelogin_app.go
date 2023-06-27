package onelogin

import (
	"context"

	// mod "github.com/onelogin/onelogin-go-sdk/internal/models"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableOneloginApp(_ context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "onelogin_app",
		Description: "Get a list of an account's registered, connected applications.",
		List: &plugin.ListConfig{
			Hydrate: listAuthorizedApps,
		},
		// Get: &plugin.GetConfig{
		// 	KeyColumns: plugin.AllColumns([]string{"id"}),
		// 	Hydrate:    getAuthorizedApp,
		// },
		Columns: []*plugin.Column{
			// {
			// 	Name:        "id",
			// 	Description: "The ID for the application.",
			// 	Type:        proto.ColumnType_INT,
			// 	Transform:   transform.FromField("ID"),
			// },
			// {
			// 	Name:        "name",
			// 	Description: "The name of the application.",
			// 	Type:        proto.ColumnType_STRING,
			// },
			// {
			// 	Name:        "description",
			// 	Description: "A short description of the application.",
			// 	Type:        proto.ColumnType_STRING,
			// },

			// JSON fields
			{
				Name:        "raw",
				Description: "An array of usernames for users who have linked the app.",
				Transform:   transform.FromValue(),
				Type:        proto.ColumnType_JSON,
			},

			// Standard Steampipe columns
			// {
			// 	Name:        "title",
			// 	Description: "The title of the resource.",
			// 	Type:        proto.ColumnType_STRING,
			// 	Transform:   transform.FromField("Name"),
			// },
		},
	}
}

//// LIST FUNCTION

func listAuthorizedApps(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	// Create client
	client, err := connectOnelogin(ctx, d)
	if err != nil {
		logger.Error("onelogin_authorized_app.listAuthorizedApps", "connection_error", err)
		return nil, err
	}

	// Limiting the results
	maxLimit := int32(1000)
	if d.QueryContext.Limit != nil {
		limit := int32(*d.QueryContext.Limit)
		if limit < maxLimit {
			maxLimit = limit
		}
	}

	// last := 0

	for {
		apps, err := client.GetApps({})
		if err != nil {
			logger.Error("onelogin_authorized_app.listAuthorizedApps", "api_error", err)
			return nil, err
		}

		d.StreamListItem(ctx, apps)
		// for _, list := range apps.Apps {

		// 	// Context can be cancelled due to manual cancellation or the limit has been hit
		// 	if d.RowsRemaining(ctx) == 0 {
		// 		return nil, nil
		// 	}
		// }

		// last = params.Offset + len(apps.Apps)
		// if last >= apps.TotalItems {
		// 	return nil, nil
		// } else {
		// 	params.Offset = last
		// }
	}

}

//// HYDRATE FUNCTIONS

// func getAuthorizedApp(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
// 	logger := plugin.Logger(ctx)

// 	id := d.EqualsQuals["id"].GetInt64Value()

// 	// List id should not be empty
// 	if id == 0 {
// 		return nil, nil
// 	}

// 	// Create client
// 	client, err := connectOnelogin(ctx, d)
// 	if err != nil {
// 		logger.Error("onelogin_authorized_app.getAuthorizedApp", "connection_error", err)
// 		return nil, err
// 	}

// 	params := gochimp3.BasicQueryParams{}
// 	if d.EqualsQuals["status"] != nil && d.EqualsQualString("status") != "" {
// 		params.Status = d.EqualsQualString("status")
// 	}

// 	list, err := client.GetAuthroizedApp(strconv.Itoa(int(id)), &params)
// 	if err != nil {
// 		logger.Error("onelogin_authorized_app.getAuthorizedApp", "api_error", err)
// 		return nil, err
// 	}

// 	return list, nil
// }
