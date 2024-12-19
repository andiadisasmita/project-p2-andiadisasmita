-- Insert Sample Data into Users
INSERT INTO Users (email, password, deposit_amt) VALUES
('alice@example.com', 'hashed_password1', 50.00),
('bob@example.com', 'hashed_password2', 30.00),
('carol@example.com', 'hashed_password3', 20.00),
('dave@example.com', 'hashed_password4', 10.00),
('eve@example.com', 'hashed_password5', 0.00);

-- Insert Sample Data into Categories
INSERT INTO Categories (name, description) VALUES
('Strategy', 'Games that involve deep planning and strategy.'),
('Family', 'Games suitable for all ages and family fun.'),
('Card Games', 'Games that revolve around cards as the main component.'),
('Cooperative', 'Games where players work together to achieve a goal.'),
('Party', 'Games designed for social and casual fun.');

-- Insert Sample Data into Boardgames
INSERT INTO Boardgames (name, availability, rental_cost, category_id) VALUES
('Terraforming Mars', 5, 20.00, 1),
('Ticket to Ride', 0, 15.00, 2),
('Exploding Kittens', 10, 7.00, 3),
('Pandemic', 3, 12.00, 4),
('Codenames', 1, 10.00, 5);

-- Insert Sample Data into Stock
INSERT INTO Stock (boardgame_id, status, location) VALUES
(1, 'warehouse', 'warehouse'),
(1, 'to_user', 'to_user'),
(2, 'with_user', 'with_user'),
(3, 'warehouse', 'warehouse'),
(4, 'to_warehouse', 'to_warehouse');

-- Insert Sample Data into RentalHistory
INSERT INTO RentalHistory (user_id, stock_id, rental_cost, status) VALUES
(1, 1, 20.00, 'reserved'),
(2, 2, 15.00, 'to_user'),
(3, 3, 7.00, 'with_user'),
(4, 4, 12.00, 'to_warehouse'),
(5, 5, 10.00, 'returned');

-- Insert Sample Data into Payments
INSERT INTO Payments (rental_id, amount, status, paid_at) VALUES
(1, 20.00, 'paid', '2024-12-01 10:00:00'),
(2, 15.00, 'unpaid', NULL),
(3, 7.00, 'paid', '2024-12-03 15:30:00'),
(4, 12.00, 'unpaid', NULL),
(5, 10.00, 'paid', '2024-12-05 18:00:00');

-- Insert Sample Data into Reviews
INSERT INTO Reviews (user_id, boardgame_id, rating, comment) VALUES
(1, 1, 5, 'Terraforming Mars is amazing! Very challenging.'),
(3, 3, 4, 'Exploding Kittens is a great casual game.'),
(2, 5, 4, 'Codenames was fun for our game night.'),
(3, 4, 5, 'Pandemic is fantastic for cooperative play.');
