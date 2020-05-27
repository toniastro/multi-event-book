# Multi-event Booking  (Vue + Go)

[![N|Solid](https://res.cloudinary.com/siteat/image/upload/v1590527203/Screenshot_2020-05-26_at_22.02.12_fi5zbr.png)](https://try.toniastro.com)


This application was built on [Golang 1.14.3](https://golang.org/), [Vue](https://vuejs.org/) and Vibes. It's basic functions is to enable you purchase tickets for events uploaded and get your receipt in PDF immediately after payment. 

  - Payment Integration : [Rave](https://ravepay.co/) by Flutterwave
  - Deployed with Docker 
  - Magic

### Installation

This application required setting up on a local enviroment for  [Golang](https://golang.org/) and creating a local database .

Clone the repo into your local and set up enviroment

```sh
$ git clone https://github.com/iamt-chadwick/multi-event-book.git
$ cd multi-event-book
$ cp .env.example .env
```
Update .env with the needed variables

| VARIABLE | *data needed |
| ------ | ------ |
| DB_PORT | your database port |
| DB_HOST | your database host |
| DB_NAME | name of empty database created |
| PORT | the port you want your application would be served on  |
| RAVE_PUBLIC_KEY | You would get this from your [Rave Dashboard](https://dashboard.flutterwave.com/) |
| RAVE_SECRET_KEY | You would get this from your [Rave Dashboard](https://dashboard.flutterwave.com/)|
| RAVE_MODE | this should be (live or test) |
| RAVE_API_TEST | https://ravesandboxapi.flutterwave.com/flwv3-pug/getpaidx/api/v2/verify |
| RAVE_API_TEST | https://api.ravepay.co/flwv3-pug/getpaidx/api/v2/verify |

### Run Application
 Run command within the multi-event-book directory you cloned
 
 ```sh
$ go run main.go
```

### Todos

 - Write Tests
 - Make a couple tweaks to improve performance

### Preview of Reciept
![N|Solid](https://res.cloudinary.com/siteat/image/upload/v1590539507/Screenshot_2020-05-27_at_01.28.00_mdykha.png)

License
----



MIT


****
