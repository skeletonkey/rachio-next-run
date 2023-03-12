# Rachio Next Run Notifier

Golang program that detects your the run of your next Rachio scheduled run and notifies you. It can also notify you when the run is complete.

This exists because the author is lazy and needs to be reminder when to bypass the water softener and when to set it back.

## WARNING

This is not even at the 0.0.1 stage - do not use!

## Running

### Env Vars

#### RACHIO_CONFIG_FILE

Location of your config JSON file.

Example config is located at: `config/local.json`.

## Config Variables

### Pushover.net

Preferred service to sending notifications to your devices.

### Account Token

This is your main 'User Key.'

### Application Token

This is the 'API Token/Key' that is generated for this application.

### Rachio

The Rachio API is rate limited, appears to be tied to your 'Person ID', and appears to be capped at 1700 per day. Please be a good citizen.

#### bearer_token

The Bearer Token to authorize HTTP requests to your device is the device's API Key.  It can be found with the following steps:

https://app.rachio.io -> [select rachio device] -> Account Settings -> GET API KEY

#### device_id

The Device ID is determined by calls to the Rachio API.

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
