horenso-reporter-slack
=====================

[![Build Status](https://travis-ci.org/ariarijp/horenso-reporter-slack.svg?branch=master)](https://travis-ci.org/ariarijp/horenso-reporter-slack)

![image](doc/image.png)

This plugin only support [horenso](https://github.com/Songmu/horenso) versions v0.0.1 and higher.

## Usage

```shell
$ go get github.com/ariarijp/horenso-reporter-slack/cmd/horenso-reporter-slack
$ HRS_SLACK_TOKEN="YOUR_SLACK_TOKEN" HRS_SLACK_CHANNEL="general" horenso -r ./horenso-reporter-slack -- [command]
$ HRS_SLACK_TOKEN="YOUR_SLACK_TOKEN" HRS_SLACK_GROUP="private_group" horenso -r ./horenso-reporter-slack -- [command]
```

### Environment Variables

* `HRS_SLACK_TOKEN`(required)
  * Slack API Token
* `HRS_SLACK_CHANNEL`(required when `HRS_SLACK_GROUP` is not provided)
  * Slack channel name
* `HRS_SLACK_GROUP`(required when `HRS_SLACK_CHANNEL` is not provided)
  * Slack private group name
* `HRS_SLACK_MENTION`(optional, default: `channel`)
  * Slack mention(`channel` or `here` is supported)
* `HRS_SLACK_ITEMS`(optional defauls: `all`):
  * Select reporting items by Comma-Separated Values
  * example: Stdout,Stderr,ExitCode
  * supported items
    * `Command`
    * `CommandArgs`
    * `Output`
    * `Stdout`
    * `Stderr`
    * `ExitCode`
    * `Result`
    * `Pid`
    * `StartAt`
    * `EndAt`
    * `Hostname`
    * `SystemTime`
    * `UserTime`

## License

MIT

## Author

[Takuya Arita](https://github.com/ariarijp)
