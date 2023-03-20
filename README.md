# Rachio Next Run Notifier

Golang program that detects your next Rachio scheduled run and notifies you. It can also send a notification when the run is complete.

This program exists because the author is lazy and needs to be reminded when to bypass the water softener and when to set it back.

## WARNING

The version is not even at the 0.0.1 stage - do not use it!

## Running

### Environmental Variables

#### RACHIO_CONFIG_FILE

Location of your config JSON file.

An example config is located at: `config/local.json`.

## Config Variables

### Pushover.net

Preferred service for sending notifications to your devices.

#### Account Token

primary 'User Key.'

#### Application Token

'API Token/Key' generated for this application.

### Rachio

The Rachio API is rate limited, is tied to your 'Person ID,' and is capped at 1700 per day. So please be a good citizen.

#### bearer_token

The device's API Key is the Bearer Token to authorize HTTP requests to your device. Following these steps to get your key:

https://app.rachio.io -> [select rachio device] -> Account Settings -> GET API KEY

#### device_id

Determined by calls to the Rachio API.

The following commands are run in a bash shell and assume you have `jq` installed:

```bash
export RACHIO_TOKEN="API TOKEN GOES HERE"

# get personal ID
curl -X GET -H "Authorization:Bearer $RACHIO_TOKEN" https://api.rach.io/1/public/person/info | jq .

export RACHIO_PERSONAL_ID="ID from above command goes here"

# get all your data
curl -X GET -H "Authorization:Bearer $RACHIO_TOKEN" https://api.rach.io/1/public/person/$RACHIO_PERSONAL_ID | jq .

export RACHIO_DEVICE="ID found from last command -> devices[0].id"

curl -X GET -H "Content-Type: application/json" -H "Authorization: Bearer $RACHIO_TOKEN" https://api.rach.io/1/public/device/$RACHIO_DEVICE | jq .
```
