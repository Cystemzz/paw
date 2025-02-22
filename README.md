<div align="center">
    <img alt="Paw" src="logo/paw.png" height="128" />
</div>

Modified version of paw to assist in privacy, all outbound connections have been removed.

# Paw

Paw is a cross platform application to manage your passwords and identities securely.

It is written in Go and uses [Fyne](https://github.com/fyne-io/fyne) as UI toolkit and [age](https://github.com/FiloSottile/age) as encryption library.

## Warning

**This software is work in progress, and didn't went through a full security audit.**

**Do not expect it to be bug free and do not rely on it for any type of security.**

## Screenshot

<div align="center">
    <img alt="Paw screenshot" src="screenshot.png" />
</div>

## Main features

* Cross platform application (linux, macOS, Windows, BSD ...) with a single codebase
* Desktop and CLI application
* Minimal direct dependencies
* Open source: code can be audited
* Audit passwords against data breaches
* TOTP support
* Password import/export

### Later goals

* Automatically detect and use password rules for known web sites that require ones
* Automatic backup / syncronization
* Mobile / Web applications
* Stateless password derivation support
* Unicode password support

## Installation

### Latest version

#### Desktop application

```
go install lucor.dev/paw/cmd/paw@latest
```

#### CLI application

```
go install lucor.dev/paw/cmd/paw-cli@latest
```

### Development version

To try the development version or help with testing:


```
# Desktop application
go install lucor.dev/paw/cmd/paw@develop

# CLI application
go install lucor.dev/paw/cmd/paw-cli@develop
```

## How it works - cryptography details

### Vault initialization

One or more vaults can be initialized to store passwords and identities.

When the vault is initialized user will be prompt for a vault name and password.
An [age](https://github.com/FiloSottile/age) key is generated and it is encrypted using an age Scrypt recipient with the provided password and saved on disk (key.age)
The X25519 identity and its recipient from the key file are used to decrypt and encrypt the vault data.
Each item is stored separately on disk so that the content can be decrypted manually using the age tool, if needed.
All the items' metadata are encrypted and stored into the vault.age file so that no information are in clear text.

### Random password

Random password are derived reading byte-by-byte the block of randomness from a [HKDF](https://pkg.go.dev/golang.org/x/crypto/hkdf) cryptographic key derivation function that uses the age key as secret. Printable characters that match the desired password rule (uppercase, lowercase, symbols and digits) are then included in the generated password.

### Custom password

Where a generated password is not applicable a custom password can be specified. 

### Vault structure

Vault internally is organized hierarchically like:
```
- vault
    ├── login
    |    └── www.example.com
    |    └── my.site.com
    ├── password
    |    └── mypassword
    └── note
         └── mysecretnote
```

where login, password and note are the Paw items, see the dedicated section for details.

### Items

Items are special templates aim to help the identity management.

Currently the following items are available:

- login
- note
- password
- ssh_key

## Threat model

The threat model of Paw assumes there are no attackers on your local machine.

## Contribute

- Fork and clone the repository
- Make and test your changes
- Open a pull request against the `develop` branch

## Contributors

See [contributors](https://github.com/lucor/paw/graphs/contributors) page

## Credits

 - [age](https://github.com/FiloSottile/age) for the encryption library
 - [Fyne](https://github.com/fyne-io/fyne) for the UI toolkit

Also thanks to these Open Source password managers that inspired Paw:

- [gopass](https://github.com/gopasspw/gopass)
- [lesspass](https://github.com/lesspass/lesspass)
- [pass](https://www.passwordstore.org/)
- [passage](https://github.com/FiloSottile/passage)
- [passgo](https://github.com/ejcx/passgo)
