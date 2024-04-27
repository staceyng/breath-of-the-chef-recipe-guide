-- +goose Up
CREATE TABLE categories(
    category VARCHAR PRIMARY KEY
);

INSERT INTO categories (category)
VALUES
       ('vegetarian'),
       ('seafood'),
       ('poultry and meat'),
       ('surf and turf'),
       ('desserts'),
       ('monster'),
       ('sides and snacks');

CREATE TABLE recipes (
    recipe_id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    category VARCHAR REFERENCES categories(category) NOT NULL,
    UNIQUE (name)
);

CREATE TABLE ingredients (
    ingredient_id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    UNIQUE (name)
);

CREATE TABLE recipe_ingredients (
    recipe_id INT,
    ingredient_id INT,
    PRIMARY KEY (recipe_id, ingredient_id),
    CONSTRAINT fk_recipe_id
        FOREIGN KEY (recipe_id) REFERENCES recipes(recipe_id) ON DELETE CASCADE,
    CONSTRAINT fk_ingredient_id
        FOREIGN KEY (ingredient_id) REFERENCES ingredients(ingredient_id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE recipe_ingredients CASCADE;
DROP TABLE ingredients;
DROP TABLE recipes;
DROP TABLE categories;
