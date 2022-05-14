## sawolabs/go-sdk


## Install
With Go toolchain:

```
go get -u github.com/sawolabs/go-sdk
```

## Examples
Let's start registering a sawo instance with apiKey and URL path:
```
func main() {
  sawoconfig := new(gosdk.SawoConfig)
	// Identifier type can be one of 'email' or 'phone_number_sms' or 'both_email_phone'
	sawoconfig.Init("your-api-key", "email", "./assets/login.html", "/login")
	r := gosdk.SawoRouter()
	http.ListenAndServe(":8080", r)
}
```
