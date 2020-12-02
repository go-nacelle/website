function updateSearchInput() {
  $("#search").val(getQuery());
}
updateSearchInput();

async function handleSearch() {
  const query = getQuery().trim().toLowerCase();
  if (query === "") {
    $(".search-results").html("No search query supplied.");
    return;
  }

  const pagesIndex = await (await fetch("/index.json")).json();

  const searchIndex = lunr(function () {
    this.field("title");
    this.field("content");
    this.ref("href");
    pagesIndex.forEach((page) => this.add(page));
  });

  const results = searchIndex
    .search(
      query
        .split(" ")
        .map((term) => `+${term.trim()}`)
        .join(" ")
        .substring(1)
    )
    .map((hit) => pagesIndex.find((page) => page.href === hit.ref));

  if (results.length === 0) {
    $(".search-results").html(`No pages match '${query}'.`);
    return;
  }

  $(".search-results ul").html(
    results
      .map(
        (hit) => `
          <li>
            <a href='${hit.href}'>${hit.title}</a>
          </li>
        `
      )
      .join("")
  );
}
handleSearch().catch((err) => console.error(err));

function getQuery() {
  return new URLSearchParams(window.location.search).get("q") || "";
}
