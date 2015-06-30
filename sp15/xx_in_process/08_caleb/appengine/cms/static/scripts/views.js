var Views = {};

function renderHeader() {
  var loggedIn = !!EMAIL;
  return h("div.page-header", [
    loggedIn ?
    h("div.pull-right", [
      h("a.btn.btn-default", { "href": "#/logout" }, [
        "Logout"
      ])
    ]) : "",
    h("h1", "CMS Example")
  ]);
}

Views["/signup"] = function() {
  var emailInput, passwordInput1, passwordInput2;

  function onSubmit(evt) {
    evt.preventDefault();

    if (passwordInput1.value !== passwordInput2.value) {
      alert("passwords don't match");
      return;
    }

    var email = emailInput.value, password = passwordInput1.value;

    api("POST", "/api/users", {
      "email": email,
      "password": password
    }, function(result, error) {
      if (error) {
        alert("Error signing up: " + error);
        return;
      }
      EMAIL = email;
      location.href = "#/";
    });
  }

  return (
    h("div.container", [
      h("div.page-header", [
        h("h1", "CMS Example")
      ]),
      h("form", { "onsubmit": onSubmit }, [
        h("div.panel.panel-default", [
          h("div.panel-heading", [
            "Signup"
          ]),
          h("div.panel-body.form-horizontal", [
            h("div.form-group", [
              h("label.control-label.col-sm-2", { "for": "email" }, "Email"),
              h("div.col-sm-10", [
                emailInput = h("input.form-control", {
                  "type": "email",
                  "name": "email",
                  "placeholder": "email"
                })
              ])
            ]),
            h("div.form-group", [
              h("label.control-label.col-sm-2", { "for": "password" }, "Password"),
              h("div.col-sm-10", [
                passwordInput1 = h("input.form-control", {
                  "type": "password",
                  "name": "password1",
                  "placeholder": "password"
                })
              ])
            ]),
            h("div.form-group", [
              h("label.control-label.col-sm-2", { "for": "password" }, ""),
              h("div.col-sm-10", [
                passwordInput2 = h("input.form-control", {
                  "type": "password",
                  "name": "password2",
                  "placeholder": "repeat password"
                })
              ])
            ]),
            h("div.form-group", [
              h("div.col-sm-offset-2.col-sm-10", [
                h("button.btn.btn-default", { "type": "submit" }, [
                  "Signup"
                ]),
                " ",
                h("a.btn.btn-link", { "href": "#/login" }, "Login")
              ])
            ])
          ])
        ])
      ])
    ])
  );
};

Views["/login"] = function() {
  var emailInput, passwordInput;

  function onSubmit(evt) {
    evt.preventDefault();

    var email = emailInput.value, password = passwordInput.value;

    api("POST", "/api/users/login", {
      "email": email,
      "password": password
    }, function(result, error) {
      if (error) {
        alert("Error logging in: " + error);
        return;
      }
      EMAIL = email;
      location.href = "#/";
    });
  }

  return (
    h("div.container", [
      h("div.page-header", [
        h("h1", "CMS Example")
      ]),
      h("form", { "onsubmit": onSubmit }, [
        h("div.panel.panel-default", [
          h("div.panel-heading", "Login"),
          h("div.panel-body.form-horizontal", [
            h("div.form-group", [
              h("label.control-label.col-sm-2", { "for": "email" }, "Email"),
              h("div.col-sm-10", [
                emailInput = h("input.form-control", {
                  "type": "email",
                  "name": "email",
                  "placeholder": "email"
                })
              ])
            ]),
            h("div.form-group", [
              h("label.control-label.col-sm-2", { "for": "password" }, "Password"),
              h("div.col-sm-10", [
                passwordInput = h("input.form-control", {
                  "type": "password",
                  "name": "password",
                  "placeholder": "password"
                })
              ])
            ]),
            h("div.form-group", [
              h("div.col-sm-offset-2.col-sm-10", [
                h("button.btn.btn-default", { "type": "submit" }, [
                  "Login"
                ]),
                " ",
                h("a.btn.btn-link", { "href": "#/signup" }, "Signup")
              ])
            ])
          ])
        ])
      ])
    ])
  )
};

Views["/logout"] = function() {
  api("POST", "/api/users/logout", null, function(result, error) {
    if (error) {
      alert("Error logging out: " + error);
      return;
    }
    EMAIL = "";
    location.href = "#/";
  });
  return h("div");
};

