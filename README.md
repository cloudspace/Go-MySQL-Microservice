# Go MySQL Microservice
A MySql Query executer for Go

#### Usage
Input must have 2 arguments:

1.  -connectionURI (refer to https://github.com/go-sql-driver/mysql DSN)
2.  -query

#### Example Input - Output
-
Input:
```
go run main.go -connectionURI="root:password@tcp(127.0.0.1:3306)/adventureworks" -query="CREATE DATABASE southwind"
```
Output (Success):
```
{"result":[]}
```
-
Input:
```
go run main.go -connectionURI="root:password@tcp(127.0.0.1:3306)/adventureworks" -query="CREATE DATABASE southwind"
```
Output (Failure):
```
{
  "error":"Error 1007: Can't create database 'southwind'; database exists"
}
```
-
Input:
```
go run main.go -connectionURI="root:password@tcp(127.0.0.1:3306)/adventureworks" -query="Show Tables"
```

Output (Shortened):

```
{
  "result":[
    {
      "Tables_in_adventureworks":"address"
    },
    {
      "Tables_in_adventureworks":"addresstype"
    },
    {
      "Tables_in_adventureworks":"awbuildversion"
    },
    {
      "Tables_in_adventureworks":"billofmaterials"
    }
  ]
}
```
-
Input:
```
go run main.go -connectionURI="root:password@tcp(127.0.0.1:3306)/adventureworks" -query="SELECT * FROM vendor LIMIT 2"
```

Output:

```
{
  "result":[
    {
      "AccountNumber":"INTERNAT0001",
      "ActiveFlag":"\u0001",
      "CreditRating":"1",
      "ModifiedDate":"2002-02-25 00:00:00",
      "Name":"International",
      "PreferredVendorStatus":"\u0001",
      "PurchasingWebServiceURL":null,
      "VendorID":"1"
    },
    {
      "AccountNumber":"ELECTRON0002",
      "ActiveFlag":"\u0001",
      "CreditRating":"1",
      "ModifiedDate":"2002-02-17 00:00:00",
      "Name":"Electronic Bike Repair \u0026 Supplies",
      "PreferredVendorStatus":"\u0001",
      "PurchasingWebServiceURL":null,
      "VendorID":"2"
    }
  ]
}
```

#### How to build docker image
Requirements:

1. Golang environment set up
2. Git
3. boot2docker running

```
go get https://github.com/cloudspace/Go-MySQL-Microservice.git
cd <Go-MySQL-Microservice directory>/Go-MySQL-Microservice
docker run --rm -v $(pwd):/src centurylink/golang-builder
docker build -t <username>/go-mysql-microservice:0.1 ./

```
