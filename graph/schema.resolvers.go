package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/marklude/go-sovtech/dataloader"
	"github.com/marklude/go-sovtech/graph/generated"
	"github.com/marklude/go-sovtech/graph/model"
	"github.com/spf13/viper"
)

func (r *queryResolver) Peoples(ctx context.Context, first *int) ([]*model.People, error) {
	// endpoints from config
	endpoint := viper.Get("go-sovtech.starwars-api.endpoint").(string)
	//page path from config
	pagePath := viper.Get("go-sovtech.starwars-api.page-path").(string)

	var url string
	if first == nil {
		// build url passing args, endpoint, path, no paging
		url = fmt.Sprintf("%s%s", endpoint, pagePath)
	} else {
		// build url passing args, endpoint, path, page number
		url = fmt.Sprintf("%s%s%d", endpoint, pagePath, *first)
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	jsonData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(jsonData), &r.Results)
	if err != nil {
		return nil, err
	}

	return r.Results.Peoples, nil
}

func (r *queryResolver) PeopleByName(ctx context.Context, name string) (*model.People, error) {
	return dataloader.CtxLoaders(ctx).PeopleByName.Load(name)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
