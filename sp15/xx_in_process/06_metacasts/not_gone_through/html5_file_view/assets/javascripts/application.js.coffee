#= require jquery

$ ->

  files = []

  handleFileSelect = (evt) ->
    files = evt.target.files
    ul = $("#list")
    ul.html("")
    for f, ind in files
      if f.type is "image/jpeg"
        path = f.webkitRelativePath.split("/").join(" / ")
        ul.append("<li><a href='#' data-index='#{ind}'>#{path}</a></li>")
    $("#list a:first").click()

  $("#files").change(handleFileSelect)

  $("#list a").live "click", (e) ->
    e?.preventDefault()
    f = files[$(e.target).attr("data-index")]
    objectURL = window.URL.createObjectURL(f)
    $("#photo").attr("src", objectURL)