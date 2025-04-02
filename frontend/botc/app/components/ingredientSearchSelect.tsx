import {
  Combobox,
  ComboboxInput,
  ComboboxOption,
  ComboboxOptions,
  ComboboxButton,
} from "@headlessui/react";
import { ChevronDownIcon } from "@heroicons/react/20/solid";

import { useState } from "react";

import RecipeSearchResults from "./recipeSearchResults";

interface Props {
  ingredients: string[];
}

function IngredientSearchSelect({ ingredients }: Props) {
  const [selectedIngredients, setSelectedIngredients] = useState<string[]>([]);
  const [query, setQuery] = useState("");

  const filteredIngredients =
    query === ""
      ? ingredients
      : ingredients.filter((ingredient) =>
          ingredient.toLowerCase().includes(query.toLowerCase())
        );

  const handleSelection = (ingredient: string) => {
    setSelectedIngredients((prevSelected) => {
      if (prevSelected.includes(ingredient)) {
        return prevSelected.filter((item) => item !== ingredient);
      } else {
        return [...prevSelected, ingredient];
      }
    });
    setQuery(""); // Clear the query after selection
  };

  return (
    <div className="grid justify-center gap-2">
      <h1>Select ingredients: </h1>
      <Combobox multiple>
        <div className="relative">
          <ComboboxInput
            key={selectedIngredients.join(",")} // reset input when selection change
            className="w-full border rounded-md py-2 px-3 focus:outline-none focus:ring-2 focus:ring-blue-500"
            onChange={(event) => setQuery(event.target.value)}
          />
          <ComboboxButton className="group absolute inset-y-0 right-0 px-2.5">
            <ChevronDownIcon className="size-4 fill-white/60 group-data-[hover]:fill-white" />
          </ComboboxButton>
          <ComboboxOptions className="absolute z-10 mt-1 max-h-60 w-full overflow-auto rounded-md bg-white py-1 text-base shadow-lg ring-1 ring-black/5 focus:outline-none">
            {filteredIngredients.length === 0 && query !== "" ? (
              <div className="relative cursor-default select-none px-4 py-2 text-gray-700">
                Nothing found.
              </div>
            ) : (
              filteredIngredients.map((ingredient) => (
                <ComboboxOption
                  key={ingredient}
                  value={ingredient}
                  onClick={() => handleSelection(ingredient)}
                  className={({ selected }) =>
                    `relative cursor-default select-none py-2 px-4 ${
                      selected ? "bg-blue-600 text-white" : "text-gray-900"
                    } ${
                      selectedIngredients.includes(ingredient)
                        ? "font-semibold"
                        : "font-normal"
                    }`
                  }
                >
                  {ingredient}
                </ComboboxOption>
              ))
            )}
          </ComboboxOptions>
        </div>
      </Combobox>
      {/* badges */}
      <div className="flex flex-wrap gap-2 mb-2">
        {selectedIngredients.map((ingredient) => (
          <span
            key={ingredient}
            className="inline-flex items-center rounded-md bg-gray-50 px-2 py-1 text-xs font-medium text-gray-600 ring-1 ring-gray-500/10 ring-inset"
          >
            {ingredient}
          </span>
        ))}
      </div>
      {/* search button */}
      <button
        className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
        onClick={() => {
          console.log("searching with ingredients:", selectedIngredients);
        }}
      >
        Search
      </button>
      <>
        <RecipeSearchResults
          recipes={[
            {
              name: "test",
              category: "test",
              ingredients: ["test"],
              image: "test",
            },
          ]}
        />
      </>
    </div>
  );
}

export default IngredientSearchSelect;
