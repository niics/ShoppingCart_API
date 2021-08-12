# ShoppingCart_API
Challenges project to build an API for product management and shopping cart


## Tech stack
 - Golang 1.16.7
 - Gin
 - SQLite


## Features
1. API for create a new product with the following fields: name, description, price **=> DONE**
2. API for show a listing of the products with pagination **=> DONE, response as JSON**
3. API for create an order in shopping cart **=> DONE**

	a. An order belongs to a user, and has multiple items. A user can have many orders. An order has a status column, which is an enum: **=> DONE**
	
	**PENDING** - In shopping cart
	
	**PURCHASED** - Purchased
	
	b. For any given user, you can only have one order with a status equal to PENDING **=> WORKING**
	
	c. When a user adds an item to the shopping cart, if there is no order with a status equal to PENDING, then create a new order for the user. **=> WORKING**
4. API for delete items from the shopping cart
    
    a. Only order status equal to PENDING are allow to delete **=> NOT START**
    
5. API for show a listing of item in shopping cart
    
    a. Shows the total price of all items => **DONE, show only list and quantity of order**
    
6. API for adjust quantity of item in shopping cart
    
    a. Minimum quantity are 0 and maximum quantity are 99  **=> NOT START**
    
    b. If user set item quantity to 0, that’s mean we need to remove that item  **=> NOT START**
    
## How to run
- To build this project
```
go build
```
- To run the project
```
go run .
```

## API
API is configurated to be access wiht URL and port below:
```
URL: localhost
port: 8080
```
The list of providing API:
- (GET) To get the product list as JSON
```
{URL}/products-json
```
- (GET) To get the product detail by ID as JSON
```
{URL}/product/:id
```
- (POST) To add the new product and fill the quantity 
```
{URL}/products
```
- (GET) To get the order list as JSON
```
{URL}/orders-json
```
- (GET) To get the product detail by user ID as JSON
```
{URL}/order/:id
```
- (POST) To create the new order with quantity 
```
{URL}/orders/create
```
