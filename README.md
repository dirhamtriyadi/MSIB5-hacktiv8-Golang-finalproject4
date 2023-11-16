# FP4-Hacktiv8

[![forthebadge made-with-go](http://ForTheBadge.com/images/badges/made-with-go.svg)](https://go.dev/)


Berikut ini adalah final project ke-4 dari hacktiv8, aplikasi ini bernama Toko Belanja, Aplikasi ini akan dilengkapi dengan proses CRUD.

## Anggota kelompok
 - Willyawan Maulana - GLNG-KS07-014
 - Dirham Triyadi - GLNG-KS07-025

## Endpoint
Berikut ini adalah seluruh endpoint yang dapat diakses melalui client.

### Users
 
Berikut ini adalah beberapa endpoint yang dapat diakses untuk tabel Users
 
| Method | URL |
| ------ | ------ |
| POST | [https://fp4-hacktiv8-production.up.railway.app/users/register] |
| POST | [https://fp4-hacktiv8-production.up.railway.app/users/login |
| PATCH | [https://fp4-hacktiv8-production.up.railway.app/users/topup] |

###### Daftar request users

POST Register User
 ```sh
{
    "full_name": "string",
    "email": "string",
    "password": "string"
}
```
#
POST Login User
 ```sh
{
    "email": "string",
    "password": "string"
}
```
#

PATCH User Topup

-Bearer Token <br />
 ```sh
{
    "balance": integer
}
```
#

> Note: Untuk method PATCH diperlukan autentikasi, sehingga perlu memasukan bearer token terlebih dahulu. Token didapatkan melalui response client saat melakukan login
#



### Categories
Berikut ini adalah beberapa endpoint yang dapat diakses untuk tabel Categories

| Method | URL |
| ------ | ------ |
| POST | [https://fp4-hacktiv8-production.up.railway.app/categories] |
| GET | [https://fp4-hacktiv8-production.up.railway.app/categories] |
| PATCH | [https://fp4-hacktiv8-production.up.railway.app/categories/:categoryId] |
| DELETE | [https://fp4-hacktiv8-production.up.railway.app/categories/:categoryId] |

###### Daftar request categories

POST Categories

-Bearer Token <br />
 ```sh
{
    "type": "string"
}
```
#

GET Categories 

-Bearer Token <br />

#

PATCH Categories

-Bearer Token <br />
-Param categoryID <br />

 ```sh
{
    "type": "string"
}
```
#
DELETE Categories

-Bearer Token <br />
-Param categoryID <br />

> Note: Seluruh endpoint hanya bisa diakses oleh user dengan role admin, sehingga perlu memasukan bearer token terlebih dahulu. Token didapatkan melalui response client saat melakukan login. Untuk methode PATCH dan DELETE diperlukan parameter Id pada URL

#


### Products
Berikut ini adalah beberapa endpoint yang dapat diakses untuk tabel Products

| Method | URL |
| ------ | ------ |
| POST | [https://fp3-hacktiv8-production.up.railway.app/products] |
| GET | [https://fp3-hacktiv8-production.up.railway.app/products] |
| PUT | [https://fp3-hacktiv8-production.up.railway.app/products/:productId] |
| DELETE | [https://fp3-hacktiv8-production.up.railway.app/products/:productId] |

###### Daftar request Products

POST Products

-Bearer Token <br />

 ```sh
{
    "title": "string",
    "price": integer,
    "stock": integer,
    "category_id": integer
}
```
#

GET Products 

-Bearer Token <br />

#

PUT Products

-Bearer Token <br />
-Param categoryID <br />

 ```sh
{
    "title": "string",
    "price": integer,
    "stock": integer,
    "category_id": integer
}
```

#

DELETE Products

-Bearer Token <br />
-Param categoryID <br />

> Note: Seluruh endpoint kecuali GET hanya bisa diakses oleh user dengan role admin, sehingga perlu memasukan bearer token terlebih dahulu. Token didapatkan melalui response client saat melakukan login. Untuk methode PUT dan DELETE diperlukan parameter Id pada URL. 
#


### TransactionHistories
Berikut ini adalah beberapa endpoint yang dapat diakses untuk tabel TransactionHistories

| Method | URL |
| ------ | ------ |
| POST | [https://fp4-hacktiv8-production.up.railway.app/transactionhistories] |
| GET | [https://fp4-hacktiv8-production.up.railway.app/transactionhistories/my-transactions] |
| GET | [https://fp4-hacktiv8-production.up.railway.app/transactionhistories/user-transactions] 

###### Daftar request TransactionHistories

POST transactions

-Bearer Token <br />

 ```sh
{
    "product_id": integer,
    "quantity": integer
}
```
#

GET my-transactions 

-Bearer Token <br />

#

GET user-transactions

-Bearer Token <br />


> Note: Seluruh endpoint diperlukan autentikasi sehingga diharuskan memasukan bearer token terlebih dahulu. Token didapatkan melalui response client saat melakukan login. Untuk methode GET user-transactions hanya dapat diakses oleh user dengan role admin. 
