// Command gmail is the Gmail plugin for Rubi: read, search and draft freely (by default), send only after
// the user's approval, get notified about replies and watched mail, with a privacy filter and folder
// access the user controls. Several Gmail accounts can be connected.
//
// It signs in with a Google app password over IMAP/SMTP: the simplest way that works for every account.
// (Google treats full Gmail access through OAuth as a restricted scope, which an open-source app without
// Google's paid security assessment can only offer to 100 hand-picked test users.)
// The user enters it in the Rubi panel, never in the agent chat, and Rubi keeps it encrypted.
package main

import (
	"strings"

	"github.com/Deikus-LXXVII/rubi-mailkit"
)

var provider = mailkit.Provider{
	ID:          "gmail",
	Name:        "Gmail",
	Version:     "v1.0.0",
	MinRubi:     "v0.6.0",
	Description: "Read, search and draft Gmail; send after your approval; get notified about replies. Several accounts.",
	Needs: "Your Gmail address and a Google app password: open myaccount.google.com/apppasswords, name it " +
		"\"Rubi\" and copy the 16 letters. (Google offers app passwords once 2-Step Verification is on.)",
	Source:     "https://github.com/Deikus-LXXVII/rubi-gmail",
	ToolPrefix: "gmail",

	AddressLabel:       "Gmail address",
	AddressPlaceholder: "name@gmail.com",
	AddressHelp:        "Your Gmail or Google Workspace address.",
	PasswordLabel:      "Google app password",
	PasswordHelp: "Open the link, sign in, name the app password \"Rubi\", and paste the 16 letters here (spaces don't matter). " +
		"If Google says app passwords aren't available, turn on 2-Step Verification first. You can revoke it on the same page at any time.",
	PasswordURL:  "https://myaccount.google.com/apppasswords",
	PasswordLink: "Open Google app passwords",
	// Google shows app passwords as "abcd efgh ijkl mnop"; any spacing is fine.
	CleanPassword: func(s string) string { return strings.Join(strings.Fields(s), "") },

	IMAPAddr:   "imap.gmail.com:993",
	SMTPAddr:   "smtp.gmail.com:587",
	Drafts:     "[Gmail]/Drafts",
	Sent:       "[Gmail]/Sent Mail",
	SavesSent:  true, // Gmail files mail sent over SMTP in Sent Mail by itself
	MailDomain: "gmail.com",
	AuthError: "Google rejected the login. Use a Google app password (16 letters from myaccount.google.com/apppasswords), " +
		"not your Google password; app passwords need 2-Step Verification on. For Google Workspace, your admin may have to allow IMAP",
}

func main() { mailkit.Main(provider) }
