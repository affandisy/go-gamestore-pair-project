-- Toko Games Digital 
-- ERD Design
-- dbdiagram.io

Table Customers {
  customerID INT [pk, increment]
  name varchar(100) not null
  email varchar(100) not null unique
  password varchar(255) not null
  createdAt timestamp
  updatedAt timestamp
}

Table Games {
  gameID INT [pk, increment]
  categoryID int not null [ref: > categories.categoryID] not null
  title varchar(100) not null
  price float64
  createdAt timestamp
  updatedAt timestamp
}

Table Categories {
  categoryID INT [pk, increment]
  name varchar(100)
}

Table Orders {
  orderID INT [pk, increment]
  customerID int not null [ref: > Customers.customerID]
  gameID int not null [ref: > Games.gameID]
  createdAt timestamp
}

Table Payments {
  paymentID INT [pk, increment]
  customerID int not null [ref: > Customers.customerID] 
  amount float64
  status varchar(50)
  createdAt timestamp
}

Table Library {
  LibraryID int [pk, increment]
  CustomerID int [ref: > Customers.CustomerID, not null]
  GameID int [ref: > Games.GameID, not null]
  CreatedAt timestamptz [not null, default: `NOW()`]
}