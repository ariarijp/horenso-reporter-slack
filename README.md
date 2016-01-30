horenso-reporter-slack
=====================

[![Build Status](https://travis-ci.org/ariarijp/horenso-reporter-slack.svg?branch=master)](https://travis-ci.org/ariarijp/horenso-reporter-slack)

![image](doc/image.png)

This plugin only support [horenso](https://github.com/Songmu/horenso) versions v0.0.1 and higher.

## Usage

```shell
$ go get https://github.com/ariarijp/horenso-reporter-slack.git
$ cd horenso-reporter-slack
$ make
$ HRS_SLACK_TOKEN="YOUR_SLACK_TOKEN" HRS_SLACK_CHANNEL="general" horenso -r ./horenso-reporter-slack -- [command]
$ HRS_SLACK_TOKEN="YOUR_SLACK_TOKEN" HRS_SLACK_GROUP="private_group" horenso -r ./horenso-reporter-slack -- [command]
```

## License

MIT

## Author

[Takuya Arita](https://github.com/ariarijp)
