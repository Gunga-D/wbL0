## Orders service

The main functions that the application performs:
- data storage
- providing the access for data in various ways

So the service allows to work with data by http requests and message queue

### Details

Service supports the following http requests:
- **GET** /api/orders/:id  — returns desired order in json format
- **GET** /view/orders/:id — returns the web page with desired order
- **POST** /api/orders/    — adds to database the order in json format

Also adding an order to database executes by streaming. All information about publishing to channel you can find in test.

### Testing

To run NATS-streaming test:

```bash
make t-ord-serv
```  

### Launching the service

To launch service:
```bash
make serv
```

Or
```bash
docker-compose up -d --build
```