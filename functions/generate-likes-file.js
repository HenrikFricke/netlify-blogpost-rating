const fetch = require("node-fetch");

exports.handler = async function(event, context, callback) {
  await fetch(
    "https://api.netlify.com/build_hooks/" + process.env.BUILD_HOOK_KEY,
    {
      method: "POST"
    }
  );

  callback(null, {
    statusCode: 200
  });
};
