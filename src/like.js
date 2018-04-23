const likeContainer = document.querySelector("#likecontainer");
const numberOfLikes = document.createElement("span");

likeContainer.appendChild(numberOfLikes);

fetch(".netlify/functions/likes?path=" + window.location.pathname)
  .then(function(res) {
    return res.text();
  })
  .then(function(likes) {
    numberOfLikes.innerHTML = likes + " likes";
  });
