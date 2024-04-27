-- +goose Up
-- +goose StatementBegin
INSERT INTO ingredients (ingredient_id, name)
VALUES
    (1, 'raw meat'),
    (2, 'raw bird drumstick'),
    (3, 'any seafood'),
    (4, 'raw bird thigh'),
    (5, 'raw prime meat'),
    (6, 'raw gourmet meat'),
    (7, 'raw whole bird'),
    (8, 'spicy pepper'),
    (9, 'hyrule bass');

INSERT INTO recipes (recipe_id, name, category)
VALUES
    (1, 'meat and seafood fry', 'surf and turf'),
    (2, 'prime meat and seafood fry', 'surf and turf'),
    (3, 'gourmet meat and seafood fry', 'surf and turf'),
    (4, 'spicy meat and seafood fry', 'surf and turf');


INSERT INTO recipe_ingredients (recipe_id, ingredient_id)
VALUES
    (1, 1),
    (1, 2),
    (1, 3),
    (2, 4),
    (2, 5),
    (2, 3),
    (3, 6),
    (3, 7),
    (3, 3),
    (4, 8),
    (4, 9),
    (4, 1);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM ingredients;
DELETE FROM recipes;
DELETE FROM recipe_ingredients;
-- +goose StatementEnd
