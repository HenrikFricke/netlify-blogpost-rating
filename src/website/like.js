const likeContainer = document.querySelector("#likecontainer");

const button = document.createElement("button");
button.innerHTML = "I like";

const numberOfLikes = document.createElement("span");
numberOfLikes.innerHTML = "0 likes";

likeContainer.appendChild(numberOfLikes);
likeContainer.appendChild(button);
