# Rachio Next Run Notifier

Golang program that detects your next Rachio scheduled run and notifies you. It can also send a notification when the run is complete.

This program exists because the author is lazy and needs to be reminded when to bypass the water softener and when to set it back.

At this time this program is a 'script' meant to be run via a cron. Future evolution is a 'service' that will simply run and log appropriately.

## WARNING

The version is 0.1.0 - it works, but things will change!

"before" alerts work; after alters probably don't

## Setup

### Environmental Variables

#### RACHIO_CONFIG_FILE

Location of your config JSON file.

An example config is located at: `config/local.json`.

## Config Variables

At this time 'secrets' are written directly into the config file.  It is _HIGHLY_ recommended that you make a copy of `config/local.json` and place it somewhere outside of the repo.

### Pushover.net

Preferred service for sending notifications to your devices.

#### Account Token

primary 'User Key.'

#### Application Token

'API Token/Key' generated for this application.

### Rachio

The Rachio API is rate limited, is tied to your 'Person ID,' and is capped at 1700 per day. So please be a good citizen.

_IMPORTANT_: at this time only one rachio device is supported.  Support for multiple devices will be added at a later date.

#### bearer_token

The device's API Key is the Bearer Token to authorize HTTP requests to your device. Following these steps to get your key:

<https://app.rachio.io/login> -> [select rachio device] -> Account Settings -> GET API KEY

#### device_id

Determined by calls to the Rachio API.

The following commands are run in a bash shell and assumes that `jq` is installed:

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

## Building

Build this on the machine you plan on running it on - or alter the commands with the appropriate platform args.

NOTE: secrets configs will be addressed at a later time.  You will need to update config/local.json AND DO NOT CHECK IT BACK IN!!!

```bash
make
```

## Running

Print out the App config data, Rachio next run information as well as 'alert' data, and any Pushover information.

```bash
./bin/rachion-next-run
```
