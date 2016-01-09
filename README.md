horenso-reporter-slack
=====================

![image](doc/image.png)

## Usage

```shell
$ go get github.com/Songmu/horenso/cmd/horenso
$ go build
$ HRS_SLACK_TOKEN="YOUR_SLACK_TOKEN" HRS_SLACK_CHANNEL="general" horenso -r ./horenso-reporter-slack -- [command]
$ HRS_SLACK_TOKEN="YOUR_SLACK_TOKEN" HRS_SLACK_GROUP="private_group" horenso -r ./horenso-reporter-slack -- [command]
```

## License

MIT

## Author

[Takuya Arita](https://github.com/ariarijp)
