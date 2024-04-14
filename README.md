# Running project

```sh
docker compose up --build

## api base url
http://localhost:9000

## Prerequisites
- Make sure you have `curl` installed on your system.

## Usage

### Subscribe Email 

You can set the email address to subscribe
```sh
make subscribe
- type EMAIL="your.email@example.com"

### Unsubscribe Email 

You can set the email address to subscribe
```sh
make unsubscribe
- type EMAIL="your.email@example.com"

### Broadcast Email to subscribe

You can set the email address to subscribe
```sh
make publish
