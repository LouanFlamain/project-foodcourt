CREATE DATABASE IF NOT EXISTS foodcourt_db;
USE foodcourt_db;

CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    picture VARCHAR(255) NOT NULL DEFAULT 'default',
    roles INT NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS roles (
    id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS restaurant (
    id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    picture VARCHAR(255) NOT NULL DEFAULT 'default',
    description LONGTEXT,
    category_id INT NOT NULL,
    draft BOOLEAN NOT NULL DEFAULT 0,
    open BOOLEAN NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS restaurant_category (
    id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS commande (
    id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
    date DATETIME NOT NULL,
    user_id INT NOT NULL,
    restaurant_id INT NOT NULL,
    content VARCHAR(255) NOT NULL,
    commentaire VARCHAR(255),
    state INT NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS commande_state (
    id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS carte (
    id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
    restaurant_id INT NOT NULL,
    description TEXT NOT NULL,
    price FLOAT NOT NULL
);

CREATE TABLE IF NOT EXISTS product (
    id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
    product VARCHAR(255) NOT NULL,
    price FLOAT NOT NULL,
    carte_id INT NOT NULL ,
    category_id INT NOT NULL
    
);
CREATE TABLE IF NOT EXISTS product_category_type (
    id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
    name VARCHAR(255) NOT NULL
    
);

CREATE TABLE IF NOT EXISTS feedback (
    id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
    commande_id INT NOT NULL,
    restaurant_id INT NOT NULL,
    description VARCHAR(255)
);

/*fake users*/

INSERT INTO users (username, email, password, roles) VALUES ('admin', 'admin@email.com', 'admin', 2);
INSERT INTO users (username, email, password, roles) VALUES ('client', 'client@email.com', 'client', 0);
INSERT INTO users (username, email, password, roles) VALUES ('restaurateur', 'restaurateur@email.com', 'restaurateur', 1);


/*fake restaurant*/

INSERT INTO restaurant (name, email, picture, description, category_id, draft) VALUES ('restaurant1', 'restaurant1@email.com', 'default', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum', 0, true);
INSERT INTO restaurant (name, email, picture, description, category_id, draft) VALUES ('restaurant2', 'restaurant2@email.com', 'default', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum', 1, false);
INSERT INTO restaurant (name, email, picture, description, category_id, draft) VALUES ('restaurant3', 'restaurant3@email.com', 'default', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum', 1, true);

/*restaurant category*/

INSERT INTO restaurant_category (name) VALUES ('burger');
INSERT INTO restaurant_category (name) VALUES ('thaïlandais');
INSERT INTO restaurant_category (name) VALUES ('mexicain');

/*roles */

INSERT INTO roles(name) VALUES ('client');
INSERT INTO roles(name) VALUES ('restaurateur');
INSERT INTO roles(name) VALUES ('admin');
/**/

-- carte

INSERT INTO carte (restaurant_id, description, price) VALUES

(1, 'Plat 1', 10.99),
(2, 'Plat 2', 12.99),
(3, 'Plat 3', 8.99);

-- FIN

-- product

INSERT INTO product (product, price, carte_id, category_id) VALUES

('Produit 1', 5.99, 1, 1),
('Produit 2', 8.99, 1, 2),
('Produit 3', 6.99, 2, 1),
('Produit 4', 10.99, 2, 3),
('Produit 5', 4.99, 3, 2);

-- FIN

-- product_category_type

INSERT INTO product_category_type (name) VALUES
('Entrée'),
('Plat principal'),
('Dessert');

-- FIN


