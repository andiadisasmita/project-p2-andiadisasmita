-- Create Users Table
CREATE TABLE Users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    deposit_amt DECIMAL(10, 2) NOT NULL DEFAULT 0.00,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create Categories Table
CREATE TABLE Categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT
);

-- Create Boardgames Table
CREATE TABLE Boardgames (
    id SERIAL PRIMARY KEY,
    name VARCHAR(150) NOT NULL,
    availability INTEGER NOT NULL CHECK (availability >= 0),
    rental_cost DECIMAL(10, 2) NOT NULL,
    category_id INTEGER REFERENCES Categories(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create Stock Table
CREATE TABLE Stock (
    id SERIAL PRIMARY KEY,
    boardgame_id INTEGER REFERENCES Boardgames(id),
    status VARCHAR(20) NOT NULL DEFAULT 'warehouse', -- warehouse, with_user, to_user, to_warehouse
    location VARCHAR(20) NOT NULL DEFAULT 'warehouse' -- warehouse, with_user, to_user, to_warehouse
);

-- Create RentalHistory Table
CREATE TABLE RentalHistory (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES Users(id),
    stock_id INTEGER REFERENCES Stock(id),
    rental_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    return_date TIMESTAMP,
    rental_cost DECIMAL(10, 2) NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'reserved' -- reserved, with_user, returned
);

-- Create Payments Table
CREATE TABLE Payments (
    id SERIAL PRIMARY KEY,
    rental_id INTEGER REFERENCES RentalHistory(id),
    amount DECIMAL(10, 2) NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'unpaid', -- paid, unpaid
    paid_at TIMESTAMP
);

-- Create Reviews Table
CREATE TABLE Reviews (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES Users(id),
    boardgame_id INTEGER REFERENCES Boardgames(id),
    rating INTEGER NOT NULL CHECK (rating BETWEEN 1 AND 5),
    comment TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT unique_user_boardgame_review UNIQUE (user_id, boardgame_id)
);
