import IngredientCard from "../components/ingredientCard";

function IngredientPicker() {
  return (
    <div className="grid justify-center">
      <div className="text-center p-4">
        <h1 className="text-2xl">Ingredient Picker</h1>
      </div>
      <IngredientCard name="hyrule-bass" description="fish" />
    </div>
  );
}

export default IngredientPicker;
