Locksmith
=========

Locksmith is a simple way to manage SSH keys across your infrastructure for your team. Using GitHub as a source of truth, Locksmith prints out your teams keys into stdout so you can pipe it to wherever you wish.

Typically, you'll want to pipe these keys into `~/.ssh/authorized_keys`.

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

## Compiling

```
go build locksmith.go
```

## Teams

In order to get your teams keys, you need to know the Team ID. There's another script in this repo called `get-teams.go`. Build it with `go build get-teams.go` and use it to get the ID of your team.
```
locksmith $> ./get-teams -access-token 12345 -org reverbdotcom
Getting teams for reverbdotcom
Name: Team1, ID: 12345
Name: Team2, ID: 123456
```
