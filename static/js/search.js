function setupFocusHandler() {
  $(document).keydown(function(ev) {
    if (ev.keyCode === 191 && !$("#search").is(":focus")) {
      $("#search").val('').focus();
      return false;
    }
  });
}
setupFocusHandler();

function setupFocusedHandler() {
  $("#search").focus(function() {
    $('.search-wrapper kbd').hide();
  });
  $("#search").blur(function() {
    $('.search-wrapper kbd').show();
  })
}
setupFocusedHandler();

function updateSearchInput() {
  $("#search").val(getQuery());
}
updateSearchInput();

function handleSearch() {
  const query = getQuery().trim().toLowerCase();
  if (query === "") {
    return showNoQuery();
  }

  fetch("/index.json")
    .then((response) => response.json())
    .then((pagesIndex) => showResults(query, search(pagesIndex, query)))
    .catch((err) => console.error(err));
}
handleSearch();

//
// Helpers

function getQuery() {
  return new URLSearchParams(window.location.search).get("q") || "";
}

function search(pagesIndex, query) {
  query = query
    .split(" ")
    .map((term) => `+${term.trim()}`)
    .join(" ")
    .substring(1);

  return makeIndex(pagesIndex)
    .search(query)
    .map((hit) => pagesIndex.find((page) => page.href === hit.ref));
}

function makeIndex(pagesIndex) {
  return lunr(function () {
    this.field("title");
    this.field("content");
    this.ref("href");
    pagesIndex.forEach((page) => this.add(page));
  });
}

function showNoQuery() {
  $("#search-results").html("No search query supplied.");
}

function showResults(query, results) {
  if (results.length === 0) {
    $("#search-results").html(`No pages match '${query}'.`);
  } else {
    const items = results.map(
      (hit) => `
        <li>
          <a href='${hit.href}'>${hit.title}</a>
        </li>
      `
    );

    $("#search-results").html(`<ul>${items.join("")}</ul>`);
  }
}
