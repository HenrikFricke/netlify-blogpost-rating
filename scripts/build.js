require("dotenv").config();

const fetch = require("node-fetch");
const fs = require("fs");
const promisify = require("util").promisify;

const writeFile = promisify(fs.writeFile);

const LIKES_FORM_ID = process.env.LIKES_FORM_ID;
const API_KEY = process.env.API_KEY;

async function main() {
  const res = await fetch(
    `https://api.netlify.com/api/v1/forms/${LIKES_FORM_ID}/submissions`,
    {
      headers: {
        Authorization: `Bearer ${API_KEY}`
      }
    }
  );

  const body = await res.json();

  const likes = {};
  for (let i = 0; i < body.length; i++) {
    const submission = body[i];

    if (likes[submission.data.id]) {
      likes[submission.data.id] = likes[submission.data.id] + 1;
    } else {
      likes[submission.data.id] = 1;
    }
  }

  await writeFile("src/likes.json", JSON.stringify(likes), "utf8");
}

main();