Views["/documents/"] = function() {
  var id = location.hash.split("/").pop();
  var linkInput, contentsInput, fileList;

  function onSubmit(evt) {
    evt.preventDefault();

    var link = linkInput.value, contents = contentsInput.value, files = [];
    for (var i=0; i<fileList.childNodes.length; i++) {
      var li = fileList.childNodes[i];
      files.push({
        ID: li.getAttribute("data-file-id"),
        Name: li.getAttribute("data-file-name")
      });
    }

    api(id ? "PUT" : "POST", id ? "/api/documents/" + id : "/api/documents", {
      "Link": link,
      "Contents": contents,
      "Files": files
    }, function(result, error) {
      if (error) {
        alert("error saving document: " + error);
        return;
      }
      location.href = "#/";
    })
  }

  function onRemoveFile(evt) {
    var node = evt.target;
    while (node && node.tagName !== "LI") {
      node = node.parentNode;
    }
    node.parentNode.removeChild(node);
  }

  function onDropFile(evt) {
    evt.preventDefault();

    var fd = new FormData();
    fd.append("file", evt.dataTransfer.files[0]);
    var name = evt.dataTransfer.files[0].name;

    var progressBar, li;
    fileList.appendChild(
      li = h("li.list-group-item",
        h("div.progress",
          progressBar = h("div.progress-bar", name)
        )
      )
    );

    var xhr = new XMLHttpRequest();
    xhr.onreadystatechange = function() {
      if (xhr.readyState === 4) {
        var result = JSON.parse(xhr.responseText);
        li.parentNode.replaceChild(renderFile(result, name), li);
      }
    };
    xhr.upload.onprogress = function(evt) {
      progressBar.style.width = Math.ceil(100 * (evt.loaded / evt.total)) + "%";
    };
    xhr.open("POST", "/api/files");
    xhr.send(fd);
  }

  function renderFile(id, name) {
    return h("li.list-group-item", {
      "data-file-id": id,
      "data-file-name": name
    }, [
      h("span.pull-right", h("span.glyphicon.glyphicon-remove-sign", { onclick: onRemoveFile })),
      h("a", {
        "href": "/api/files/" + id + "?name=" + encodeURIComponent(name),
        "download": true
      }, name)
    ]);
  }

  var el = (
    h("div.container", [
      renderHeader(),

      h("form", { "onsubmit": onSubmit }, [
        h("div.panel.panel-default", [
          h("div.panel-heading", id ? "Edit Document" : "New Document"),
          h("div.panel-body.form-horizontal", [
            h("div.form-group", [
              h("label.control-label.col-sm-2", { "for": "link" }, "Link"),
              h("div.col-sm-10", [
                linkInput = h("input.form-control", {
                  "type": "text",
                  "name": "link",
                  "placeholder": "Link"
                })
              ])
            ]),
            h("div.form-group", [
              h("label.control-label.col-sm-2", { "for": "contents" }, "Contents"),
              h("div.col-sm-10", [
                contentsInput = h("textarea.form-control", {
                  "name": "contents",
                  "placeholder": "Contents"
                })
              ])
            ]),
            h("div.form-group", [
              h("label.control-label.col-sm-2", {  }, "Files"),
              h("div.col-sm-10", [
                fileList = h("ul.list-group"),
                h("div.file-target", {
                  ondragover: function(evt) {
                    evt.preventDefault();
                  },
                  ondragend: function(evt) {
                    evt.preventDefault();
                  },
                  ondrop: onDropFile
                }, "Drop File Here")
              ])
            ]),
            h("div.form-group", [
              h("div.col-sm-offset-2.col-sm-10", [
                h("button.btn.btn-primary", { "type": "submit" }, "Save"),
                " ",
                h("a.btn.btn-link", { "href": "#/" }, "Back")
              ])
            ])
          ])
        ])
      ])
    ])
  );

  if (id) {
    api("GET", "/api/documents/" + id, null, function(doc, error) {
      if (error) {
        alert("error retrieving document: " + error);
        return;
      }

      linkInput.value = doc.Link;
      contentsInput.textContent = doc.Contents;
      var files = doc.Files || [];
      for (var i=0; i<files.length; i++) {
        var f = files[i];
        fileList.appendChild(renderFile(f.ID, f.Name));
      }
    });
  }

  return el;
};

function renderDocumentList() {
  var el, tbody;

  function onRemove(evt) {
    var node = evt.target;
    while (node && !node.getAttribute("data-id")) {
      node = node.parentNode;
    }
    var id = node.getAttribute("data-id");
    api("DELETE", "/api/documents/" + id, null, function(result, error) {
      if (error) {
        alert("error deleting document: " + error);
        return;
      }
      el.parentNode.replaceChild(renderDocumentList(), el);
    });
  }

  el = (
    h("div.panel.panel-default", [
      h("div.panel-heading", "Documents"),
      h("div.panel-body", [
        h("table.table", [
          h("thead", [
            h("tr", [
              h("th", "Link"),
              h("th", "Contents"),
              h("th")
            ])
          ]),
          h("tfoot", [
            h("tr", [
              h("td", { "colSpan": "3" }, [
                h("a.btn.btn-primary", { "href": "#/documents/" }, "New Document")
              ])
            ])
          ]),
          tbody = h("tbody")
        ])
      ])
    ])
  );

  api("GET", "/api/documents", null, function(docs, error) {
    if (error) {
      alert("error listing documents: " + error);
      return;
    }
    for (var i=0; i<docs.length; i++) {
      var doc = docs[i];
      tbody.appendChild(
        h("tr", { "data-id": doc.ID }, [
          h("td", [
            h("a", { "href": "#/documents/" + doc.ID }, doc.Link)
          ]),
          h("td", [
            doc.Contents
          ]),
          h("td", h("span.glyphicon.glyphicon-remove-sign", {
            onclick: onRemove
          }))
        ])
      );
    }
  });

  return el;
}

Views["/"] = function() {
  var loggedIn = !!EMAIL;
  return (
    h("div.container", [
      renderHeader(),
      loggedIn ?
      renderDocumentList() :
      h("div.btn-group", {
        "role": "group"
      }, [
        h("a.btn.btn-default", { "href": "#/login" }, [
          "Login"
        ]),
        h("a.btn.btn-default", { "href": "#/signup" }, [
          "Signup"
        ])
      ])
    ])
  );
};
