window.onload = function() {
  //<editor-fold desc="Changeable Configuration Block">
  // the following lines will be replaced by docker/configurator, when it runs in a docker-container
  var file_list = []
  fetch("http://localhost:8080/file/").then(function(response) {
    //console.log(response.headers.values)
    //console.log(response.headers.get("file_data"))
   return response.headers.get("file_data");
  }).then(function(data) {
    const arr_list = data.split(",")
    arr_list.forEach(file => {
      var file_name = file.replace(".swagger.json","")
      var obj = {url:`./proto/${file}`,name:`${file_name}`}
      file_list.push(obj)  
    });
    window.ui = SwaggerUIBundle({
      urls: file_list,
      "urls.primaryName": "home",
      dom_id: '#swagger-ui',
      validatorUrl : null,
      deepLinking: true,
      presets: [
        SwaggerUIBundle.presets.apis,
        SwaggerUIStandalonePreset
      ],
      plugins: [
        SwaggerUIBundle.plugins.DownloadUrl
      ],
      layout: "StandaloneLayout"
    });
  }).catch(function() {
    console.log("Booo");
  });
};