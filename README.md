# Stake ASX Golang Client

stakego is an unofficial golang client for the [Stake](https://hellostake.com) ASX trading platform.

**Note:** This is a pet project that will only be as in-depth as is required to meet my own needs.

## Sample usage

By default, the `NewCredentials()` function will attempt to load your credentials from the following environment variables:
- `STAKE_USERNAME`
- `STAKE_PASSWORD`
- `STAKE_OTP_SECRET` (optional, not recommended)
- `STAKE_SESSION_TOKEN`

`STAKE_SESSION_TOKEN` will take priority over `STAKE_USERNAME` and `STAKE_PASSWORD` if it is set.

`STAKE_OTP_SECRET` can be set to generate your TOTP tokens automatically, but it is recommended to use `STAKE_SESSION_TOKEN` instead.

### Retrieving your session token
If you need to find your session token, you can do that by opening the developer tools of your browser and going to the "Network" tab. Once you've logged in, inspect some of the requests and look for the `Stake-Session-Token` header in the request headers section.

### Login with session token
```
func main() {
	creds := stakego.NewCredentials()
	creds.StakeSessionToken = "48d345ae2321ebb7db112a6745cb7f11"

	c := stakego.NewASXClient()
	c.Credentials = &creds
	_ = c.Login()
	fmt.Printf("Hello %s\n", c.User.FirstName)
}
```


### Login with username and password
```
func main() {
	creds := stakego.NewCredentials()
	creds.Username = "your-username"
	creds.Password = "your-password"

	c := stakego.NewASXClient()
	c.Credentials = &creds
	_ = c.Login()
	fmt.Printf("Hello %s\n", c.User.FirstName)
}
```

### Login with username, password and OTP code
This one isn't recommended but it does work if you want to automate your OTP login
```
func main() {
	creds := stakego.NewCredentials()
	creds.Username = "your-username"
	creds.Password = "your-password"
	creds.OTPSecret = "YOUR2TOTP465ECRET"

	c := stakego.NewASXClient()
	c.Credentials = &creds
	_ = c.Login()
	fmt.Printf("Hello %s\n", c.User.FirstName)
}
```