function api(method, endpoint, data, callback) {
  var xhr = new XMLHttpRequest();
  xhr.open(method, endpoint);
  if (data) {
    xhr.send(JSON.stringify(data));
  } else {
    xhr.send(null);
  }
  xhr.onreadystatechange = function(evt) {
    if (xhr.readyState === 4) {
      var res;
      try {
        res = JSON.parse(xhr.responseText);
      } catch(e) {
        res = { "error": xhr.responseText };
      }
      if (res && res.error) {
        callback(null, res.error);
      } else {
        callback(res);
      }
    }
  };
}

function main() {
  var prev = null;
  function renderView(url) {
    if (prev) {
      document.body.removeChild(prev);
    }
    var view = Views[url];
    if (!view) {
      var arr = url.split("/");
      arr.pop();
      view = Views[arr.join("/") + "/"];
    }
    prev = view();
    document.body.appendChild(prev);
  }
  window.addEventListener("hashchange", function() {
    renderView(location.hash.substr(1) || "/");
  }, false);
  renderView(location.hash.substr(1) || "/");
}

main();
