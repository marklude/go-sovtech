package dataloader

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/marklude/go-sovtech/graph/model"
	"github.com/spf13/viper"
)

type ctxKeyType struct{ name string }

var ctxKey = ctxKeyType{"peopleCtx"}

type loaders struct {
	PeopleByName *PeopleLoader
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		ldrs := loaders{}
		ldrs.PeopleByName = &PeopleLoader{
			maxBatch: 100,
			wait:     1 * time.Millisecond,
			fetch: func(keys []string) ([]*model.People, []error) {
				name := keys[0]

				// Endpoints from config
				endpoint := viper.Get("go-sovtech.starwars-api.endpoint").(string)
				// Search path from config
				searchPath := viper.Get("go-sovtech.starwars-api.search-path").(string)

				searchUrl := fmt.Sprintf("%s%s%s", endpoint, searchPath, url.QueryEscape(name))
				errors := make([]error, 2)
				// Get call to Star Wars REST api
				res, err := http.Get(searchUrl)
				if err != nil {
					errors = append(errors, err)
					return nil, errors
				}

				jsonData, err := ioutil.ReadAll(res.Body)
				if err != nil {
					errors = append(errors, err)
					return nil, errors
				}

				var result *model.Results
				err = json.Unmarshal([]byte(jsonData), &result)
				if err != nil {
					errors = append(errors, err)
					return nil, errors
				}
				return result.Peoples, nil
			},
		}
		ctx := context.WithValue(r.Context(), ctxKey, ldrs)
		next.ServeHTTP(rw, r.WithContext(ctx))
	})
}

func CtxLoaders(ctx context.Context) loaders {
	return ctx.Value(ctxKey).(loaders)
}
