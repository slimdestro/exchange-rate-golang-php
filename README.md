# exchange-rate-golang-php
#### Exchange rate of last 10 days:Backend(Golang) | UI(PHP/HTML) | Deployment(Docker + K8 + TF)

<p>
  <a href="https://www.modcode.dev/">
    <img src="https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Go_Logo_Blue.svg/60px-Go_Logo_Blue.svg.png" height="50" alt="Go">
  </a>
  <a href="https://www.modcode.dev/">
    <img src="https://juancenteno.info/wp-content/uploads/2017/02/php.png" height="50" alt="PHP">
  </a>
</p>


## Loom shoot

[![Watch the video](https://static-00.iconduck.com/assets.00/loom-icon-512x155-uq8gnrp3.png)](https://www.loom.com/share/c02cf6d2b2694751af7caa4961165381?sid=0ea56c1b-bb60-45a2-b06a-a18fa883c034)

## Setup 

```sh
./go run backend/exchangeServer.go
```

```sh
Backend:http://localhost:8080
Frontend:http://localhost

APIs:
- Sync : /[GET]http://localhost:8080/syncRates
- Fetch : /[GET]http://localhost:8080/frontend
```

## Run as container:

```sh
docker-compose up -d
```

Stop container:

```sh
docker-compose down  
```

## Run on kubernetes: 
```sh
cd kubemenifests
kubectl apply -f{all 4 files one by one}
this will deploy and create servicve for both app
```

## Via Terraform: 
```sh
theres a Terraform folder in root. just need to run main.tf
terraform init 
terraform apply 
```

## Author

[Checkout my Website](https://www.modcode.dev)
