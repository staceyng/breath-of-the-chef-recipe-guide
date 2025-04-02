import IngredientSearchSelect from "../components/ingredientSearchSelect";

function IngredientPicker() {
  const ingredients = ["hyrule-bass", "seafood", "raw prime meat"];
  return (
    <div className="grid justify-center">
      <div className="text-center p-4">
        <h1 className="text-2xl">Ingredient Picker</h1>
      </div>
      <IngredientSearchSelect ingredients={ingredients} />
    </div>
  );
}

export default IngredientPicker;
