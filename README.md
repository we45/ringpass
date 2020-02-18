# Ringpass - A trivially simple Password Manager, backed by Keyrings

> This is really meant as a golang learning project. Please use accordingly :)

## Compatibility

* Works on *Nix
* Works on OSX

### Dependencies

#### Linux

The Linux implementation depends on the [Secret Service][SecretService] dbus
interface, which is provided by [GNOME Keyring](https://wiki.gnome.org/Projects/GnomeKeyring).

It's expected that the default collection `login` exists in the keyring, because
it's the default in most distros. If it doesn't exist, you can create it through the
keyring frontend program [Seahorse](https://wiki.gnome.org/Apps/Seahorse):

 * Open `seahorse`
 * Go to **File > New > Password Keyring**
 * Click **Continue**
 * When asked for a name, use: **login**

## Features
* Set Secret (stores in keyring, keychain)
* Retrieve secret to clipboard and JSON (stdout) for things like JQ, etc or use in automated scripts

## Usage

[![asciicast](https://asciinema.org/a/sQWpHC1LoFefLAiQ3PDq4nLry.svg)](https://asciinema.org/a/sQWpHC1LoFefLAiQ3PDq4nLry)

* Download binary for your OS from the releases to `/usr/local/bin` or other exec directory

### Store secrets

```bash
ringpass set --key hello --service aws 
Enter value: 
```

* Has the ability to handle `getpasswd()` style password prompt

OR

```bash
ringpass set --key hello --service aws --value supersecret
2020/02/18 07:51:25 Successfully stored Key, Value and Service in Keyring
```
* Reads string literals from stdin

> Short-codes for flags also available

```bash
ringpass set -k hello -s aws -v supersecret
```

## Retrieve secret

### Copy to clipboard

```bash
ringpass get -c -s aws -k hello
2020/02/18 07:55:27 Successfully written secret to system clipboard
```
* Copies to clipboard

### Write to JSON - stdout

```bash
ringpass get -j -s aws -k hello
{"service":"aws","key":"hello","value":"supersecret"}
```

* Additional cmdline-fu

```bash
ringpass get -j -s aws -k hello | jq -r .value
```