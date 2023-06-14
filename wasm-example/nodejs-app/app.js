const fs = require("fs");
const { loadPolicy } = require("@open-policy-agent/opa-wasm");
const axios = require('axios');

// Define the URL of the raw file on GitHub
const rawFileUrl = 'https://raw.githubusercontent.com/pkbhowmick/opa-demo/master/wasm-example/nodejs-app/policy.wasm';

// Make a GET request to fetch the raw file
axios.get(rawFileUrl)
  .then(response => {
    policyWasm = response.data
      // Load the policy module asynchronously
loadPolicy(policyWasm).then((policy) => {
  // Use console parameters for the input, do quick
  // validation by json parsing. Not efficient.. but
  // will raise an error
  const input = JSON.parse(process.argv[2]);
  // Provide a data document with a string value
  policy.setData({ world: "world" });

  // Evaluate the policy and log the result
  const result = policy.evaluate(input);
  console.log(JSON.stringify(result, null, 2));
}).catch((err) => {
  console.log("ERROR: ", err);
  process.exit(1);
});
  })
  .catch(error => {
    console.error('Error reading raw file:', error);
  });





