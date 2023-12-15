package stakego

import (
  "encoding/json"
  "time"
  "github.com/pquerna/otp/totp"
)

// NewCredentials - creates a new Credentials and tries to populate
// it from environment variables
func NewCredentials() Credentials {
  r := Credentials{}
  r.FromEnv()
  return r
}

// Credentials - holds info for creating a UserSession
type Credentials struct {
  Username string `json:"username"`
  Password string `json:"password"`
  OTPSecret string `json:"-"`
  OTP string `json:"otp,omitempty"`
  RememberMeDays int `json:"rememberMeDays"`
  PlatformType string `json:"platformType"`
  StakeSessionToken string `json:"-"`
}

// FromEnv - retrieves login details from the environment
func (r *Credentials) FromEnv() {
  r.Username = GetEnv("STAKE_USERNAME", "")
  r.Password = GetEnv("STAKE_PASSWORD", "")
  r.OTPSecret = GetEnv("STAKE_OTP_SECRET", "")
  r.StakeSessionToken = GetEnv("STAKE_SESSION_TOKEN", "")
  r.RememberMeDays = GetEnvInt("STAKE_REMEMBER_ME_DAYS", 30)
  r.PlatformType = GetEnv("STAKE_PLATFORM_TYPE", "WEB_f5K2x3") // note: find this in the login page
}

// AsJSON - returns properties as JSON byte slice
func (r *Credentials) AsJSON() []byte {
  if r.OTPSecret != "" {
    r.OTP = r.GetOTP()
  }
  j, err := json.Marshal(r)
  if err != nil {
    return []byte("{}")
  }
  return j
}

// GetOTP - uses OTPSecret to get a current OTP token
func (r *Credentials) GetOTP() string {
  if r.OTPSecret != "" {
    o, err := totp.GenerateCode(r.OTPSecret, time.Now())
    if err == nil {
      return o
    }
  }
  return ""
}
