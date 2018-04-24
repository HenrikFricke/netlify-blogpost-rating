const likeContainer = document.querySelector("#likecontainer");
const likeform = document.querySelector("#likeform");
const likeid = document.querySelector("#likeid");

const numberOfLikes = document.createElement("span");
likeContainer.appendChild(numberOfLikes);

fetch(".netlify/functions/likes?id=" + likeid.value)
  .then(function(res) {
    if (!res.ok) {
      return "";
    }
    return res.text();
  })
  .then(function(likes) {
    numberOfLikes.innerHTML = likes + " likes";
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
  }).then(() => {
    numberOfLikes.innerHTML = "Thanks for your feedback!";
  });
});
