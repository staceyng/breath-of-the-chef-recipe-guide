import { type RouteConfig, route } from "@react-router/dev/routes";

export default [
  route("/", "./routes/home.tsx"),
  route("/ingredient-picker", "./routes/ingredientPicker.tsx"),
  route("/recipe-list", "./routes/recipeList.tsx"),
  // pattern ^           ^ module file
] satisfies RouteConfig;
