const sleep = ms => new Promise(resolve => setTimeout(resolve, ms));
const urls = [
  "server.example.com:8080"
]
const promises = urls.map(x => fetch(`http://${x}/ping`))
Promise.all(promises)
  .then(() => sleep(200)) // timeout for max ping
  .then(() => {
    const results = performance.getEntriesByType("resource")
    return results.map(res => {
      return {
        name: res.name.substr(7,10),
        time: (res.responseStart - res.requestStart).toFixed(2)
      }
    })
  })