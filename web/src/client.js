const axios = require('axios').default;

// TODO: Separate to external config file
async function getFromBackend(network) {
  const result = await axios
    .get('api/watch?network='+network)
    .then((response) => response.data)
    .catch((error) => {
      console.log(`failed to get data from server: ${error}`);
      return error;
    });
  return result;
}

async function Watch(network) {
  const response = await getFromBackend(network)
    .then((response) => response);
  return response;
}

export default Watch;
