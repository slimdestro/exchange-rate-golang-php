# exchange-rate-golang-php
#### Create an application that will show the exchange rates of USD, EUR and GBP for the last 10 days

[![N|Solid](https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Go_Logo_Blue.svg/60px-Go_Logo_Blue.svg.png)](https://dev.to/slimdestro)
 
## Setup

Start container:

```sh
docker-compose up -d
```

Stop container:

```sh
docker-compose down  
```

## Example

```sh
Backend:http://localhost:8080
Frontend:http://localhost

APIs:
- Sync : /[GET]http://localhost:8080/syncRates
- Fetch : /[GET]http://localhost:8080/frontend
```


## Author

[slimdestro(Mukul Mishra)](https://linktr.ee/slimdestro)