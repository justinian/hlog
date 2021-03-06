# hlog

`hlog` is a super-simple Hakka Logs CLI client. Just set up your `~/.hakkarc`
file and then run `hlog <log message>` to post a log.

## Config file

hlog looks for its config in `$HOME/.hakkarc`, in the following format:

```
[logs]
token = your_hakka_logs_webhook_token
default-privacy = private
```

Optionally, you can substitute a different URL for the default one by adding a
`url = http://some.custom.domain/path?foo=bar` - just be sure to leave off the
`&token=your_token..` part, as hlog appends this automatically.

The `default-privacy` parameter defaults to `private`, but you can also set
it to `public` or `anonymous`. You can also set this per-log with the
`-priv=<anonymous|public|private>` command-line flag.

If you're having TLS handshake issues, please let me know. As a temporary
workaround, you can try turning off TLS certificate validation by adding the
following to your .hakkarc:

```
[connection]
verify-ssl = false
```

## Installing

If you've got Go installed, just run `go get github.com/justinian/hlog`.
