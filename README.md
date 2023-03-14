# media-cloud-vod-open-sdk-go

StreamLake SDK for the Go programming language.

## NOTICE

The golang SDK is still in early staging.

## Convention

For any `Service` like AI, CDN, Metric or VOD, there will be a `_Service_Client`.

For any `API`, there will be an API func for it

```go
func (_Service_Client) API (`_API_Request`) (`_API_Response`, `error`) {
    ...
}
```
