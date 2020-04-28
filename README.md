# timeular-api-client
API Client for Timeular APIv2

Based on the API Docs from Timeular: https://developers.timeular.com/public-api/


Create new API
```
api, err := timeular.NewAPI(
	conf.TimeularBaseURL,
	conf.TimeularKey,
	conf.TimeularSecret,
)
```

