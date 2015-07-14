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
```

Install locksmith and run it for your team:
```
sudo wget https://s3.amazonaws.com/adamenger/locksmith --directory-prefix /usr/local/bin/ && chmod +x /usr/local/bin/locksmith && locksmith -access-token 12345 -team-id 12345 > ~/.ssh/authorized_keys
```

So you want to remove a user from your infrastructure? Just delete them from your team in GitHub and run locksmith again:
```
locksmith -access-token 12345 -team-id 12345 > ~/.ssh/authorized_keys
```

## Compiling

If you're compiling on the platform that you're targeting, use this:
```
go build locksmith.go
```
If you're compiling for a different platform you can pass the `GOOS` environment variable to the build command to specify your target platform.

## Teams

In order to get your teams keys, you need to know the Team ID. There's a flag `-get teams` you can pass to locksmith to tell it to grab your teams and id's.
```
locksmith $> ./locksmith -get teams -access-token 12345 -org reverbdotcom
Getting teams for reverbdotcom
Name: Team1, ID: 12345
Name: Team2, ID: 123456
```
