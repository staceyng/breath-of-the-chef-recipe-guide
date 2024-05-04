const form = document.querySelector("form");
const searchResultsDiv = document.querySelector("#search-results");

form.addEventListener("submit", async (e) => {
  e.preventDefault();
  let name = document.querySelector("#name").value;
  name = name.replaceAll(" ", "-");

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
    <br>
    <div class="card" style="width: 20rem;">
      <img src="../src/imgs/${name}.png" class="card-img-top" alt="...">
      <div class="card-body">
        <h5 class="card-title">${recipe.name}</h5>
        <h5><span class="badge text-bg-secondary">${recipe.category}</span></h5>
        <p class="card-text">Ingredients</p>
        <ul>
          ${recipe.ingredients
            .map((ingredient) => `<li>${ingredient}</li>`)
            .join("")}
        </ul>
      </div>
    </div>
  `;
  searchResultsDiv.innerHTML += recipeHTML;
});
