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
    picture VARCHAR(255) NOT NULL DEFAULT 'default.png',
    description LONGTEXT,
    category_id INT NOT NULL,
    draft BOOLEAN NOT NULL DEFAULT 0,
    open BOOLEAN NOT NULL DEFAULT 0,
    user_id INT NOT NULL
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

INSERT INTO users (username, email, password, roles) VALUES ('admin', 'admin@admin.com', '$2a$14$5o/fhb57T/hbPZyrYmjbYe1vpU17CaEY4GvYQNfXf4RxBhYFDKIfu', 3);
INSERT INTO users (username, email, password, roles) VALUES ('client', 'client@client.com', '$2a$14$mw1v2qkxJb2gWFypV4x0JuFI5L8dWx3JmTNTfM.JJKF4/86I4wWXa', 1);
INSERT INTO users (username, email, password, roles) VALUES ('restaurateur', 'restaurateur@restaurateur.com', '$2a$14$HnMJ4li8JFt8Z2eDMP.38OI7WpRc8m8M45yEnfa75/nJOhRljQ6Im', 2);
INSERT INTO users (username, email, password, roles) VALUES ('restau', 'restau@restau.com', '$2a$14$lEfQY9e/UYFGp.0PCHptq.xqYzaS1oBRae6xwMkgf1bsV6e49LfIO', 2);
INSERT INTO users (username, email, password, roles) VALUES ('restau2', 'restau2@restau.com', '$2a$14$vsesRWLHmv/nRPK.E0dEQOFPtycMeRIrSvp5p/W.ki4.4ay1ZOawu', 2);


/*fake restaurant*/

INSERT INTO restaurant (name, email, picture, description, category_id, draft, user_id) VALUES ('restaurant1', 'restaurant1@email.com', 'default', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum', 0, true, 3);
INSERT INTO restaurant (name, email, picture, description, category_id, draft, user_id) VALUES ('restaurant2', 'restaurant2@email.com', 'default', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum', 1, false, 4);
INSERT INTO restaurant (name, email, picture, description, category_id, draft, user_id) VALUES ('restaurant3', 'restaurant3@email.com', 'default', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum', 1, true, 5);

/*restaurant category*/

INSERT INTO restaurant_category (name) VALUES ('burger');
INSERT INTO restaurant_category (name) VALUES ('thaïlandais');
INSERT INTO restaurant_category (name) VALUES ('mexicain');

/*roles */

INSERT INTO roles(name) VALUES ('client');
INSERT INTO roles(name) VALUES ('restaurateur');
INSERT INTO roles(name) VALUES ('admin');
/**/
