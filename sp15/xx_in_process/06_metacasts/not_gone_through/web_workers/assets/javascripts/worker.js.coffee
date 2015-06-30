self.addEventListener 'message', (e) ->
  obj =
    name: "Mark"
    site: "MetaCasts.tv"
  self.postMessage(obj)
, false