# twitter-login-go

A simple Go package for [Login with Twitter](https://developer.twitter.com/en/docs/authentication/guides/log-in-with-twitter)

## Installation

```bash
go get github.com/pandodao/twitter-login-go
```

## Preparation

- Sign in to [Twitter Developer Dashboard](https://developer.twitter.com/en/portal/dashboard), create a new app.
- And then you can get the `TwitterApiKey` and `TwitterApiSecret` of the app.

## Usage

### 1. Init client

```go
twitterClient := twitter.New(TwitterApiKey, TwitterApiSecret)
```

### 2. Generate the Auth URL

```go
requestUrl, err := twitterClient.GetAuthURL(callbackURL)
if err != nil {
  return err
}
// ask user to visit the requestUrl
```

Twitter will ask the user to confirm the login, and then redirect user to the `callbackURL` with a `oauth_token` and `oauth_verifier` parameter.

```
callbackURL?oauth_token=xxx&oauth_verifier=xxx
```

You must read the `oauth_token` and `oauth_verifier` parameters from the URL, and then use it to get the access token.

### 3. Exchange code for access token

```go
oauthToken := r.URL.Query().Get("oauth_token")
oauthVerifier := r.URL.Query().Get("oauth_verifier")
accessToken, err := twitterClient.GetAccessToken(oauthToken, oauthVerifier)
if err != nil {
  return err
}

user, err := s.twitterClient.Verify()
if err != nil {
  return err
}

// save `user` and `accessToken` for later use
```
