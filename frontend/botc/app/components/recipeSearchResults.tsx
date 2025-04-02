type Recipe = {
  name: string;
  category: string;
  ingredients: string[];
  image: string | null | undefined;
};

interface Props {
  recipes: Recipe[];
}

function RecipeSearchResults({ recipes }: Props) {
  return (
    <div className="relative overflow-x-auto">
      <table className="w-full text-sm text-left rtl:text-right text-gray-500 dark:text-gray-400">
        <thead className="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
          <tr>
            <th scope="col" className="px-6 py-3">
              Recipe Name
            </th>
            <th scope="col" className="px-6 py-3">
              Category
            </th>
            <th scope="col" className="px-6 py-3">
              Ingredients
            </th>
          </tr>
        </thead>
        {recipes.map((recipe) => (
          <tbody>
            <tr className="bg-white border-b dark:bg-gray-800 dark:border-gray-700 border-gray-200">
              <th
                scope="row"
                className="px-6 py-4 font-medium text-gray-900 whitespace-nowrap dark:text-white"
              >
                {recipe.name}
              </th>
              <td className="px-6 py-4">{recipe.category}</td>
              <td className="px-6 py-4">{recipe.ingredients}</td>
            </tr>
          </tbody>
        ))}
      </table>
    </div>
  );
}

export default RecipeSearchResults;
