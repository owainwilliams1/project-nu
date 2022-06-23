# Project Nu

Project-nu's monorepo.

# Project Nu bot

## Development

Firstly, download and run GCloud firestore emulator:

```
gcloud beta emulators firestore start
```

Create a .env file in `bot/` containing:

```
BOT_TOKEN=
GUILD_ID=
PROJECT_ID=utility-ratio-353814
FIRESTORE_EMULATOR_HOST={similar to [::1]:8190 found after starting emulator}
```

Then to run the go files:

```
go run .
```

## Production

Firstly, create a GCP Virtual Machine in Compute with API access, then:

```
sudo apt-get install git
git clone https://github.com/owainwilliams1/project-nu.git
sudo apt install golang-go
export BOT_TOKEN=
```

Create a .env file in `bot/` containing:

```
BOT_TOKEN=
GUILD_ID=
PROJECT_ID=utility-ratio-353814
LOG_NAME=DiscordBot
```

Then to run the go files you must create a service:

`sudo nano /etc/systemd/system/discordbot.service`
```
[Unit]
Description=Discord Bot Go Service
ConditionPathExists=/home/owain/project-nu/bot
After=network.target
[Service]
Type=simple
User=owain
Group=owain
WorkingDirectory=/home/owain/project-nu/bot
ExecStart=/usr/bin/go run .
SyslogIdentifier=discordbotservice
[Install]
WantedBy=multi-user.target
```

Then run these commands to begin the service:

```
sudo systemctl daemon-reload &&
sudo service discordbot start &&
sudo service discordbot status &&
sudo systemctl enable discordbot &&
sudo systemctl start discordbot
```

Checking logs

`journalctl -u discordbot -e`

## Updating production

```
git pull
sudo service discordbot restart
```
