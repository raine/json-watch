# json-watch

A small cli tool for monitoring JSON data for new items.

```sh
# the first run never emits
echo '[{"id": 1}, {"id": 2}]' | json-watch test --key id
# the second run finds two new items
echo '[{"id": 1}, {"id": 2}, {"id": 3}, {"id": 4}]' | json-watch test --key id
{"id": 3}
{"id": 4}
```

Go rewrite of [json-notify][json-notify].

## install

```sh
go install github.com/raine/json-watch@latest
```

## usage

```sh
<curl json etc.> | json-watch <name>
```

Takes a list of objects as JSON through stdin.

The first execution will "prime" the internal watch file (stored at
`$HOME/.config/json-watch/watches/<name>`) with existing items and won't print
output.

On further executions, unseen JSON objects in the array will be printed to
stdout as newline delimited JSON.

The name parameter uniquely identifies an instance of json-watch usage, so if
watching multiple JSONs for new objects, each of the json-watch calls should
have a distinct name.

If the key parameter is not provided, an object's content is calculated to a
checksum and that is used for identification.

## options

```
  -h, --help         show this help
  -k, --key string   prop in json objects that identifies them (basically the id)
      --version      show installed version
```

## example use

The tool works great combined with crontab, curl, [ramda-cli][ramda-cli] and
[jsonargs][jsonargs].

```sh
curl -s -G --data-urlencode "q=stuff" "https://api.tori.fi/api/v1.1/public/ads" \
  | ramda \
    '.list_ads' \
    'map (.ad)' \
    'filter -> it.type.code is "s" and it.company_ad is false and it.list_price' \
    'map -> { id: it.ad_id, message: it.subject, url: it.share_link, price: it.list_price.price_value }' \
  | json-watch tori-stuff --key id \
  | jsonargs curl -s -G \
    --data-urlencode "chat_id=$chat_id" \
    --data-urlencode "text=[{{.message}}]({{.url}}) {{.price}}€" \
    --data-urlencode "parse_mode=Markdown" \
    "https://api.telegram.org/bot$telegram_bot_token/sendMessage"
```

When running the pipeline periodically with crontab, new items matching search
query on tori.fi are sent to telegram.

[jq]: https://stedolan.github.io/jq/
[ramda-cli]: https://github.com/raine/ramda-cli
[json-notify]: https://github.com/raine/json-notify
[jsonargs]: https://github.com/mattn/jsonargs
