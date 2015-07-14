Locksmith
=========

Locksmith is a simple way to manage SSH keys across your infrastructure for your team. Using GitHub as a source of truth, Locksmith prints out your teams keys into stdout so you can pipe it to wherever you wish.

Typically, you'll want to pipe these keys into an `~/.ssh/authorized_keys` file.

## Usage

Printing keys to stdout:
```
locksmith -access-token 12345 -team-id 12345
ssh-rsa AAAAB3NzaC1yc2EAAA...
```

Piping your teams public keys to your service user account authorized_keys file:
```
locksmith -access-token 12345 -team-id 12345 > ~/.ssh/authorized_keys
ssh-rsa AAAAB3NzaC1yc2EAAA...
```

Install locksmith and run it for your team:
```
wget https://s3.amazonaws.com/adamenger/locksmith --directory-prefix /usr/local/bin/ && locksmith -access-token 12345 -team-id 12345 > ~/.ssh/authorized_keys

```
