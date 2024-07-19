# Cloudvalid Go SDK
Unofficial SDK for working with the [CloudValid APIs](https://docs.cloudvalid.com/api-swagger/).
## Usage
### Create the client
```go
cv := cloudvalid.NewCloudValidClient("CLOUDVALID_API_KEY")
```
### Create new DNS Configuration Link
```go
link, err := cv.CreateDNSConfigurationLink(cloudvalid.CreateDNSConfigurationLinkRequest{
    Domain: "example.com",
    Variables: map[string]string{
        "dmarc": "v=DMARC1; p=quarantine; sp=none; adkim=s; aspf=s;"
    },
    // You can set UseCases and RawDNSRecords as well.
})
// Use link.ID and link.PublicURL
```
### Get previously created DNS Configuration link
```go
page, err := cv.GetHostedPage("LINK_ID")
```
### Get DNS configuration link propagation status
```go
response, err := cv.GetPropagationStatus("LINK_ID")
for _, record := range response.Result {
    if record.Propagated {
        fmt.Printf("%s record for %s is propagated.", record.Type, record.Host)
    } else {
        fmt.Printf("%s record for %s is not propagated: %s", record.Type, record.Host, record.Message)
    }
}
```
### Cancel DNS Configuration Link
```go
response, err := cv.CancelDNSConfigurationLink("LINK_ID")
```