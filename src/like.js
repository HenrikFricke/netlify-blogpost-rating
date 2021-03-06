const likeContainer = document.querySelector("#likecontainer");
const likeform = document.querySelector("#likeform");
const likeid = document.querySelector("#likeid");

fetch("likes.json")
  .then(function(res) {
    if (!res.ok) {
      return "";
    }
    return res.json();
  })
  .then(function(res) {
    const likes = res[likeid.value] || 0;
    likeContainer.innerHTML = likes + " likes";
  });

likeform.addEventListener("submit", function(event) {
  event.preventDefault();

  const inputElements = event.target.elements;
  const action = event.target.action;

  var body = new FormData();
  for (let i = 0; i < inputElements.length; i++) {
    const element = inputElements[i];
    if (element.name && element.value) {
      body.append(element.name, element.value);
    }
  }

  fetch(action, {
    method: "POST",
    body: body
  })
    .then(() => {
      likeContainer.innerHTML = "Thanks for your feedback!";
    })
    .then(() => fetch(".netlify/functions/generate-likes-file"));
});
