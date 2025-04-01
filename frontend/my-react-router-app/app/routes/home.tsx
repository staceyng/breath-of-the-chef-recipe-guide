import { useNavigate } from "react-router";

function Home() {
  let navigate = useNavigate();

  const onIngredientPickerButtonClick = () => {
    navigate("/ingredient-picker");
  };

  const onRecipeListButtonClick = () => {
    navigate("/recipe-list");
  };

  return (
    <div className="grid justify-center">
      <div className="text-center p-4">
        <h1 className="text-2xl">Breath of the chef guide</h1>
      </div>
      <button
        className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded my-4"
        onClick={onIngredientPickerButtonClick}
      >
        Ingredient picker
      </button>
      <button
        className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded my-4"
        onClick={onRecipeListButtonClick}
      >
        Recipe list
      </button>
    </div>
  );
}

export default Home;
