const form = document.querySelector("form");
const searchResultsDiv = document.querySelector("#search-results");

form.addEventListener("submit", async (e) => {
  e.preventDefault();
  let name = document.querySelector("#name").value;
  name = name.replace(" ", "-");

  let headers = new Headers();
  headers.append("Content-Type", "application/json");
  headers.append("Accept", "application/json");
  headers.append("Access-Control-Allow-Origin", "http://localhost:8080");
  headers.append("Access-Control-Allow-Credentials", "true");
  headers.append("GET", "OPTIONS");

  const response = await fetch(
    `http://localhost:8080/recipes?name=${name}`,
    headers
  );
  const result = await response.json();

  // Update the HTML to display the results
  searchResultsDiv.innerHTML = "";
  const recipe = result.data;
  const recipeHTML = `
      <h3>${recipe.name}</h3>
      <p>Category: ${recipe.category}</p>
      <p>Ingredients:</p>
        <ul>
          ${recipe.ingredients
            .map((ingredient) => `<li>${ingredient}</li>`)
            .join("")}
        </ul>
    `;
  searchResultsDiv.innerHTML += recipeHTML;
});
