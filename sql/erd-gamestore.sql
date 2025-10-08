-- Toko Games Digital 
-- ERD Design
-- dbdiagram.io

Table Customers {
  customerID INT [pk, increment]
  name varchar(100)
  email email
  password password
  created_at timestamp
  updated_at timestamp
}

Table Games {
  gameID INT [pk, increment]
  categoryID int [ref: > categories.categoryID]
  titles varchar(100)
  price float64
  created_at timestamp
  updated_at timestamp
}

Table Categories {
  categoryID INT [pk, increment]
  name varchar(100)
}

Table Orders {
  orderID INT [pk, increment]
  customerID int [ref: > Customers.customerID]
  gameID int [ref: > Games.gameID]
  created_at timestamp
}

Table Payments {
  paymentID INT [pk, increment]
  orderID int [ref: > Orders.orderID]
  amount float64
  status varchar(50)
  created_at timestamp
}