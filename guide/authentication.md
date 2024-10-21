# Authentication

During the registration phase, you will have been issued a clientId and clientSecret.
These cannot be used to access the API directly,
but will instead be used as a part of a standard
[OAuth2.0 Authorization Code Grant Flow](https://datatracker.ietf.org/doc/html/rfc6749#section-4.1).

At a high-level, the process will be as follows:

1. A Zentail customer will enable your application, starting the OAuth authorization flow.
1. Then you will complete the OAuth exchange
1. You will be automatically issued a token associated with that particular seller and with the proper scopes that can be used for the Open API

Each Zentail seller that sells on your channel will have a separate token that will need to be associated with their seller account on your channel.

## Detailed Authentication Process Breakdown

This part of the guide will outline the exact process needed to complete the OAuth exchange
and get tokens on behalf of sellers.

### Prerequisites

In order to utilize this authentication flow,
you must first have received a Client ID and Client Secret
from [registering an application](./application-registration.md) with Zentail.

### OAuth Flow

The process for acquiring an Access Token follows the
[OAuth2.0 Authorization Code Grant Flow](https://datatracker.ietf.org/doc/html/rfc6749#section-4.1)
and will be initialized by a Zentail specialist or customer
clicking an "Install" button in the Zentail platform.

Once that button is clicked the process will begin:

1. The Zentail system will verify that your application is active and registered.
2. Zentail will generate a one-time use Authorization Code 
and redirect the user currently installing the application 
to the Redirect URL provided during registration.
The request will be of the format:
```
GET <RedirectURL>/?client_id=<client_id>&code=<authorization code>
```

> [!NOTE]
> This token cannot be used to make API calls, it is used only to securely retrieve the long-lived access token.

3. The application should then have the user log in to their system
(or use their session if they are already logged in).
4. From the backend, the application should send a POST request to https://api.zentail.com/v1/token
with a `Content-type` of `application/x-www-form-urlencoded` and the following data: `grant_type=authorization_code&code=<authorization code>&redirect_uri=<redirect uri>`.

> [!IMPORTANT]
> The redirect URI must match the one provided during registration.

This request needs to be authenticated with basic auth:

* base64 encode the string of `<ClientID>:<ClientSecret>`
* Provide this string in the header: `Authorization: Basic <base64 encoded string>`

If the request is successful, you will receive a response with the Access Token:

```json
{
  "access_token": "<access token>",
  "token_type": "bearer"
}
```

> [!IMPORTANT]
> The access token returned there should be associated with
> the currently authenticated user's account.
> It should be used to make all subsequent requests 
> on behalf of that user and will apply 
> only to the Zentail customer account that started the process.


#### Example Token Exchange Request

Given:

| Field | Value |
|----|----|
| Redirect URI | https://YOUR_DOMAIN/oauth_callback |
| Client ID | 6779ef20e75817b79602 |
| Client Secret | 31408421fb39334a730fc92e41e96986b60f6522a13f631ec93afa6f6c4e0cb3 |
| Authorization Code | l4b6iplj3r |

Base64 Encode the ClientID and Secret:

```bash
$ echo 6779ef20e75817b79602:31408421fb39334a730fc92e41e96986b60f6522a13f631ec93afa6f6c4e0cb3 | base64 -w0
> Njc3OWVmMjBlNzU4MTdiNzk2MDI6MzE0MDg0MjFmYjM5MzM0YTczMGZjOTJlNDFlOTY5ODZiNjBmNjUyMmExM2Y2MzFlYzkzYWZhNmY2YzRlMGNiMwo=
```

Then send this POST request:

```bash
curl -X POST "https://api.zentail.com/v1/token" \
  -H "AUTHORIZATION: Basic Njc3OWVmMjBlNzU4MTdiNzk2MDI6MzE0MDg0MjFmYjM5MzM0YTczMGZjOTJlNDFlOTY5ODZiNjBmNjUyMmExM2Y2MzFlYzkzYWZhNmY2YzRlMGNiMwo=" \
  -d grant_type=authorization_code \
  -d authorization_code=l4b6iplj3r \
  -d redirect_uri=https%3A%2F%2YOUR_DOMAIN%2Eoauth_callback
```

If the application is successfully authorized
the response to that POST request will be a 200 status code with the following JSON body:

```json
{
  "access_token": "v2podKNZ395VqsQ7dz0afA1T9HpfPQkvoJHHo3Vow4U=",
  "token_type": "bearer"
}
```

### Using the Token

The Access Token can now be used to make API calls by providing it in the Authorization header.

For example:

```bash
curl -X GET "https://api.zentail.com/v1/salesOrder?status=PENDING" \
  -H "AUTHORIZATION: v2podKNZ395VqsQ7dz0afA1T9HpfPQkvoJHHo3Vow4U=" \
  -H "accept: application/json"
```

## Next Steps

Now that you have completed the authentication flow,
you should proceed to sending us your 
[Taxonomy Data](./taxonomy-ingestion.md)
so we can begin the process of mapping the taxonomy data.
