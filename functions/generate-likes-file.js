const request = require("request");

exports.handler = function(event, context, callback) {
  request.post(
    {
      url: "https://api.netlify.com/build_hooks/" + process.env.BUILD_HOOK_KEY
    },
    function() {
      return callback(null, {
        statusCode: 200,
        body: "Done."
      });
    }
  );
};
