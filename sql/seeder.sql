INSERT INTO categories (name) VALUES
('Action'),
('Adventure'),
('RPG'),
('Strategy'),
('Sports'),
('Racing'),
('Shooter'),
('Puzzle'),
('Horror'),
('Fighting')
;


INSERT INTO Games (categoryID, title, price) VALUES
-- Action
(1, 'God of War', 799000),
(1, 'Devil May Cry 5', 599000),
(1, 'Sekiro: Shadows Die Twice', 749000),

-- Adventure
(2, 'Uncharted 4: A Thief’s End', 699000),
(2, 'Tomb Raider', 299000),
(2, 'Life is Strange', 249000),

-- RPG
(3, 'The Witcher 3: Wild Hunt', 399000),
(3, 'Elden Ring', 899000),
(3, 'Skyrim Special Edition', 349000),

-- Strategy
(4, 'Civilization VI', 499000),
(4, 'Age of Empires IV', 599000),
(4, 'Total War: Warhammer III', 899000),

-- Sports
(5, 'FIFA 25', 899000),
(5, 'NBA 2K25', 899000),
(5, 'Tony Hawk’s Pro Skater 1+2', 429000),

-- Racing
(6, 'Forza Horizon 5', 799000),
(6, 'Need for Speed: Heat', 599000),
(6, 'F1 24', 899000),

-- Shooter
(7, 'Call of Duty: Modern Warfare III', 999000),
(7, 'Apex Legends', 0),
(7, 'Overwatch 2', 0),

-- Puzzle
(8, 'Portal 2', 99000),
(8, 'The Witness', 149000),
(8, 'Tetris Effect', 199000),

-- Horror
(9, 'Resident Evil 4 Remake', 899000),
(9, 'Outlast', 79000),
(9, 'The Medium', 429000),

-- Fighting
(10, 'Tekken 8', 999000),
(10, 'Street Fighter 6', 899000),
(10, 'Mortal Kombat 1', 999000);
