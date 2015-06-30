#= require jquery
#= require agnes
#= require filesaver

$ ->

  CSV = []

  buildTable = (event) ->
    $("tbody").html("")
    result = event.target.result
    CSV = agnes.csvToArray2d(result)
    if CSV.length > 0
      for row in CSV
        tr = $("<tr>")
        for column in row
          tr.append($("<td>").text(column))
        $("tbody").append(tr)
      $(".content").show()

  $("#file-selector").bind "change", (e) ->
    $(".content").hide()
    file = e.target.files[0]
    reader = new FileReader()
    reader.onload = buildTable
    reader.readAsText(file)

  buildCSVFile = ->
    csv = [["ID", "Name", "Duration"]]
    for row in CSV
      csv.push row[0..2]
    json = agnes.array2dToJson(csv)
    file = agnes.jsonToCsv(json)
    return file

  $(".export").click (e) ->
    e?.preventDefault()

    blob = new Blob([buildCSVFile()], {type: "text\/csv"})
    saveAs(blob, "data.csv")