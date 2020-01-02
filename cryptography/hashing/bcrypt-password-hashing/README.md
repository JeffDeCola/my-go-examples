# bcrypt-password-hashing

_Store your password hash, not your password.
Set your password hash and then check your password using `bcrypt`._

Documents and reference,

* The genesis for this example
  [here](https://gowebexamples.com/password-hashing/)

## PREREQUISITES

```bash
go get -u -v golang.org/x/crypto/bcrypt
```

## OVERVIEW

Teh beauty of this, your secret and personal password
is never stored on the server.

* You first set it and it's hashed at the server
* When you use it, it's checked against that hash

This illustration may help,

![IMAGE - bcrypt-password-hashing - IMAGE](../../../docs/pics/bcrypt-password-hashing.jpg)

## RUN

```bash
go run bcrypt-password-hashing
```

Now go into the code and change `passwordEntered` and note it's
invalid.
