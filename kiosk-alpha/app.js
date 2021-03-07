const http = require("http");
const fs = require("fs");
const port = 3000;

const server = http.createServer(function (req, res) {
  res.writeHead(200, { "Content-Type": "text/html" });
  fs.readFile("index.html", function (e, data) {
    if (e) {
      res.writeHead(404);
      res.write("Error: File Not Found");
    } else {
      res.write(data);
    }
    res.end();
  });
});

server.listen(port, function (e) {
  if (e) {
    console.log("Something went wrong", e);
  } else {
    console.log("Server is listening on port " + port);
  }
});
