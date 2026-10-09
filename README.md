# Gmail for Rubi

A [Rubi](https://github.com/Deikus-LXXVII/rubi) plugin that lets your AI agent work with Gmail, with
your approval:

- read and search mail (never marks anything as read);
- save drafts to your Gmail Drafts folder;
- send email only after you approve it in the Rubi panel (with your passkey by default);
- optionally get notified when someone replies to an email you sent;
- let your agent watch for new mail from given senders or with given words (say, delivery updates for an
  order) and react to each one;
- keep private mail private: sign-in codes, password resets and anything you choose stay hidden from your
  agent unless you approve showing one;
- choose which folders your agent can see, and whether it may ask for the others;
- connect several Gmail accounts. Your agent picks one by its address; otherwise it uses the default
  (the first one you connected).

## Install

Ask your agent: "Connect my Gmail through Rubi." It installs the plugin from the Rubi store (you approve
what it can do), then sends you a setup link. There you enter your Gmail address and a Google app password:

1. Open [myaccount.google.com/apppasswords](https://myaccount.google.com/apppasswords) (the setup page has
   a button for it) and sign in.
2. Name the app password "Rubi" and press *Create*.
3. Copy the 16 letters into Rubi. Spaces don't matter.

App passwords need 2-Step Verification; if Google says they aren't available, turn it on first. You can
revoke the app password on the same page at any time. Rubi keeps it encrypted; your agent never sees it.
Gmail sign-in through OAuth isn't used because Google only allows it for open-source apps like this one
after a paid security assessment.

To add another account, open *Settings* in the Rubi panel and press *Add account* at Gmail.

## Permissions

| | |
|---|---|
| Asks for | Gmail address, Google app password |
| Read and search mail | no approval by default |
| Save drafts | no approval by default |
| Send email | passkey / password by default |
| Connects to | `imap.gmail.com:993`, `smtp.gmail.com:587` |
| Watch for new mail (wakes your agent) | no approval by default |
| Show a private email | passkey / password, always |
| Give an attachment to your agent (when set to "may ask") | passkey / password, always |
| Open a closed folder for a while | passkey / password, always |
| Can notify your agent about | replies to tracked emails, mail matching its watches |

You can change the approval levels in Rubi's settings, except the two marked "always".

## Your settings

Right after connecting, and later in the Rubi panel under *Settings*, at each account of Gmail:

- **Folders your agent can see** (for each account): all, or only the ones you check; and whether it may
  ask for the others.
- **Hide sign-in codes, one-time passwords and confirmation links:** on by default.
- **Hide password reset emails:** on by default.
- **Hide sign-in and security alerts:** off by default, so your agent can warn you about new sign-ins or
  suspicious activity. (An alert that contains a code stays hidden by the first switch.)
- **Hidden senders** and **hidden words:** anything else you want to keep from your agent.

Formatted emails: your agent reads emails as text, and can ask for the formatted version when it needs a
button the text doesn't show (an "Unsubscribe" link, say). It gets the links and a copy of the email that
loads nothing from the internet, so opening it doesn't tell the sender you read it.

Attachments: per mailbox, choose whether your agent can open the documents and photos attached to emails,
may ask you first (the default), or can't open them at all and sees only how many there are.

The privacy switches apply to all your Gmail accounts. Only you can change these settings, with your
passkey or password; your agent can't see or change them. Hidden mail shows up to your agent only as
"a private email from <sender>".

## Development

The mail engine lives in [rubi-mailkit](https://github.com/Deikus-LXXVII/rubi-mailkit), shared with the
iCloud Mail plugin; this repository only describes Gmail (`main.go`). Tests are in rubi-mailkit.

```bash
go build -o gmail . && ./gmail --manifest
```

To develop against local checkouts, use a workspace (not committed):

```bash
go work init . ../rubi-mailkit ../rubi
```

## Releasing

Tag `vX.Y.Z`. The release workflow tests, builds for linux/amd64, linux/arm64 and darwin/arm64, writes
`rubi-plugin.json`, signs `SHA256SUMS` with the plugin key (`PLUGIN_SIGNING_KEY`; public half in
`publisher-key.pub.pem`), attests the build, and publishes the release. It prints the SHA-256 of
`SHA256SUMS` for the Rubi store catalog.

## License

[MIT](LICENSE)
