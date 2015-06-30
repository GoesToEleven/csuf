#= require jquery

$ ->

  worker = new Worker("assets/worker.js")

  worker.addEventListener 'message', (e) ->
    $("#main").append(JSON.stringify(e.data))
  , false

  worker.postMessage('Hello World!')

  $("#main").prepend("Some Text Here...")