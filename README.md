# Go MySQL Microservice
A MySql Query executer for Go

#### Usage
Input must have 6 arguments:

1.  -username
2.  -password
3.  -ip
4.  -port
5.  -databasename
6.  -query

#### Example Input - Output
Input:
```
go run main.go -username=root -password=password -ip=127.0.0.1 -port=3306 -databasename=adventureworks -query="CREATE DATABASE southwind"
```
Output (Success):
```
[]
```

Input:
```
go run main.go -username=root -password=password -ip=127.0.0.1 -port=3306 -databasename=adventureworks -query="CREATE DATABASE southwind"
```
Output (Failure):
```
{"error":"Error 1007: Can't create database 'southwind'; database exists"}
```

Input:
```
go run main.go -username=root -password=password -ip=127.0.0.1 -port=3306 -databasename=adventureworks -query="Show Tables"
```

Output:

```
[{"Tables_in_adventureworks":"address"},{"Tables_in_adventureworks":"addresstype"},{"Tables_in_adventureworks":"awbuildversion"},{"Tables_in_adventureworks":"billofmaterials"}]
```

Input:
```
go run main.go -username=root -password=password -ip=127.0.0.1 -port=3306 -databasename=adventureworks -query="SELECT * FROM vendor LIMIT 2"
```

Output:

```
[{"AccountNumber":"INTERNAT0001","ActiveFlag":"\u0001","CreditRating":"1","ModifiedDate":"2002-02-25 00:00:00","Name":"International","PreferredVendorStatus":"\u0001","PurchasingWebServiceURL":null,"VendorID":"1"},{"AccountNumber":"ELECTRON0002","ActiveFlag":"\u0001","CreditRating":"1","ModifiedDate":"2002-02-17 00:00:00","Name":"Electronic Bike Repair \u0026 Supplies","PreferredVendorStatus":"\u0001","PurchasingWebServiceURL":null,"VendorID":"2"}]
```